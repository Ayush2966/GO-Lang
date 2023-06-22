package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
		days :=[]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		fmt.Println(days)
		fmt.Println("01 loop")
		for d:=0; d<len(days); d++{
			fmt.Println(days[d])
		}	
		fmt.Println("02 loop")
		for i:= range days{
			fmt.Println(days[i])
		}
		fmt.Println("03 loop")
		for index, days :=range days{
			fmt.Printf("index is %v and value is %v\n",index, days)
		}

		rougueValue := 1
		for rougueValue<10{
			if rougueValue==5{
				break
			}
			fmt.Println("value is :", rougueValue)
			rougueValue++
		}
		
		rouguValue := 1
		for rouguValue<10{
			if rouguValue ==2{
				goto lco
			}
			if rouguValue==5{
				rouguValue++
				continue
			}
			fmt.Println("value is :", rouguValue)
			rouguValue++
		}

		lco:
		fmt.Println("Jumping at ayush jain")
}
