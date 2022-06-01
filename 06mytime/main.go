package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package")

	parseTime := time.Now()
	fmt.Println(parseTime)
	fmt.Println(parseTime.Format("01-02-2006 15:30:15 Monday"))

	createDate := time.Date(2020, time.July, 23, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:30:12 Monday"))

}
