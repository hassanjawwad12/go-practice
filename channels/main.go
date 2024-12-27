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

func emailSender(emailChan chan string, complete chan bool) {
	defer func() { complete <- true }() //this is the go routine syncronizer
	for email := range emailChan {
		fmt.Println("sending email to:", email)
		time.Sleep(time.Second)
	}
}

func main() {

	//we can send limited amount of data without blocking with buffer channel
	emailChan := make(chan string, 100)
	complete := make(chan bool)

	go emailSender(emailChan, complete)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("done sending")

	//this is important to counter deadlock
	close(emailChan)
	<-complete

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
