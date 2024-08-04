package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	DecodeJson()
}
type course struct {
	Name     string   
	Price    int
	Platform string   
	Password string   
	Tags     []string 
}

func DecodeJson() {
	jsonData := []byte(`
	{
                "Name": "Reactjs",
                "Price": 299,
                "Platform": "learncode.com",
                "Tags": ["xyz","abc"]
        }
	`)
	var lcocourses course
	checkvalid := json.Valid(jsonData)

	if checkvalid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonData,&lcocourses)
		fmt.Printf("%#v \n", lcocourses)
	}else{
		fmt.Println("Json was not valid")
	}

	//in some cases where you just want to add data to key value pair

	var onlineData map[string] interface{}	//we do not know the coming value is just string or not, it can be a slice too thats why we use interface here
		json.Unmarshal(jsonData,&onlineData)
		fmt.Printf("%#v \n", onlineData)

	for k,v := range onlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", k,v,v)
	}
}