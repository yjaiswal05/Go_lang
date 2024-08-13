package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

type Myevent struct{
	Name string 	`json: "what is you name?"`
	Age int 		`json:"How old are you?"`
}

type MyResponse struct{
	Message string	`json:"Answer"`
}

func HandleLambdaEvent(event Myevent)(MyResponse, error)  {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!",event.Name,event.Age)},nil
}