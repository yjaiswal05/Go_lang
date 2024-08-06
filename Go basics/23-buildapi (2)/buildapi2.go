package main

import (
	"buildapi2/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":9098",r))
	fmt.Println("Listening at port 9098")

}