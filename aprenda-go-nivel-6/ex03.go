package main

import "fmt"


func execFirst() string  {
	fmt.Println("Executa primeiro")
	return ""
}
func execAfter() string  {
	fmt.Println("Executa depois")
	return ""
}




func main()  {
defer execFirst()
	  execAfter()
}