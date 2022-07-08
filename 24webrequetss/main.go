package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb lesson")
	PerformGetRequest()
	PerformPostRequest()

	PerformPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Reponse status code:", response.StatusCode)

	fmt.Println("content length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
func PerformPostRequest() {

	const myurl = "http://localhost:3000/post"
	requestBody := strings.NewReader(`
		{
			"coursename": "Learn Go",
			"price": "100",
			"platform": "LCO"
		}
	`)

	response, err := http.Post("http://localhost:3000/post", "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	//dataform

	data := url.Values{}

	data.Add("coursename", "Learn Go")
	data.Add("price", "100")
	data.Add("platform", "LCO")
	data.Add("name", "John")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
