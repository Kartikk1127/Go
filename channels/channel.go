package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channels are used to send and receive data.
// imagine there are multiple goroutines running, you want to pass some data from one goroutine to another goroutine. It can be done with the help of the channels.
// communication between 2 coroutines can be done with the help of channels.
// deadlock is a concept of operating system, when multiple processes running concurrently hold a resource. Other processes are waiting for this resource but the first isn't releasing as it is waiting for another resource which in turn is held by the process waiting for the first resource. Deadlock!!
// this can be avoided using some sleep methods?
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

	numChan := make(chan int)

	// wg.Add(1)
	go processNum(numChan) // this is an example of sending data from one goroutine(main) to another goroutine(processNum) via channel
	// wg.Wait()

	for {
		numChan <- rand.Intn(100)
	}

	// numChan <- 5

	// time.Sleep(time.Second * 2)

}
