package main

import "fmt"


func funcVariadic(a ...int) int  {
	soma := 0
	for _, v := range a {
		soma += v
	}

	return soma
}

func funcVariadic1(a []int) int  {
	soma := 0
	for _, v := range a {
		soma += v
	}

	return soma
}


func main()  {
	x := []int{1,1,1,1,1,1,1,1,1,1}
	y := []int{1,1,1,1,1,1,1,2,1,1}

	fmt.Println(funcVariadic(x...))
	fmt.Println(funcVariadic1(y))
}