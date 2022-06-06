package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)

	}
}

//basically what defer does is it execute at the end of the main function,if there are multiple defers it execute in LIFO
