package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PostReqJson()	//this is post request uses json payload
	PostReqForm()	//this is post request uses form payload
}

func PostReqJson()  {
	const myurl = "https://lco.dev" 	//this could also be your localhost url if you have a custom server running
	requestBody :=strings.NewReader(`
		{
		"coursename" : "Lets go with golang"		
		"price" : "0"
		"platform" : "https://xyz.com"
		}
		
		`)		

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	
}

func PostReqForm()  {
	const myurl = "https://lco.dev" 
	data := url.Values{}
		data.Add("firstname","yash")
		data.Add("lastname","jaiswal")
		data.Add("email","e@mail.com")

	response, err := http.PostForm(myurl,data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content,_ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	
}