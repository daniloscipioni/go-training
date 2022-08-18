package main
import (
	"fmt"
)

func main() {
	anoQueNasci := 1989
	anoQueQueroContar:= 2022
	for anoQueNasci <= anoQueQueroContar {
		fmt.Printf("%d\n", anoQueNasci)
		anoQueNasci++
	}

} 