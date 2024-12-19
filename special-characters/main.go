package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	intValue := 42
	floatValue := 123.456
	boolValue := true
	stringValue := "Hello, Go!"
	p := &intValue // Pointer example
	person := Person{Name: "John", Age: 30}

	// Print examples in the specified order
	fmt.Printf("%%v: %v\n", person)      // Default format for structs
	fmt.Printf("%%+v: %+v\n", person)    // Struct with field names
	fmt.Printf("%%#v: %#v\n", person)    // Go syntax representation
	fmt.Printf("%%T: %T\n", person)      // Type of the value
	fmt.Printf("%%d: %d\n", intValue)    // Decimal format for integers
	fmt.Printf("%%b: %b\n", intValue)    // Binary format for integers
	fmt.Printf("%%o: %o\n", intValue)    // Octal format for integers
	fmt.Printf("%%x: %x\n", intValue)    // Hexadecimal (lowercase)
	fmt.Printf("%%X: %X\n", intValue)    // Hexadecimal (uppercase)
	fmt.Printf("%%f: %f\n", floatValue)  // Decimal format for floats
	fmt.Printf("%%e: %e\n", floatValue)  // Scientific notation (lowercase e)
	fmt.Printf("%%E: %E\n", floatValue)  // Scientific notation (uppercase E)
	fmt.Printf("%%s: %s\n", stringValue) // String format
	fmt.Printf("%%q: %q\n", stringValue) // Double-quoted string with escapes
	fmt.Printf("%%p: %p\n", p)           // Pointer address
	fmt.Printf("%%t: %t\n", boolValue)   // Boolean format
}
