package main

import (
	"fmt"
	"sync"
)

func PrintOdd(n int, oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i += 2 {
		<-oddChan
		fmt.Println("Odd goroutine value: ", i)
		if i+1 <= n {
			evenChan <- true
		}

	}
}

func PrintEven(n int, oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= n; i += 2 {
		<-evenChan
		fmt.Println("Even chan value: ", i)
		if i+1 <= n {
			oddChan <- true
		}
	}
}

//func PrintOdd(n int, oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for i := 1; i <= n; i += 2 {
//		<-oddChan
//		fmt.Println("Odd goroutine val:", i)
//		if i+1 <= n { // only signal if an even number still exists
//			evenChan <- true
//		}
//
//	}
//}
//
//func PrintEven(n int, oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for i := 0; i <= n; i += 2 {
//		<-evenChan
//		fmt.Println("even goroutine val:", i)
//		if i+1 <= n { // only signal if an even number still exists
//			oddChan <- true
//		}
//
//	}
//
//}
