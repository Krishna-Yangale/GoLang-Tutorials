package main

import "fmt"

func main() {

	//different ways to print data using struct
	fmt.Println("Welcome to strcut in golang")
	krishna := User{"Krishna", "krishna.yangale@radiant.co", true, 24}
	fmt.Println(krishna)

	fmt.Printf("The details are :%+v\n", krishna)
	fmt.Printf("The name is %v and email is %v", krishna.Name, krishna.Email)

}

//defining struct in golang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
