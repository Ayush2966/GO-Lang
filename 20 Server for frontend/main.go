package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("hey")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)
	fmt.Println("Byecount is :", bytecount)
	fmt.Println(responseString.String())
	// fmt.Println(content)
	// fmt.Println(string(content))
}