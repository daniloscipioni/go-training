package main

import "fmt"

func main() {
	soma := Soma(1, 2, 3, 4)
	mult := Multiplicacao(1, 2, 3, 4)

	fmt.Println(soma, mult)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func Multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}

	return total
}
