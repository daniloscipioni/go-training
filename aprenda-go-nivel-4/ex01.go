package main

import "fmt"

func main()  {

	firstArray := [5]int{1,2,3,4,5}

	for index, value := range firstArray {
		
		fmt.Println(index, value)
	}

	fmt.Printf("Type: %T  \n", firstArray)
	
}