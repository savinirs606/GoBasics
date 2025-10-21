package main

import (
	"fmt"
	"sync"
)

var unbufferedWG *sync.WaitGroup

func MakeUnbufferedChannel() {
	unbufferedWG := &sync.WaitGroup{}
	unbufferedChannel := make(chan int)
	unbufferedWG.Add(1)
	go func() {
		defer unbufferedWG.Done()
		v := <-unbufferedChannel
		fmt.Printf("unbufferedChannel value: %d\n", v)
	}()
	unbufferedChannel <- 8
	unbufferedWG.Wait()
}

func MakeBufferedChannel() {
	bufferedChannel := make(chan int, 3)

	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3

LOOP:
	for {
		select {
		case v, _ := <-bufferedChannel:
			fmt.Printf("Buffered Channel value: %d \n", v)
		default:
			break LOOP
		}

	}
}
