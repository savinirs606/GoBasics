package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func NewNode(val int, next *Node) *Node {
	return &Node{
		Value: val,
		Next:  next,
	}
}

func ReverseList(head *Node) {
	if head == nil || head.Next == nil {
		return

	}
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head = prev

	PrintList(head)
	return

}

func PrintList(head *Node) {

	for curr := head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}

func ListOperations() {
	fifth := NewNode(0, nil)
	fourth := NewNode(1, fifth)
	third := NewNode(2, fourth)
	second := NewNode(3, third)
	first := NewNode(4, second)
	head := NewNode(5, first)
	ReverseList(head)

}
