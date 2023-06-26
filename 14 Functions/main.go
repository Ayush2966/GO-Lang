package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greet()
	greet2()
	result := adder(3,5)
	fmt.Println("result is:",result)
}
func adder(valone int, valtwo int)int{
	return valone + valtwo

}
func greet() {
	fmt.Println("Hello, Ayush this side!!")
}
func greet2(){
	fmt.Println("Another method")
}