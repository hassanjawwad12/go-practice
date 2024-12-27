package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}
func main() {

	for i := 0; i < 10; i++ {
		go task(i) //adding the word go will make it run concurrently
	}

	time.Sleep(time.Second * 2)
}
