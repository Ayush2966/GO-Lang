package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("hello")
	myDefer()	
}
// world two three
// 0 1 2 3 4
// hello
 func myDefer(){
	for i:= 0; i<5; i++{
		defer fmt.Println(i)
	}
 }