package main

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	fmt.Println(string(r))
	return string(r)
}

func ReverseWordsInAString(s string) string {
	words := strings.Split(s, " ")
	fmt.Printf("Words %v \n", words)
	result := ""
	for _, word := range words {
		result = result + " " + ReverseString(word)
	}
	return result
}
