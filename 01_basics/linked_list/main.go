package main

import (
	"fmt"
)

// TC:   SC:

func main() {

	type ListNode struct {
		Val  int
		Next *ListNode
	}

	// Create
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 33}
	// Point node1 to node2
	node1.Next = node2

	// Read
	fmt.Printf("List1 Val: %v \nList1 Next: %v\nList1 Memory Address: %v\n", node1.Val, node1.Next, node1)
	fmt.Printf("List2 Val: %v \nList2 Next: %v\nList2 Memory Address: %v\n", node2.Val, node2.Next, node2)
	fmt.Println("------------------------------")

	// Traversal Linked List
	fmt.Print("Traversal Linked List: ")
	for node1 != nil {
		fmt.Printf("%d -> ", node1.Val)
		node1 = node1.Next
	}
	fmt.Printf("nil\n")

	// Update
	fmt.Println("------------------------------")
	// target := 5

	// Delete
	fmt.Println("------------------------------")
}
