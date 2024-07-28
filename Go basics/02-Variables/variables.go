package main

import "fmt"

const LoginToken string = "gibberish"	//if we are using a const with capital letter it automatically becomes Public

func main() {
	var name string = "Yash"
	fmt.Println(name)
	fmt.Printf("Variable is of type : %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)			// values till 0 to 255 are stored in uint8 (unsigned) or we can say it stores 8 bit integer from 0 to 255
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.9684687469874
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)	//has a predefined value of 0
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)		//has no predefined value
	fmt.Printf("Variable is of type : %T \n", anotherString)

	//implicit way of defining a varible
	var word = "Yo"
	fmt.Println(word)				//here the lexers defines the datatype on its own just like in python
	fmt.Printf("Variable is of type : %T \n", word)

	//no var usage
	something := 985				//this := is known as walarus operator and it can be used only inside a method, it can't be used in global scope
	fmt.Println(something)				
	fmt.Printf("Variable is of type : %T \n", something)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

	var a,b int = 2,3				//multiple variable declarance at once
	fmt.Println(a,b)
}