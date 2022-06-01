package main

import (
	// "crypto/rand"
	"fmt"
	"time"

	"math/rand"
)

func main() {

	//by using math/random package
	fmt.Printf("Generate a random number:")
	rand.Seed(time.Now().UnixNano())
	input1 := rand.Intn(50)
	input2 := rand.Intn(20)
	fmt.Println(input1)
	fmt.Println(input2)

	//operation on random numbers
	// 1.Addition
	add := input1 + input2
	fmt.Println("Addition is:", add)

	//2.subtraction
	sub := input1 - input2
	fmt.Println("Subtraction is:", sub)

	//3.multiplication
	mul := input1 * input2
	fmt.Println("Multiplication is:", mul)

	//4.division
	div := float64(input1) / float64(input2)
	fmt.Println("Division is :", div)

	//by using crpto/random package

	// myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(myRandomNum)

}
