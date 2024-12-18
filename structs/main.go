package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //this has nano second precision
	customer            //embedding
}

// constructor for struct order
func newOrder(id string, amount float32, status string) *order {
	Myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &Myorder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	fmt.Println("hello world")
	newCustomer := customer{
		name:  "John",
		phone: "0300-0987654",
	}
	Myorder := order{
		id:       "1",
		amount:   500,
		status:   "received",
		customer: newCustomer,
	}

	Myorder.createdAt = time.Now()

	Myorder.changeStatus("confirmed")
	fmt.Println(Myorder.getAmount())
	fmt.Println(Myorder)

	Myorder2 := newOrder("3", 580, "received")

	fmt.Println(Myorder2.getAmount())

	//if struct is to be used only once
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
