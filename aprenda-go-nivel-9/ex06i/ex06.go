package main
import "fmt"
import "runtime"

func main()  {
	fmt.Println("OS:", runtime.GOOS, "ARCH:", runtime.GOARCH)
}