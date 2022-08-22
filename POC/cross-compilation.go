package main

import "fmt"
import "runtime"
// Comando para compilar para o Windows
// GOOS=windows GOARCH=amd64 go build <filename>.go
func main()  {
	fmt.Println("Esse é programa do exercicio de compilação cruzada. Foi compilado num Linux/amd64 e agora está rodando num sistema:", runtime.GOARCH, runtime.GOOS)
}