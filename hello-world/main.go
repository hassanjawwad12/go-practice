package main
import "fmt"  //this is coming from go library

//this is important the main key word 
//we don't use semi-colons in go 
func main () {
	fmt.Println("Hello, World!")
	fmt.Println(1+4)
	fmt.Println(true)
}

//either you run go build hello-world/main.go and then /.main after the executable is created
//or you can run go run hello-world/main.go