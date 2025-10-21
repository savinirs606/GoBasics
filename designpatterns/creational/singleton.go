package main

import "sync"

/*
Implemented using functions and interfaces
*/

var once *sync.Once

func AddOnce() {

	once.Do(func() int {

	})
}
