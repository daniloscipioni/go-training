package main
import (
	"fmt"
)

// Números sequenciais
const (
	a = iota * 10
	b 
	c 
	_ 
	e 
)

func main() {
	fmt.Println(a, b, c, e)

} 