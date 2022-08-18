package main

import "fmt"

func main()  {
	// Callback
	functionReturn := somaGeral(soma, 32)
	fmt.Println(functionReturn)
}

func soma(i,j int) int  {
	return i + j
}

func somaGeral(callback func(i,j int) int, valor int) int  {
	return callback(1,9) + valor
}