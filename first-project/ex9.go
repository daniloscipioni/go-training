package main
import (
	"fmt"
)

func main() {
	anoQueNasci := 1989
	anoQueQueroContar:= 2022
	for {
		fmt.Printf("%d\n", anoQueNasci)
		anoQueNasci++

		if anoQueNasci > anoQueQueroContar {
			break
		}
	}

} 