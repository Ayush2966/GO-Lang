package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitlist [4]string
	fruitlist[0]="apple"
	fruitlist[1]="banana"
	fruitlist[3]="peach"

	fmt.Println("Fruit list is:", fruitlist)
	fmt.Println("Fruit list is:", len(fruitlist))
	
	var veglist =[3]string{"potato", "tommato", "ladyfinger"}
	fmt.Println("vegy list is", veglist)
	fmt.Println("vegy list is", len(veglist))
}
