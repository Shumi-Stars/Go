package main

import (
	"errors"
	"fmt"
)

// hello принимает строку и возвращает строку "Привет, name!"
func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

// printEven принимает два целых числа и выводит четные числа в этом диапазоне
func printEven(start int, end int) error {
	if start > end {
		return errors.New("левая граница больше правой")
	}
	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

// apply выполняет математическое действие над двумя числами
func apply(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("действие не поддерживается")
	}
}

func main() {
	// Тесты для функции hello
	fmt.Println(hello("Go"))
	fmt.Println(hello("world"))

	// Тесты для функции printEven
	if err := printEven(1, 10); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := printEven(10, 1); err != nil {
		fmt.Println("Ошибка:", err)
	}

	// Тесты для функции apply
	result1, err1 := apply(3, 5, "+")
	if err1 != nil {
		fmt.Println("Ошибка:", err1)
	} else {
		fmt.Println("Результат:", result1)
	}

	result2, err2 := apply(7, 10, "*")
	if err2 != nil {
		fmt.Println("Ошибка:", err2)
	} else {
		fmt.Println("Результат:", result2)
	}

	result3, err3 := apply(3, 5, "/")
	if err3 != nil {
		fmt.Println("Ошибка:", err3)
	} else {
		fmt.Println("Результат:", result3)
	}

	result4, err4 := apply(3, 5, "#")
	if err4 != nil {
		fmt.Println("Ошибка:", err4)
	} else {
		fmt.Println("Результат:", result4)
	}
}
