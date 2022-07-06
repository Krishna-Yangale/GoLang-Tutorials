package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=12345"

func main() {
	fmt.Println("welcome to handling urls")
	fmt.Println("url is:", myurl)

	//parsing

	url.Parse(myurl)
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("the type of qparams is %T\n", qparams)

	fmt.Println(qparams["coursename"])
	for _, val := range qparams {
		fmt.Println("param is:", val)

	}
	partsofurl := &url.URL{
		Scheme: "http",
		Host:   "lco.dev",
		Path:   "/learn",
	}

	anotherurl := partsofurl.String()
	fmt.Println("another url is:", anotherurl)
}
