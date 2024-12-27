package main

import "fmt"

// enums are enumerated types
type OrderStatus int

const (
	Received OrderStatus = iota + 1 //iota gets auto-incremented
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to", status)

}
func main() {
	changeOrderStatus(Received)

}
