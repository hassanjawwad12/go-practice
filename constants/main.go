package main
import "fmt" 

const price = 10.56;

//cannot use :price outside of main function
func main () {
	//constant cannot be changed later on
	const name string="John"
	fmt.Println(name)
	fmt.Println(price)

	//this is const grouping
	const (
		port = 8080
		host = "localhost"
	)
	fmt.Println(port,host)
}

