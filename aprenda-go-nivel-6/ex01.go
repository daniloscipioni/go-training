package main

import "fmt"


func returnAnInt() int  {
	return 10
}

func returnAnIntAndString() (int, string) {
	return 10, "Bom dia"
}

func main()  {
	
	fmt.Println(returnAnInt())
	fmt.Println(returnAnIntAndString())

}