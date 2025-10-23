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

	inputSliceForBubbleSort := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("Bubble Sort \n %v \n", BubbleSort(inputSliceForBubbleSort))

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

func BubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s)-i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = swap(s[j-1], s[j])
				// s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
	return s
}

func swap(a, b int) (int, int) {
	if a == b {
		return a, b
	}
	a = a + b
	b = a - b
	a = a - b
	return a, b
}
