package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Golang lets practice Slices")
	var myFruitList = []string{"apple", "mango", "banana", "peach", "watermelon"}

	fmt.Println(myFruitList)
	fmt.Printf("The type of my fruitlist is %T\n", myFruitList)
	fmt.Println(len(myFruitList))

	//Adding items to slices without defining the size using append method
	myFruitList = append(myFruitList, "Oranges", "Berries", "Blackberries")
	fmt.Println(myFruitList)
	fmt.Println(len(myFruitList))

	//printing or sliceing the given array in the range
	myFruitList = append(myFruitList[2:5])
	fmt.Println(myFruitList)

	//using make for memory management

	highScores := make([]int, 4)

	highScores[0] = 290
	highScores[1] = 310
	highScores[2] = 355
	highScores[3] = 458
	// highScores[4] = 555
	fmt.Println(highScores)

	highScores = append(highScores, 222, 980, 900)
	fmt.Println(highScores)
	fmt.Println(len(highScores))
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
