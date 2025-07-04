package main

import (
	"fmt"
	"time"
)

// channels are used to send and receive data.
// imagine there are multiple goroutines running, you want to pass some data from one goroutine to another goroutine. It can be done with the help of the channels.
// communication between 2 coroutines can be done with the help of channels.
// deadlock is a concept of operating system, when multiple processes running concurrently hold a resource. Other processes are waiting for this resource but the first isn't releasing as it is waiting for another resource which in turn is held by the process waiting for the first resource. Deadlock!!
// this can be avoided using some sleep methods?

// sending
func processNum(numChan chan int) {
	// defer wg.Done()

	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}
}

func main() {
	// messageChannel := make(chan string)
	// messageChannel <- "ping" // channels are blocking i.e. this operation is blocking. Till when? Until the receiver is ready to receive it. Is receiver ready yet? No.
	// message := <-messageChannel
	// fmt.Println(message)
	// var wg sync.WaitGroup
	// numChan := make(chan int)
	// // wg.Add(1)
	// go processNum(numChan) // this is an example of sending data from one goroutine(main) to another goroutine(processNum) via channel
	// // wg.Wait()
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// numChan <- 5
	// time.Sleep(time.Second * 2)
	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result // this is also a blocking operation
	// fmt.Println(res)

	// done := make(chan bool)
	// go task(done)

	// <-done // this stays blocked unless it gets the result

	// emailChan := make(chan string, 100) // means we can send 100 size of data without it being a blocking operation, beyond this would be blocking.
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("Done sending")

	// close(emailChan)
	// // fmt.Println(<-emailChan)
	// // fmt.Println(<-emailChan)

	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}

}

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending mail to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing task...")
// 	// done <- true
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }
