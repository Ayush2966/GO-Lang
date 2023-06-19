package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter:")
	//comma ok // error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for reading,", input)
	fmt.Printf("type of this reading %T \n", input)
	fmt.Println("enter:")
	//comma ok // error ok
	inpu, _ := reader.ReadString('\n')
	fmt.Println("thanks for reading,", inpu)
	fmt.Printf("type of this reading %T", inpu)

}
