package main

import "fmt"

func main() {
	fmt.Println("If-Else in Golang")


	loginCount :=23
	var result string 
	if loginCount<10{
	result="Regular user"
	} else if loginCount>10 {
		result="Ghar ka Admi"
	}else{
		result="huehue"
	}
	fmt.Println(result)
	
	if 9%2 == 0{
		fmt.Println("Number is even")
	} else{
		fmt.Println("Number is odd")
	}

	if num:=3; num<10{
		fmt.Println("huehue")
	}else{
		fmt.Println("not huehue")
	}
}
