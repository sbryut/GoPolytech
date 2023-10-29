package main

import (
	"fmt"
)

func main() {
	var op1, op2 float64
	var sign string
	var tmp string

	for {
		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scanln(&op1); err != nil {
			fmt.Scanln(&tmp)
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			continue
		}
		break
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scanln(&sign)

	for {
		fmt.Print("Введите второе число: ")
		if _, err := fmt.Scanln(&op2); err != nil {
			fmt.Scanln(&tmp)
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			continue
		}
		break
	}

	switch sign {
	case "+":
		fmt.Printf("Результат: %.2f\n", op1+op2)
	case "-":
		fmt.Printf("Результат: %.2f\n", op1-op2)
	case "*":
		fmt.Printf("Результат: %.2f\n", op1*op2)
	case "/":
		if op2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно!")
			return
		}
		fmt.Printf("Результат: %.15f\n", op1/op2)
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
	}
}
