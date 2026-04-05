package main

import (
	"fmt"
	"sync"
	"time"
)

// worker consumes jobs from the jobs channel
func FanOutWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Millisecond * 300) // simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}
