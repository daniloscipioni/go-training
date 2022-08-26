package main

import "testing"

func BenchmarkAdd(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// Execução
// go test -bench . ". para executar todos os bechs da pasta"
// go test -cover
// go test -coverprofile c.out
// go tool cover -html=c.out ← abre no browser
// go tool cover -h ← para mais detalhes