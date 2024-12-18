package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, int, bool) {
	return "Golang", 4, true
}

func process(fn func(a int) int) {
	fn(1)
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func summation(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total

}

func main() {
	nums := []int{1, 2, 3, 5}
	sum := 0

	//_ is the index
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println(sum)

	m := map[string]string{"name": "Hassan", "command": "Go"}

	for k, v := range m {
		fmt.Println(v)
		fmt.Println(k)
	}

	//this prints the unicode for each alphabet
	for a, b := range "Hassan" {
		fmt.Println(a, b)
	}

	result := add(3, 5)
	fmt.Println(result)

	//i dont want to use lang3 so i used _
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	fn := func(a int) int {
		return 4
	}
	process(fn)

	funct := processIt()
	funct(6)

	myNumbers := []int{3, 4, 76, 54, 2}
	summed := summation(myNumbers...)
	fmt.Println(summed)

}
