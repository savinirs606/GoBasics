package main

import (
	"fmt"
	"sort"
)

/*
Arrays
slices
arrays vs slices

make command
*/

var sampleIntegerArray = [3]int{}

// Keeps a pointer
// a slice using make command

// playing with slices
// [start to x]
// [x to end]
// reverse -1

func slices() {
	sampleIntegerSlice := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("Sample Slice:", sampleIntegerSlice)

	makeIntegerSlice := make([]int, 5)
	fmt.Println("Sample Slice via make:", makeIntegerSlice)
	makeIntegerSlice = append(makeIntegerSlice, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("Sample Slice via make:", makeIntegerSlice)

	// slice of a slice
	subSlice := makeIntegerSlice[5:]
	fmt.Println("Sample Slice via subSlice:", subSlice)

}

func removeElementFromSlice(slice []int, i int) []int {
	return append(slice[:i-1], slice[i+1:]...)
}

func sortSlice(slice []int) []int {
	sort.Ints(slice)
	return nil
}

func findMax(s []int) int {
	maxVal := s[0]
	for _, v := range s {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
