package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var one *int		//* defines that one is going to be a pointer which is storing a int type value
						// the default value of a pointer is nil if its not initialised

	fmt.Println("Value of pointer is", one)

	myNum := 23
	var ptr = &myNum	// here we are referencing the pre initialised value
	fmt.Println("Value of pointer is", ptr)		//this gives the actual memory location
	fmt.Println("Value of pointer is", *ptr)	//this gives the value

	*ptr = *ptr + 2
	fmt.Println("The new value of pointer is", myNum)	//pointer guarantees that the operations are being performed on the original value not to its copy


}