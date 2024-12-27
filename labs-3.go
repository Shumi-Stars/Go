package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



// 1-я задача
func handler(c *gin.Context) {
	// Получение query-параметров
	name := c.DefaultQuery("name", "")
	age := c.DefaultQuery("age", "")

	// Проверка, переданы ли параметры
	if name == "" || age == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствуют параметры name или age"})
		return
	}

	// Формирование ответа
	response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)
	c.String(http.StatusOK, response)
}




// 2-я задача
func getParams(c *gin.Context) (float64, float64, error) {
	// Получение параметров
	aStr := c.DefaultQuery("a", "")
	bStr := c.DefaultQuery("b", "")

	// Проверка, что параметры не пустые
	if aStr == "" || bStr == "" {
		return 0, 0, fmt.Errorf("отсутствуют параметры a или b")
	}

	// Преобразование строк в числа
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("параметр a должен быть числом")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("параметр b должен быть числом")
	}

	return a, b, nil
}

func addHandler(c *gin.Context) {
	a, b, err := getParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := a + b
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func subHandler(c *gin.Context) {
	a, b, err := getParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := a - b
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func mulHandler(c *gin.Context) {
	a, b, err := getParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := a * b
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func divHandler(c *gin.Context) {
	a, b, err := getParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if b == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "деление на 0 невозможно"})
		return
	}
	result := a / b
	c.JSON(http.StatusOK, gin.H{"result": result})
}





// Структура для запроса в 3-й задаче
type RequestBody struct {
	Text string `json:"text"`
}


// 3-я задача
func countCharacters(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат JSON"})
		return
	}

	// Подсчёт символов
	counts := make(map[rune]int)
	for _, char := range requestBody.Text {
		counts[char]++
	}

	// Возвращение результата
	c.JSON(http.StatusOK, counts)
}






func main() {
	// Создание роутера
	r := gin.Default()

	// 1-я "http://localhost:8080/?name=Alice&age=19"
	r.GET("/", handler) 

	// 2-я 
	r.GET("/add", addHandler) // "http://localhost:8080/add?a=10&b=5"

	r.GET("/sub", subHandler) // "http://localhost:8080/sub?a=10&b=5"

	r.GET("/mul", mulHandler) // "http://localhost:8080/mul?a=10&b=5"

	r.GET("/div", divHandler) // "http://localhost:8080/div?a=10&b=5"


	// 3-я  http://localhost:8080/count
	r.POST("/count", countCharacters)

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}