package main

import (
	"fmt"
)

func ex1Basic() {
	fmt.Println("EX1: Basic Unbuffered send/receive")
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Println("sent", i)
		}
		close(ch)
	}()

	// receiver until the ch is closed
	for v := range ch {
		fmt.Println("received ", v)
	}

}

/*
package main

import (
	"fmt"
	"sync"
	"time"
)

// 1) basic unbuffered send/receive
func ex1Basic() {
	fmt.Println("EX1: basic unbuffered send/receive")
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Println(" sent", i)
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(" recv", v)
	}
	fmt.Println()
}

// 2) buffered channels (send won't block until buffer full)
func ex2Buffered() {
	fmt.Println("EX2: buffered channel")
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println(" two values buffered")
	// this send would block if buffer full and no receiver:
	go func() {
		ch <- 30
		fmt.Println(" background sent 30")
	}()
	time.Sleep(100 * time.Millisecond) // allow goroutine to run (it will block)
	fmt.Println(" recv", <-ch)
	fmt.Println(" recv", <-ch)
	fmt.Println(" recv", <-ch)
	fmt.Println()
}

// 3) directional channels (send-only and receive-only)
func sendOnly(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
}
func recvOnly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(" recvOnly got", v)
	}
}
func ex3Directional() {
	fmt.Println("EX3: directional channels")
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go recvOnly(ch, &wg)
	sendOnly(ch, 3)
	close(ch)
	wg.Wait()
	fmt.Println()
}

// 4) fan-in and fan-out
func ex4FanInFanOut() {
	fmt.Println("EX4: fan-in and fan-out")
	// fan-in: multiple producers to single channel
	in := make(chan int)
	var wg sync.WaitGroup
	producer := func(id int) {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			in <- id*10 + i
		}
	}
	wg.Add(2)
	go producer(1)
	go producer(2)

	// closer: waits for producers then closes channel
	go func() {
		wg.Wait()
		close(in)
	}()

	// fan-out: one consumer stage that spawns workers to process incoming values
	out := make(chan int)
	var w sync.WaitGroup
	worker := func(id int, in <-chan int, out chan<- int) {
		defer w.Done()
		for v := range in {
			fmt.Printf(" worker %d processing %d\n", id, v)
			out <- v * 2
		}
	}
	w.Add(2)
	go worker(1, in, out)
	go worker(2, in, out)

	// close out when workers done
	go func() {
		w.Wait()
		close(out)
	}()

	for v := range out {
		fmt.Println(" result:", v)
	}
	fmt.Println()
}

// 5) simple worker pool
func ex5WorkerPool() {
	fmt.Println("EX5: worker pool")
	type job struct{ id, payload int }
	jobs := make(chan job)
	results := make(chan int)
	var wg sync.WaitGroup

	// start workers
	worker := func(id int, jobs <-chan job, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for j := range jobs {
			time.Sleep(50 * time.Millisecond) // simulate work
			fmt.Printf(" worker%d handled job%d\n", id, j.id)
			results <- j.payload * 2
		}
	}
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results, &wg)
	}

	// producer
	numJobs := 6
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- job{i, i * 10}
		}
		close(jobs)
	}()

	// closer for results
	go func() {
		wg.Wait()
		close(results)
	}()

	// collect
	for r := range results {
		fmt.Println(" collected result", r)
	}
	fmt.Println()
}

// 6) pipeline with cancel and timeout using select
func ex6PipelineCancelTimeout() {
	fmt.Println("EX6: pipeline with cancel/timeout")
	type stageFunc func(<-chan int) <-chan int

	// generator
	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}
	// stage that multiplies
	mult := func(factor int) stageFunc {
		return func(in <-chan int) <-chan int {
			out := make(chan int)
			go func() {
				for v := range in {
					time.Sleep(30 * time.Millisecond)
					out <- v * factor
				}
				close(out)
			}()
			return out
		}
	}

	// build pipeline: gen -> mult(2) -> mult(3)
	c := gen(1, 2, 3, 4)
	c = mult(2)(c)
	c = mult(3)(c)

	// read with timeout
	timeout := time.After(200 * time.Millisecond)
	for {
		select {
		case v, ok := <-c:
			if !ok {
				fmt.Println(" pipeline finished")
				return
			}
			fmt.Println(" got", v)
		case <-timeout:
			fmt.Println(" timeout; cancel reading")
			return
		}
	}
}

func main() {
	ex1Basic()
	ex2Buffered()
	ex3Directional()
	ex4FanInFanOut()
	ex5WorkerPool()
	ex6PipelineCancelTimeout()
}
*/
