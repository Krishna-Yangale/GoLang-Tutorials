package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	res := adder(10, 5)
	fmt.Println(res)

	res2, message := proAdder(1, 5, 6, 9, 8)
	fmt.Println(res2, " ", message)

}

func adder(value1, value2 int) int {
	result := value1 + value2
	return result
}

func proAdder(values ...int) (int, string) {
	totalRes := 0
	for _, value := range values {
		totalRes += value

	}
	return totalRes, "is pro result is "

}
