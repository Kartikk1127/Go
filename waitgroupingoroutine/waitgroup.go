package main

import (
	"fmt"
	"sync"
)

// wait groups are used to ensure that all running goroutines were finished or not? If yes, can exit the program

// make sure to receive the pointer
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer basically runs after this method's execution
	fmt.Println("Doing some task", id)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // wait until everything finishes
}
