package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time")

	presentTime := time.Now()
	fmt.Println(presentTime)

	//changing the format of date and time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))	// this is the standard to format date & time and can't be changed

	//custom date 
	createdDate := time.Date(2020, time.November, 5, 10, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}