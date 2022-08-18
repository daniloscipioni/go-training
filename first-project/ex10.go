package main
import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("startNumber = %v  mod 4 = %v\n",i, i % 4)
	}
	// for {
	// 	fmt.Printf("startNumber = %v  mod 4 = %v\n",startNumber, startNumber % 4)
	// 	startNumber++

	// 	if startNumber > finalNumber {
	// 		break
	// 	}
	// }

} 