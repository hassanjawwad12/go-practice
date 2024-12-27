package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sending and receiving of unbuffer channel have blocking
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// go routine syncronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing the task...")

}

func main() {

	done := make(chan bool)
	go task(done)
	<-done //block

	//receiving
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result
	fmt.Println(res)

	//sending
	numChan := make(chan int)
	go processNum(numChan)
	for {
		numChan <- rand.Intn(100)
	}
}

//deadlock error code
// messageChan := make(chan string)
// messageChan <- "ping"
// msg := <-messageChan
// fmt.Println(msg)
