package main

import (
	"fmt"
	"sync"
)

// mutex is mutual exclusion, used to counter the race conditions
// multiple processes modifying the same resource, this modification is not atomic
// good practice is to make the mutex inside the struct for which it is used instead of making it globally
type post struct {
	Views int
	mu    sync.Mutex
}

func (p *post) inc(w *sync.WaitGroup) {

	//unlock is called inside the defer function because complex logics can throw error after locking, in this case
	//the mutex lock will never be released
	defer func() {
		w.Done()      //runs after the function is complete
		p.mu.Unlock() //go routines wait until one goroutine is doing the modification
	}()

	//this can become a bottle-neck
	p.mu.Lock() //saves the resource from being over-written (prevents the race condition)
	p.Views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{Views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait() //until all the waitgroups are complete
	fmt.Println(myPost.Views)
}
