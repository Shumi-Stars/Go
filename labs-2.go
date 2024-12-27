package main

import (
	"fmt"
	"errors"
	"math"
)



// Задание 1. Массивы и срезы
// 1.1 formatIP принимает массив из четырех байт и возвращает строку с IP-адресом в формате "x.x.x.x"
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// 1.2 listEven принимает два целых числа, и возвращает срез с четными числами в заданном диапазоне и ошибку, если диапазон некорректный
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var slice []int
	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			slice = append(slice, i)
		}
	}
	return slice, nil
}



// Задание 2
// countCharacters принимает строку и возвращает карту, где каждому символу соответствует количество его вхождений
func countCharacters(s string) map[rune]int {
	// Создаем пустую карту для хранения результатов
	charCount := make(map[rune]int)
	
	
	for index := 0; index < len(s); index++ {
		// Получаем символ по индексу и преобразуем его в rune
		char := rune(s[index])
		
		// Увеличиваем счетчик для символа в карте
		charCount[char]++
	}
	
	
	return charCount
}



// Задание 3
type Point struct {
	X float64 
	Y float64
}


type Segment struct {
	Start Point
	End Point
}

// Метод  вычисляет длину отрезка
func (s Segment) Length() float64 {
	return math.Sqrt(math.Pow(s.End.X-s.Start.X, 2) + math.Pow(s.End.Y-s.Start.Y, 2))
}


type Triangle struct {
	A Point
	B Point 
	C Point
}

// Метод Area для Triangle возвращает площадь треугольника
func (t Triangle) Area() float64 {
	// Используем формулу Герона для нахождения площади треугольника
	a := Segment{t.A, t.B}.Length()
	b := Segment{t.B, t.C}.Length()
	c := Segment{t.C, t.A}.Length()
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}


type Circle struct {
	Center Point
	Radius float64
}


func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}


type Shape interface {
	Area() float64
}


func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}




// Задание 4

func Map(slice []float64, square func(float64) float64) []float64 {
	result := make([]float64, len(slice))
	copy(result, slice)

	for i := range result {
			result[i] = square()
	}

	return result
}






func main() {
	// 1.1 
	fmt.Println(`Задание 1.1: formatIP принимает массив из четырех байт и возвращает строку с IP-адресом в формате 'x.x.x.x'`)
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println(formatIP(ip)) 



	// 1.2
	fmt.Println(`Задание 1.2: listEven принимает два целых числа, и возвращает срез с четными числами в заданном диапазоне и ошибку, если диапазон некорректный`)
	// Пример корректного диапазона
	evenNumbers, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Четные числа:", evenNumbers) 
	}

	// Пример некорректного диапазона
	evenNumbers, err = listEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err) 
	} else {
		fmt.Println("Четные числа:", evenNumbers)
	}



	// 2
	fmt.Println("Задание 2: countCharacters принимает строку и возвращает карту, где каждому символу соответствует количество его вхождений")
	// Пример строки
	s := "hello world"
	result := countCharacters(s) 
	fmt.Println("Количество вхождений каждого символа:", result)



	// 3
	fmt.Println("Задание 3:")
	// Создаем треугольник и круг
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{0, 3},
		C: Point{4, 0},
	}
	circle := Circle{
		Center: Point{0, 0},
		Radius: 5,
	}

	// Выводим площади фигур
	printArea(triangle) 
	printArea(circle)   



// 4 
// 1. Создаем и заполняем срез
values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
fmt.Println("Исходный срез:", values)

// 2. Создаем функцию для возведения в квадрат и присваиваем её переменной
square := func(x float64) float64 {
		return x * x
}


// 3. Применяем функцию Map с функцией square
squaredValues := Map(values, square)

// Выводим результаты
fmt.Println("Срез после применения функции (копия):", squaredValues)
fmt.Println("Исходный срез после Map (остался неизменным):", values)
}