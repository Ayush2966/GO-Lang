package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greet()
	greet2()
	result := adder(3,5)
	fmt.Println("result is:",result)

	pror := proAdder(2,3,5,6,4)
	fmt.Println("pro resullt is:",pror)
}


func adder(valone int, valtwo int)int{
	return valone + valtwo
}


func proAdder(values...int)int{
	total :=0

	for _, val:= range values{
		total +=val		
	}
	return total
}


func greet() {
	fmt.Println("Hello, Ayush this side!!")
}


func greet2(){
	fmt.Println("Another method")
}