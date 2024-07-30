package main

import "fmt"
//function inside function is not allowed

func greeter() {
	fmt.Println("Hello")
}
func adder(i int, j int ) int {		
	return (i+j)
}
func proadder(l... int) int {	
	total := 0
	for val := range l{
		total += val
	}	
	return (total)
}
func main()  {
	greeter()
	fmt.Println(adder(1,2))
	fmt.Println(proadder(1,2,3,4,5,6,7))

}