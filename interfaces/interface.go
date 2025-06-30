// package main

// import "fmt"

// type payment struct {
// 	gateway stripe
// }

// // but this is the violation of open closed principle -- open for extension and closed for modification
// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGateway := razorpay{}
// 	// razorpayPaymentGateway.pay(amount)

// 	// stripePaymentGateway := stripe{}
// 	// stripePaymentGateway.pay(amount)

// 	p.gateway.pay(amount)
// }

// type razorpay struct {
// }

// func (r razorpay) pay(amount float32) {
// 	// logic to call the razorpay apis
// 	fmt.Println("Making payment using razorpay", amount)
// }

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

// func main() {
// 	stripePaymentGateway := stripe{}
// 	newPayment := payment{
// 		gateway: stripePaymentGateway,
// 	}
// 	newPayment.makePayment(100)
// }
