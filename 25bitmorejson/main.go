package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcometo json lesson")
	// EncodeJson()
	DecodeJson()

}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// func EncodeJson() {
// 	learnCourses := []course{
// 		{"ReactBootcamp", 100, "LCO", "123456", []string{"web-dev", "js"}},
// 		{"GoBootcamp", 200, "LCO", "123456", []string{"web-dev", "go"}},
// 		{"PythonBootcamp", 300, "LCO", "123456", []string{"web-dev", "python"}}}

// 	//packaga this data into a json string

// 	jsonData, err := json.MarshalIndent(learnCourses, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", jsonData)
// }

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "ReactBootcamp",
			"Price": 100,
			"website": "LCO",
			"tags": ["web-dev","js"]
		}
		
	`)
	var lcoCourse course

	checkvalid := json.Valid(jsonData)

	if checkvalid {
		fmt.Println("json data is valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("json data is not valid")
	}

	var myonlineData map[string]interface{}

	json.Unmarshal(jsonData, &myonlineData)
	fmt.Printf("%#v\n", myonlineData)

	for k, v := range myonlineData {
		fmt.Printf("Key is %v and value is %v and type is :%T\n", k, v, v)
	}
}
