package main

import (
	"fmt"
	"time"
)

// goroutines are lightweight threads in go used to handle concurrency/ multithreading

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {

	// but this is blocking right?
	// for i := 0; i <= 10; i++ {
	// 	task(i)
	// }

	// for i := 0; i <= 10; i++ {
	// 	go task(i) // now this will run the task parallely
	// 	// you'll notice that nothing gets printed, since Go has its own scheduler to run goroutines. So when you run this code, the scheduler schedules the goroutine.
	// 	// hence this becomes non-blocking
	// 	// main also runs on a goroutine
	// 	// in order to see the output, we need to somehow keep the code alive since the scheduler once schedules the program terminates as it has nothing else apart from this for loop.
	// }
	// time.Sleep(time.Second * 2) // now the output you'll see won't be in any specific order since it is running inside a goroutine now.

	for i := 0; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Second * 2) // but is it safe to put sleep like this and block your main method? Of course not.
	// checkout waitgroupingoroutine.go to fix the problem
}
