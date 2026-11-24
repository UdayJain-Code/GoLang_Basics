package main
import "fmt"
const age  = 30
//ERROR Cant use this syntax here name:=golang

func main()  {
	const name string = "Uday Jain"
	fmt.Println(age)

	const(
		port = 5000
		host = "local host"
	)
	
	fmt.Println(port,host)

}