package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "this needs to go in a file by Ayush Jain"
	file, err :=os.Create("./ayushfile.txt")
	if err !=nil{
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err !=nil{
		panic(err)
	}
		fmt.Println("length is:", length)
		defer file.Close()
		readFile("./ayushfile.txt")
	
}
 func readFile (filename string){
	databyte, err := ioutil.ReadFile(filename)
	if err !=nil{
		panic(err)
	}
	fmt.Println("text data inside the file is \n", string(databyte))
 }
  func checkError (err error){
	if err!=nil{
		panic(err)
	}
  }