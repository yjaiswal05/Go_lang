package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of response is %T\n",response)	//gives the http response reference (pointer)
	readReq(*response)


	defer response.Body.Close()	//it is required to close the connection after it has been used
}
func readReq(response http.Response){
	databyte,err := ioutil.ReadAll(response.Body)	
	if err != nil {
		panic(err)
	}
	fmt.Println("Data is :", string(databyte))	

}