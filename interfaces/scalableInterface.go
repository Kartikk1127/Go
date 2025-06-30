package main

import "fmt"

// in other languages we explicitly tell teh compiler that a class is implementing an interface and then implement its member functions
// in Go, if a struct has a method definition which has a declaration in an interface -- compiler implicitly implements the interface -- is it good?

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

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

type fakePaymentGateway struct{}

func (f fakePaymentGateway) pay(amount float32) {
	fmt.Println("Making payment using fake payment gateway", amount)
}

func main() {

	// fakePaymentGateway := fakePaymentGateway{}
	razorpayPG := razorpay{}
	newPayment := payment{
		// gateway: fakePaymentGateway,
		gateway: razorpayPG,
	}
	newPayment.makePayment(100)
}
