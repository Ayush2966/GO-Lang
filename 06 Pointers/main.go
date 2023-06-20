package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var myptr = &myNumber
	fmt.Println("value of actual pointer is ", myptr)
	fmt.Println("value of actual pointer is ", *myptr)
     *myptr= *myptr * 2
	 fmt.Println("new value is", myNumber)
}