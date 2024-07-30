package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["ENG"] = "English"
	languages["HIN"] = "Hindi"
	languages["MATH"] = "Mathematics"
	languages["PHY"] = "Physics"

	fmt.Println(languages["ENG"])	//retrieving values based on key

	delete(languages,"PHY")
	fmt.Println(languages)	//deleting values based on key

}