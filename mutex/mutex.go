package main

import (
	"fmt"
	"sync"
)

// race conditions occur when multiple processes try to modify the same resource, then that modification would not be atomic.

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock() // but since we might not know when a function might execute completely hence it is a good pratice to put it in deferred block
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views) // with this code snippet race condition will occur. To prevent this, mutex can be used:: lock the resource when it is being modified
}
