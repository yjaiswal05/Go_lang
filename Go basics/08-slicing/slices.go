package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Mango"}	//it is like an arraylist
	fruitList = append(fruitList, "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highscores := make([]int, 4)
	highscores[0] = 234
	highscores[1] = 345
	highscores[2] = 456
	highscores[3] = 678
	//highscores[4] = 789		//this will give index out of reange error

	highscores = append(highscores, 963,852,741)	// this will append the new values and increase the size doesn't matter what was the initial size
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)

	//removing element from a slice
	var courses = [] string {"react","ruby","python","java"}
	var index int = 2
	fmt.Println(append(courses[:index],courses[index+1:]...))	//we don't have a direct function so we print the rest of the slice excluding the index position one

}