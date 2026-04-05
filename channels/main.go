package main

import (
	"context"
	"fmt"
	"time"
)

func countTillCtxDone(ctx context.Context) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Counting stopped, err: ", ctx.Err())
			return
		default:
			time.Sleep(1 * time.Second)
			i += 1
			fmt.Println(i)
		}
	}
}

func main() {

	ctx := context.Background()

	//cancelCtx, cancel := context.WithCancel(ctx)
	//go countTillCtxDone(cancelCtx)

	//timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(3*time.Second))
	//go countTillCtxDone(timeoutCtx)
	//
	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	go countTillCtxDone(deadlineCtx)

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}
