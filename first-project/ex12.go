package main
import (
	"fmt"
)

func main() {

	valor := 5	
	switch {
		
	case valor < 4:
		fmt.Println("menor que 4")
	case valor < 6:
		fmt.Println("menor que 6")	
	}
} 