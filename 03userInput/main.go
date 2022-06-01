package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the first number:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Print(input)

	number1, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	fmt.Println(number1)

	fmt.Print("Enter the second number:")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	// fmt.Print(input)

	number2, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)

	fmt.Println(number2)

	sum := number1 + number2

	fmt.Println(sum)

	var userName string
	fmt.Println("Enter your name")
	fmt.Scan(&userName)
	fmt.Println("Username is ", userName)

	var number, number3 int
	fmt.Printf("Enter the first number:")
	fmt.Scan(&number, &number3)

	add := number + number3
	fmt.Println(add)

}
