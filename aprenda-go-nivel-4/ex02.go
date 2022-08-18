package main

import "fmt"

func main()  {

	firstSlice := []int{1,2,3,4,5,6,7,8,9,105}
	for index, value := range firstSlice {
		
		fmt.Println(index, value)
	}

	fmt.Printf("Type: %T  \n", firstSlice)
	
}