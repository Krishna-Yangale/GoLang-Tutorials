package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch cases")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("The roll is one you can open")
	case 2:
		fmt.Println("The roll is 2 you ccan move two step forward")
	case 3:
		fmt.Println("You can move three steps further")

	case 4:
		fmt.Println("You can move four steps forward")
	case 5:
		fmt.Println("You can move five steps forward")
	case 6:
		fmt.Println("You can move 6 steps forward and can roll back again")
	default:
		fmt.Println("I dint understad")
	}

}
