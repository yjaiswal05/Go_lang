package main

import "fmt"

func main()  {
	defer fmt.Println("World")	//this should run first but since we have used defer it will execute just before the function is going to end
	defer fmt.Println("One")	//It works on LIFO principle
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i<5; i++{
		defer fmt.Println(i)
	}
}
