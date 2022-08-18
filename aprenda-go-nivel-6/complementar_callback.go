package main

import "fmt"

func main()  {
	t := somaSomenteImpares(soma, []int{1,2,3}...)
	fmt.Println(t)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	
	return n
}

func somaSomenteImpares(f func(x ...int) int, y ...int) int  {
	var slice []int

	for _, v := range y {
		if(v%2 != 0) {
			slice = append(slice, v)
		}
	}

	total := f(slice...)

	return total

}