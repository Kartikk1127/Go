package main

import (
	"fmt"
	"time"
)

// use pointers to modify the struct else not needed
// if you don't set a particular field, it will fallback to its default value
// there is this convention to start the names in UPPERCASE
type customer struct {
	name  string
	phone string
}

func newCustomer(name string, phone string) customer {
	customer := customer{
		name:  name,
		phone: phone,
	}
	return customer
}

// order struct - this is how you create a struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nano seconds precision
	customer            // -- this is an example of struct embedding
}

// this is the convention to form a constructor in Go, you have to return a pointer of the struct
func newOrder(id string, amount float32, status string, customer customer) *order {
	// initial setup goes here.
	order := order{
		id:       id,
		amount:   amount,
		status:   status,
		customer: customer,
	}
	return &order
}

// receiver type - now this function is attached with the struct
// unless you give the pointer, the struct will behave as a deep copy and won't modify the original object
func (o *order) changeStatus(status string) {
	o.status = status
}

// while returning something, there's no need to use pointer(although using it won't give you any error)
func (o order) getAmount() float32 {
	return o.amount
}

// structs are custom data structures
func main() {

	customer := newCustomer("kartikey", "9794159484")
	order := newOrder("1", 30.5, "confirmed", customer)

	order.createdAt = time.Now()

	fmt.Println(order.getAmount())
	// order.changeStatus("paid")
	fmt.Println("order struct", *order)

	// below is an example of an inline struct -- does not require a type
	language := struct {
		name   string
		isGood bool
	}{
		name:   "Kartik",
		isGood: true,
	}

	fmt.Println(language)
	fmt.Println(customer)
	fmt.Println(order.customer.name)
}
