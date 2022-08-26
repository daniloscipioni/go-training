package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println("Erro vai CuRiNtHiAnS", err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erro := fmt.Errorf("Não pode valor negativo %v", f)
		return 0, sqrtError{"50.225 N","55.695 W", erro}
	}
	return 42, nil
}
