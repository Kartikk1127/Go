package main

import "fmt"

//  enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota // -- iota is an untyped integer
	Confirmed
	Prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}
func main() {
	changeOrderStatus(Received)
}
