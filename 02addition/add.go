package main

import "fmt"

func main() {
	fmt.Println("addition is :", 10+9)

	var number1 int = 56
	var number2 int = 69
	var number3 float64 = 69.5

	fmt.Println("The addition is :", number1+number2+int(number3))
	fmt.Println("the Addition is :", float64(number1)+float64(number2)+number3)
}
