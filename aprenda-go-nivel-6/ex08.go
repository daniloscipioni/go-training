package main

import "fmt"

func main()  {
	x := returnAFunction()
	fmt.Println(x)
	y := x(3)

	z := returnAFunction()(8)

	fmt.Println(y)
	fmt.Println(z)
}

// Retornando uma função
func returnAFunction() func(int) int  {
	return func (i int) int {
		return i * 10
	}
}