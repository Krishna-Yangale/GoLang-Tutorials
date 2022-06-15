package main

import (
	"fmt"
)

func main() {

	user := bodmas{15, 10, 8, 7, 6}

	res := user.calculateEquation()

	fmt.Println(res)

}

type bodmas struct {
	number1 float32
	number2 float32
	number3 float32
	number4 float32
	number5 float32
}

func (b bodmas) calculateEquation() float32 {
	equation := (b.number1 + b.number2 - b.number3) * b.number4 / b.number5

	return equation
}
