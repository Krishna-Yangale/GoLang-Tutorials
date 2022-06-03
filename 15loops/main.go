package main

import "fmt"

func main() {
	fmt.Println("Welcome to the loops in ")

	days := []string{"Sunday", "Monday", "Tuesday", "Thursday", "Friday", "Saturday"}
	// fmt.Println(days)

	//one of the way to use for loop traditional way
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])

	// }

	// for i := range days {
	// 	fmt.Println(days[i])

	// }

	for index, values := range days {
		fmt.Printf("At the index %v ,the day is %v\n", index, values)
	}

}
