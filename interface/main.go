package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// open clsoe principle must be followed
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

func main() {
	stripePaymentGw := stripe{}
	newPayment := payment{
		gateway: stripePaymentGw,
	}
	razorpayPaymentGw := razorpay{}
	newPayment2 := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
	newPayment2.makePayment(200)
}
