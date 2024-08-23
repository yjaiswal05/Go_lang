package main

import (
	"encoding/json"
	"fmt"	
	"io"
	"net/http"
	"os"
	"strconv"


)

func main() {
	os.Setenv("PexelsToken", "")
	var TOKEN = os.Getenv("PexelsToken")
	
	var c = NewClient(TOKEN)
	result, err := c.searchPhotos("sky",5,1)
	if err != nil {
		fmt.Errorf("Search error:%v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}
	fmt.Println(result)
}

const (
	PhotoApi = ""
	VideoApi = ""
)

type Client struct{
	Token 			string
	hc 				http.Client
	RemainingTimes 	int32
}

func NewClient(token string) *Client{
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type SeachResult struct{
	Page 			int32
	PerPage 		int32
	TotalResult 	int32
	NextPage 		string
	Photos 			[]Photo
}

type Photo struct{
	Id				int32
	Width			int32
	Height			int32
	Url				string
	Photographer	string
	Src				PhotoSource
}

type PhotoSource struct{
	Original		string
	Large			string
	Medium			string
	Small			string
	Portrait		string
}

func (c *Client) searchPhotos(query string, perPage, page int) (*SeachResult,error)  {
	url := fmt.Sprintf(PhotoApi+"/search?query=%s&per_page=%d&page=%d",query,perPage,page)
	resp , err := c.requestDoWithAuth("GET",url)
	defer resp.Body.Close()
	data,err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var result SeachResult
	err = json.Unmarshal(data, &result)
	return &result,err
}

func (c *Client) requestDoWithAuth(method, url string) (*http.Response,error)  {
	req,err := http.NewRequest(method,url,nil)
	if err != nil {
		return nil,err
	}
	req.Header.Add("Autorization",c.Token)
	resp,err := c.hc.Do(req)
	if err != nil {
		return resp,err
	}
	times, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return resp,nil
	}else{
		c.RemainingTimes = int32(times)
	}
	return resp,nil
}

