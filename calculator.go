package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Foydalanuvchidan kiritish
	fmt.Print("Birinchi sonni kiriting: ")
	fmt.Scanln(&num1)
	fmt.Print("Ikkinchi sonni kiriting: ")
	fmt.Scanln(&num2)
	fmt.Print("Amaliyotni kiriting (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Amaliyotga qarab natijani hisoblash
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Xatolik: Nolga bo'lish mumkin emas!")
		}
	default:
		fmt.Println("Xatolik: Noto'g'ri amaliyot kiritildi!")
	}
}
