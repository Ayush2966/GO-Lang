package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's learn Slices")
	var fruitlist = []string{"apple", "peach", "pineaple"}
	fmt.Printf("Type of Fruitlist is %T \n", fruitlist)

	fruitlist = append(fruitlist, "mango", "guava")
	fmt.Println(fruitlist)
	fruitlist=append(fruitlist[1:])
	fmt.Println(fruitlist)
	fruitlist=append(fruitlist[:3])
	fmt.Println(fruitlist)
	fruitlist=append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScore := make([]int, 4)
	highScore[0]=234
	highScore[1]=945
	highScore[2]=456
	highScore[3]=567
	highScore= append(highScore, 55, 666 , 555)
	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index
	var courses =[]string{"js", "html", "go", "css", "shell", "pypy"}
	fmt.Println(courses)
	var index int =2
	courses = append(courses[:index], courses[index+1:]...)
    fmt.Println(courses)
}
