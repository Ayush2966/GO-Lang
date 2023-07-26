package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("hiiii")
	formRequest()
	// postJsonRequest()
}

// func postJsonRequest() {
// 	const myurl = "http://localhost:8000/post"

//		//fake json paload
//		requestBody := strings.NewReader(`
//		{
//			"coursename": "reactjs",
//			"price":1000,
//			"platforrm" : "learnCodeOnline.in"
//		}
//		`)
//		response, err := http.Post(myurl, "application/json", requestBody)
//		if err != nil {
//			panic(err)
//		}
//		defer response.Body.Close()
//		content, _ := ioutil.ReadAll(response.Body)
//		fmt.Println(string(content))
//	}
func formRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "ayush")
	data.Add("lasttname", "jain")
	data.Add("email", "ayush@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
