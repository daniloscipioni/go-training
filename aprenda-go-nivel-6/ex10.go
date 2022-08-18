package main

import "fmt"

func main()  {
	x := returnAFunction()
	z := returnAFunction()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println("---------")
	fmt.Println(z())
}

// Retornando uma função
func returnAFunction() func() int  {
	// Clojure - Usa a variavel dentro do escopo e cada funcão chamada
	x := 0
	return func() int {
		x++
		return x
	}
}