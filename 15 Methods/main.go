package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	 Ayush := User{ "Ayush Jain", "ajsa02812@gmail.com", true, 20}
	fmt.Println(Ayush)
	fmt.Printf("Ayush details are %+v \n", Ayush)
	fmt.Printf("Name is %v and email is %v \n ", Ayush.Name, Ayush.Email)
	Ayush.GetStatus()
	Ayush.Newmail()
	fmt.Printf("Name is %v and email is %v \n ", Ayush.Name, Ayush.Email)
}
// no inheritance in golang : no super or paren
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int	
	}


	func (u User) GetStatus(){
		fmt.Println("Is user active:", u.Status)

	}

	func (u User) Newmail(){
		u.Email="test@go.dev"
		fmt.Println("Email of this user is: ", u.Email)
	}
