package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	
	content := "This is going inside the file"

	file, err := os.Create("./myfile.txt")	//to create a file we use os package
	checkNilError(err)
	
	length , err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is :", length)
	readFile("./myfile.txt")
	defer file.Close()
}

func readFile(filename string){
	databyte,err := ioutil.ReadFile(filename)	//to read a file we use ioutil package
	checkNilError(err)
	fmt.Println("Data is :", databyte)
	fmt.Println("Data is :", string(databyte))	//the raw data is in the form of bytes

}

func checkNilError(err error){
	if err != nil {
		panic(err)		//panic will stop the execution of the program and will return you the error
	}
}
