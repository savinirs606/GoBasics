package main

import (
	"fmt"
	"time"
)

func FanInWorker(id int, out chan<- string) {
	for i := 0; i < 3; i++ {
		out <- fmt.Sprintf("Worker %d: message %d", id, i)
		time.Sleep(time.Millisecond * 200)
	}
}
