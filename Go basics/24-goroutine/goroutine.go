package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup
func main() {
	//go greeter("hello")		// go routine is same as threads in java, it is a part of concurrency
	//greeter("world")
	
	websiteList := [] string{
		"https://google.com",
		"https://facebook.com",
	}	
	for _,web := range websiteList{
		go getStatusCode(web)
		wg.Add(1)
	}	
	wg.Wait()
}
		

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("error in endpoint")
	}else{
	fmt.Printf("%d status code for %s\n",res.StatusCode,endpoint)
}
}