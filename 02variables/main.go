package main

import "fmt"

func main() {
	var firstName string = "krishna"
	var empId int = 101

	fmt.Println("My firstname is ", firstName)
	fmt.Printf("The tpe of firstname is %T\n", firstName)
	fmt.Println("My employee Id is ", empId)
	fmt.Printf("The type of employee Id %T\n", empId)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("the type of this variable is %T\n", isLoggedin)

	lastName := "krishna"
	fmt.Println(lastName)
	fmt.Printf("The type of lastname is %T \n", lastName)

	const mobileNumber int = 8169285583
	fmt.Println(mobileNumber)
	fmt.Printf("The type of variale is %T", mobileNumber)

}
