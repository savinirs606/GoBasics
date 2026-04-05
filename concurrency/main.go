package main

import (
	"fmt"
)

func OddEvenPrint() {
	n := 10
	wg := &sync.WaitGroup{}

	oddChan, evenChan := make(chan bool), make(chan bool)

	wg.Add(2)
	go PrintOdd(n, oddChan, evenChan, wg)

	go PrintEven(n, oddChan, evenChan, wg)
	oddChan <- true
	wg.Wait()
}

func FanIn() {
	// Create a single channel to collect results
	out := make(chan string)

	// Start multiple workers (fan-in)
	for i := 1; i <= 3; i++ {
		go FanInWorker(i, out)
	}

	// Collect results from all workers
	for i := 0; i < 9; i++ { // 3 workers × 3 messages each
		fmt.Println(<-out)
	}
}

func FanOut() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	wg := &sync.WaitGroup{}

	// Start multiple workers (fan-out)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go FanOutWorker(w, jobs, wg)
	}

	// Send jobs into the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // important: close channel so workers exit

	wg.Wait()
}

func main() {
	//OddEvenPrint()
	//FanIn()
	FanOut()
}
