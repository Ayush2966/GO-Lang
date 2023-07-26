package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to handling url in golang")
	fmt.Println(myurl)
	//parsing-- extracting info 
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of Query params are %T \n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val:= range qparams{
		fmt.Println("Param is:", val)
	}

	partofUrl := &url.URL{
		Scheme: "https",
		Host : "lco.dev",
		Path : "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURl :=partofUrl.String()	
	fmt.Println(anotherURl)
}
