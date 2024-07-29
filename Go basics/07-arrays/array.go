package main

import "fmt"

func main() {
	var array [4] string	//first way to initialise
	var array2  = [4] string{"a","b","c","d"}	//second way to initialise

	array[0] = "Apple"
	array[1] = "Orange"
	array[3] = "Peach"
	//array[0] = "Grapes"

	fmt.Println(array)		//in the place of missing value it places a blank space and the values are just written without comma
	fmt.Println(len(array))	//len doesn't actually calculate the length it just returns what length you have initialised

	
	fmt.Println(array2)		

}