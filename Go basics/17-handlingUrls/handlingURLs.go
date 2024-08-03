package main
import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev"

func main() {

	//Parsing
	result,_ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()		//query parameter
	fmt.Println(qparams)
}