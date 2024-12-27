package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite" // Новый драйвер SQLite
)

var db *sql.DB

type Booking struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TableNumber int    `json:"table_number"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

func main() {
	// Подключение к базе данных
	var err error
	db, err = sql.Open("sqlite", "./database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализация базы данных
	initDB()

	// Маршруты
	http.HandleFunc("/api/bookings/options", handleOptions)
	http.HandleFunc("/api/bookings/", handleBookings)
	http.HandleFunc("/api/bookings", handleBookings)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// initDB создаёт таблицу Booking и заполняет её данными
func initDB() {
	query := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		table_number INTEGER NOT NULL,
		date TEXT NOT NULL,
		start_time TEXT NOT NULL,
		end_time TEXT NOT NULL
	);
	INSERT OR IGNORE INTO bookings (id, name, table_number, date, start_time, end_time)
	VALUES
		(1, 'Alice', 1, '2024-12-28', '18:00', '20:00'),
		(2, 'Bob', 2, '2024-12-28', '19:00', '21:00');
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}
}



// Обработчик OPTIONS для CORS
func handleOptions(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.WriteHeader(http.StatusOK)
}



func enableCors(w http.ResponseWriter) {
	// Разрешаем все домены
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Разрешаем методы GET, POST, PUT, DELETE и OPTIONS
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// Разрешаем заголовки Content-Type
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// handleBookings обрабатывает запросы к /api/bookings
func handleBookings(w http.ResponseWriter, r *http.Request) {
	// Включаем CORS
	enableCors(w)

	switch r.Method {
	case http.MethodGet:
		getBookings(w)
	case http.MethodPost:
		createBooking(w, r)
	case http.MethodPut:
		updateBooking(w, r)
	case http.MethodDelete:
		deleteBooking(w, r)
	case http.MethodOptions:
		// Для OPTIONS-запросов сразу возвращаем OK
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}


// getBookings возвращает список всех бронирований
func getBookings(w http.ResponseWriter) {
	rows, err := db.Query("SELECT id, name, table_number, date, start_time, end_time FROM bookings")
	if err != nil {
		http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bookings []Booking
	for rows.Next() {
		var b Booking
		if err := rows.Scan(&b.ID, &b.Name, &b.TableNumber, &b.Date, &b.StartTime, &b.EndTime); err != nil {
			http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// createBooking добавляет новое бронирование
func createBooking(w http.ResponseWriter, r *http.Request) {
	var b Booking
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Проверка пересечений времени
	if isTimeOverlap(b.TableNumber, b.Date, b.StartTime, b.EndTime) {
		http.Error(w, "Время бронирования пересекается с существующим", http.StatusConflict)
		return
	}

	_, err := db.Exec("INSERT INTO bookings (name, table_number, date, start_time, end_time) VALUES (?, ?, ?, ?, ?)",
		b.Name, b.TableNumber, b.Date, b.StartTime, b.EndTime)
	if err != nil {
		http.Error(w, "Ошибка добавления данных", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Бронирование добавлено"))
}

// updateBooking обновляет существующее бронирование
func updateBooking(w http.ResponseWriter, r *http.Request) {
	var b Booking
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Проверка пересечений времени
	if isTimeOverlap(b.TableNumber, b.Date, b.StartTime, b.EndTime, b.ID) {
		http.Error(w, "Время бронирования пересекается с существующим", http.StatusConflict)
		return
	}

	_, err := db.Exec("UPDATE bookings SET name = ?, table_number = ?, date = ?, start_time = ?, end_time = ? WHERE id = ?",
		b.Name, b.TableNumber, b.Date, b.StartTime, b.EndTime, b.ID)
	if err != nil {
		http.Error(w, "Ошибка обновления данных", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Бронирование обновлено"))
}

// deleteBooking удаляет бронирование
func deleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM bookings WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Ошибка удаления данных", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Бронирование удалено"))
}

// isTimeOverlap проверяет пересечение времени бронирования
func isTimeOverlap(tableNumber int, date, startTime, endTime string, excludeID ...int) bool {
	query := "SELECT COUNT(*) FROM bookings WHERE table_number = ? AND date = ? AND (start_time < ? AND end_time > ?)"
	args := []interface{}{tableNumber, date, endTime, startTime}

	if len(excludeID) > 0 {
		query += " AND id != ?"
		args = append(args, excludeID[0])
	}

	var count int
	err := db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		log.Printf("Ошибка проверки пересечений: %v", err)
		return true
	}

	return count > 0
}
