package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	languages := make(map[string]string)
	languages["JS"]= "Javascript"
	languages["GO"]= "GO"
	languages["PY"]="python"
	fmt.Println("List of all languages", languages)
	fmt.Println("Js shorts for", languages["JS"])
	delete(languages, "GO")
	fmt.Println("List of all langauges", languages)

	// loops are interesting in go
	for key, value := range languages{
		fmt.Printf("For key %v, value is %v\n",key ,value)
	}

}
