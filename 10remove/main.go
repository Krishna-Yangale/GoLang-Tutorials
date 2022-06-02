package main

import "fmt"

func main() {
	fmt.Println("Welcome to learn about how to remove slices")

	var courses = []string{"react", "python", "java", "ruby", "cloud"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
