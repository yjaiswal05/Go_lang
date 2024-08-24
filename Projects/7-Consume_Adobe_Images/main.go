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
	os.Setenv("PexelsToken", "s1cmhRgBCgxMS6tWPP5q5GfTtxXVjX22aPNIwoA1CDlR7cmaXwn63aQb")
	var TOKEN = os.Getenv("PexelsToken")
	
	var c = NewClient(TOKEN)
	result, err := c.SearchPhotos("sky",5,1)
	if err != nil {
		fmt.Errorf("Search error:%v", err)
	}
	if result.Page == 0 {
		fmt.Errorf("search result wrong")
	}
	fmt.Println(result)
}

const (
	PhotoApi = "https://api.pexels.com/v1/"
	//VideoApi = "https://api.pexels.com/videos/"
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
	Page 			int32			`json:"page"`
	PerPage 		int32			`json:"per_page"`
	TotalResults 	int32			`json:"total_Results"`
	NextPage 		string			`json:"next_page"`
	Photos 			[]Photo			`json:"photos"`
}

type Photo struct{
	Id				int32			`json:"id"`
	Width			int32			`json:"width"`
	Height			int32			`json:"height"`
	Url				string			`json:"url"`
	Photographer	string			`json:"photographer"`
	Src				PhotoSource		`json:"src"`
}

type PhotoSource struct{
	Original		string			`json:"original"`
	Large			string			`json:"large"`
	Medium			string			`json:"medium"`
	Small			string			`json:"small"`
	Portrait		string			`json:"portrait"`
}
type CuratedResult struct{
	Page			int32			`json:"page"`
	PerPage			int32			`json:"per_page"`
	NextPage		int32			`json:"next_page"`
	Photos			[]Photo			`json:"photos"`
}
func (c *Client) SearchPhotos(query string, perPage, page int) (*SeachResult,error)  {
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

func (c *Client) CuratedPhotos(perPage, page int) (*CuratedResult,error)  {
	url := fmt.Sprintf(PhotoApi+"/curated?per_page=%d&page=%d",perPage,page)
	resp , err := c.requestDoWithAuth("GET",url)
	defer resp.Body.Close()
	data,err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var result CuratedResult
	err = json.Unmarshal(data, &result)
	return &result,err
}