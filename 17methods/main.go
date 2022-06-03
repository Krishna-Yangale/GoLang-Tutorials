package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")

	krishna := User{"Krishna", "krishna.yangale@radiantcloud.co", true, 24}
	fmt.Printf("the name is %v and the email is %v\n", krishna.Name, krishna.Email)

	krishna.NewMail()

	input := Area{100, 200}

	fmt.Println(input.calculateArea(100, 200))

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
type Area struct {
	Length  int
	Breadth int
}

func (u User) NewMail() {
	u.Email = "krishna.yangale@radint.co"
	fmt.Println(u.Email)
}

func (a Area) calculateArea(l, b int) int {

	area := a.Length * a.Breadth

	return area
}
