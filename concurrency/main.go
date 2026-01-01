package main

import (
	"fmt"
)

func printOdd(oddChan, evenChan chan int, limit int, done chan bool) {
	for num := range oddChan {
		fmt.Println(num)
		if num >= limit-1 {
			close(evenChan) // stop even goroutine
			done <- true
			return
		}
		evenChan <- num + 1
	}
}
func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	done := make(chan bool)
	limit := 102
	// Odd printer
	go printOdd(oddChan, evenChan, limit, done)

	// Even printer
	go func(limit int) {
		for num := range evenChan {
			fmt.Println(num)
			if num >= limit {
				close(oddChan) // stop odd goroutine
				done <- true
				return
			}
			oddChan <- num + 1
		}
	}(limit)

	// Start sequence
	oddChan <- 1

	// Wait for completion (only one signal needed)
	<-done
}

//package main
//
//func main() {
//	// Channels:
//	/*
//		// Unbuffered Channel
//		MakeUnbufferedChannel()
//		// buffered channel
//		MakeBufferedChannel()
//
//	*/
//	fanOut()
//
//}
