package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	GetReq()
}

func GetReq()  {
	const myurl = "https://lco.dev" 	//this could also be your localhost url if you have a custom server running
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("The type of response is %T\n",response)
	fmt.Printf("The response code is %T\n",response.StatusCode)
	fmt.Println(response)

	var responseString strings.Builder
	content,_ := ioutil.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)

	fmt.Println("Byte count is : ", byteCount)	
	fmt.Println(responseString.String())	//	can also be written directly as fmt.Println(string(content))

}