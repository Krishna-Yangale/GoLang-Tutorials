package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Welcome to web requests")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

}
