package main
import "fmt" 

func main () {

	//you cannot use an unused variable in go and also single coma 
	var name string= "Hassan"

	//golang infer the type itself
	var shortName="Hassan"

	//much shorter syntax (used when read and assigning value at same time to variable)
	short:="Hassan"

	var age int=10

	//value agr baad main assign karni ho to we have to explicitly mention the type
	var later string;
	later="Hassan"

	price := 10.5

	fmt.Println(name)
	fmt.Println(shortName)
	fmt.Println(age)
	fmt.Println(short)
	fmt.Println(later)
	fmt.Println(price)
}
