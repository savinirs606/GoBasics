package main

import (
	"fmt"
	"sync"
)

/*
 write a simple working code on concurrency

- fan in fan out
- worker pools
*/

var fanInWg *sync.WaitGroup

func fanInSender(id int, ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- id*10 + i
	}
}

func fanInReceiver(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}
	wg.Done()
}

func fanIn() {
	fanInWg = &sync.WaitGroup{}
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go fanInSender(i, ch)
	}
	fanInWg.Add(1)
	go fanInReceiver(ch, fanInWg)
	fanInWg.Wait()
}

var fanOutWg *sync.WaitGroup

// one sender multiple receivers

func fanOutSender(ch chan int) {
	for i := 0; i <= 15; i++ {
		ch <- i
	}
}

func fanOutReceivers(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	wg.Done()
}

func fanOut() {
	fanOutWg = &sync.WaitGroup{}
	ch := make(chan int)
	go fanOutSender(ch)

	for i := 0; i < 3; i++ {
		fanOutWg.Add(1)
		go fanOutReceivers(ch, fanOutWg)
	}
	fanOutWg.Wait()
}
