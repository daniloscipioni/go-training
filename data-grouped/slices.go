package main

import "fmt"


func main()  {

	slice:=[]int{1,2,3,4,900}
	total := 0
	for index,valor:=range slice {
		total += valor
		fmt.Println(index, valor)
	}
	fmt.Println(total)
}