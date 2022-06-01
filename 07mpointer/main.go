package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	number := 56
	var ptr = &number
	fmt.Println("The value of pointer is ", ptr)
	fmt.Println("The value of pointer is ", *ptr)

	*ptr = *ptr + 6
	fmt.Println("the value of the pointer is ", number)

	firstName := "krishna"
	fmt.Println(firstName)

	var ptr1 = &firstName
	fmt.Println("Firstname is ", *ptr1)

	*ptr1 = *ptr1 + " " + "Yangale"
	fmt.Println(*ptr1)

}
