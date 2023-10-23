package main

import (
	"fmt"
	"os"
)

func main() {
	var firstNumber float64
	var secondNumber float64
	var operation string

	fmt.Print("Введите первое число: ")

	if _, err := fmt.Fscan(os.Stdin, &firstNumber); err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	fmt.Print("Введите операцию (+, -, *, /): ")

	if _, err := fmt.Fscan(os.Stdin, &operation); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	switch operation {
	case "+", "-", "*", "/":
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		return
	}

	fmt.Print("Введите второе число: ")

	if _, err := fmt.Fscan(os.Stdin, &secondNumber); err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	if secondNumber == 0 {
		fmt.Println("Ошибка: деление на ноль невозможно.")
		return
	}

	var result float64
	switch operation {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		result = firstNumber / secondNumber
	}

	fmt.Printf("Результат: %.2f\n", result)
}
