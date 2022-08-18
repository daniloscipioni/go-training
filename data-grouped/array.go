package main

import "fmt"

var x [5]int
var y [5]int
func main()  {
	x[0] = 5
	x[1] = 6

	y[0] = 2
	y[1] = 3

	fmt.Println(y)
}