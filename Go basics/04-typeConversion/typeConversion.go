package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out pizza app")
	fmt.Println("Please rate our Pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Printf("The type of input is : %T \n", input)	// this will give string

	//numRating, _ := strconv.ParseFloat(input, 64)		// this will not work because when we hit enter \n automatically goes to the input and the input function ends
	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input),64)	//we trimmed the space and only took the value

	fmt.Println(numRating)
	fmt.Printf("The type of input is : %T \n", numRating)	// this will give float as we used ParseFloat

}