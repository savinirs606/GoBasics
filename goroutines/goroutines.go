package main

import (
	"context"
	"fmt"
	"time"
)

/*

Why Golang , What is so special about golang?

What is Go routine in Go ? How it is different from Thread?

how goroutines are run. go scheduler

*/

func doSomething(ctx context.Context, val chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		case v, ok := <-val:
			fmt.Println(v, ok)
		default:
			fmt.Println("Doing something...")
		}
		time.Sleep(500 * time.Millisecond)
	}

}

//func main() {
//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//defer cancel()
//val := make(chan int)
//go doSomething(ctx, val)
//val <- 5
//
//time.Sleep(10 * time.Second)
//}

func printOdds(limit int, done chan bool) {
	for i := 1; i <= limit; i += 2 {
		fmt.Println("Odd:", i)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func printEvens(limit int, done chan bool) {
	for i := 2; i <= limit; i += 2 {
		fmt.Println("Even:", i)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func main() {
	limit := 20
	done := make(chan bool)

	go printOdds(limit, done)
	go printEvens(limit, done)

	// Wait for both goroutines to finish
	<-done
	<-done

	fmt.Println("Finished printing odd and even numbers.")
}
