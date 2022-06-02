package main

import "fmt"

func main() {
	fmt.Println("welcome to maps in Golang")

	languages := make(map[string]string)

	languages["js"] = "javaScript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	fmt.Println(languages)
	//printing values from key in map
	fmt.Println("js is short for ", languages["js"])

	//deleting value from key
	delete(languages, "rb")
	fmt.Println(languages)

	//loop for map

	for key, value := range languages {
		fmt.Printf("For key %v,the value is %v\n", key, value)
	}
}
