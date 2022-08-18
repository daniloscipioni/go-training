package main
import (
	"fmt"
)

func main() {

	esporteFavorito := "StockCar"	

	switch esporteFavorito{
		
	case "Futebol", "Futebol Salao":
		fmt.Println("Futebol")
	case "F1","StockCar":
		fmt.Println("Corrida")	
	}
} 