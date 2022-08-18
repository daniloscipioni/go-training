package main

import "fmt"


func main()  {

	slice:=[]int{1,2,3,4,900}
	slice2:=slice[0:3]
	// Precisa transformar o slice em itens do tipo do slice
	slice3:=append(slice, slice2...)

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}