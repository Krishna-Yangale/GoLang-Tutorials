package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcom to control flow in golang")
	fmt.Println("Enter the number:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	number, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if number%2 == 0 {
		fmt.Println("number is even number")

	} else {
		fmt.Println("number is odd number")
	}

	logincount := 30

	var result string

	if logincount < 10 {
		result = "verified user"
	} else if logincount > 10 {
		result = "something fishy"
	} else {
		result = "it is exactly 10"
	}

	fmt.Println(result)

	//also we can use if else without defining variable

	if 10%2 == 0 {
		fmt.Println("the number is Even")
	} else {
		fmt.Println("the number is odd")
	}

	//or we can intialize variable in single and also check for condition

	if num := 3; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("num is greate than 10")
	}

}
