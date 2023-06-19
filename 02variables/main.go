package main

import "fmt"

func main() {
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	 
	var smallFloat float32= 255.56464545454455455
	fmt.Println(smallFloat)
	fmt.Printf("variable of type: %T \n", smallFloat)

	var smallyFloat float64= 255.56464545454455455
	fmt.Println(smallyFloat)
	fmt.Printf("variable of type: %T \n", smallyFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("varaible is of stype: %T \n", anotherVariable)

	//implicit type
	var website ="ayush jain"
	fmt.Println(website)
	
	//no var style
	 numberOfUser :=300000.0
	 fmt.Println(numberOfUser)
}
