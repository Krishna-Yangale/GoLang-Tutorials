package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// var a [3]int
	// a[0] = 1
	// a[1] = 3
	// a[2] = 5
	// a := [4]int{1, 2 , 5}

	// fmt.Println(a)

	var a [4]string
	a[0] = "apple"
	a[3] = "Orange"

	// a := [4]string{"apple", "orange"}
	fmt.Println(len(a))
	fmt.Println(a)

}
