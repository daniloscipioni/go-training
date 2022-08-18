package main

import "fmt"

func main()  {
	x := 4
	// função na variavel
	funcvar := func (y int) int  {
		return y * 2
	}

	fmt.Println(x, "vezes 2 =", funcvar(x))
}