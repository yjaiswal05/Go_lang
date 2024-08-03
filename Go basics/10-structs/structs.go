package main

import "fmt"

func main() {
	// there is no inheritance in go and we don't have super or parent
	u1 := User{"Yash", "y@j.com", true, 22}
	fmt.Println(u1)
	u1.GetStatus()
}

type User struct {  //instead of having a class, we have structs
	Name   string	//the thinngs declared in uppercase is Public 
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)
}