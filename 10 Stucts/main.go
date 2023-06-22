package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	 Ayush := User{ "Ayush Jain", "ajsa02812@gmail.com", true, 20}
	fmt.Println(Ayush)
	fmt.Printf("Ayush details are %+v \n", Ayush)
	fmt.Printf("Name is %v and email is %v ", Ayush.Name, Ayush.Email)

}	

// no inheritance in golang : no super or paren
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

