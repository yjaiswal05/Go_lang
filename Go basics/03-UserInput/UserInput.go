package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)		// we took the shortcut as we don't know what type is the user going to pass
	fmt.Println("Enter the rating for this lang:")

	//comma ok syntax or error ok syntax
	 input,_ := reader.ReadString('\n')	// since go lang doesn't have try and catch it uses comma ok where input tries and err catches the eror
	 println("Thanks for reading", input)	//we gave "\n" to tell the compiler when to stop reading the input
}