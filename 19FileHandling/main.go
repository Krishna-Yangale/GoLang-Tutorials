package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling")
	content := "This is course about Golang to learn about more go to its original website and read documentation"
	file, err := os.Create("./myfile.txt")
	checkNilErr(err)

	// if err != nil {
	// 	panic(err)
	// }
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("the length is ", length)
	defer file.Close()
	redaFile("./myfile.txt")

}

func redaFile(file string) {
	databyte, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("the file databyte is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
