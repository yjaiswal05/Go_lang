package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"` //this json thing is not mandatory
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}
func EncodeJson(){
	lcocourses := []course{
		{"Reactjs", 299, "learncode.com", "abcs123",[]string{"xyz","abc"}},
		{"Nodejs", 99, "learncode.com", "bcs123",[]string{"dfd","abtdc"}},
		{"Java", 199, "learncode.com", "abc23",[]string{"xyhyhz","abhtdc"}},
	}

	finaljson,err := json.MarshalIndent(lcocourses,"","\t")
	if err!= nil {
		panic(err)
	}
	fmt.Printf("%s\n",finaljson)
}