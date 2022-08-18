package main
import (
	"fmt"
)
t := 5
type test int

var (
	y int
	x test
)

func main() {
	fmt.Println(t)
	fmt.Printf("%v\t%T\n" , x, x)
	x = 42
	fmt.Printf("%v\n" , x)
	y = int(x)
	fmt.Printf("%v\t%T\n" , y, y)
} 