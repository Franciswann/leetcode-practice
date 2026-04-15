package main

import (
	"fmt"
)

// =============================================================================
// Linked List Data Structure Implementation
// =============================================================================
//
// Data Structure Definition

// Core Operations
// - Traversal r
// - Insertion c
// - Deletion d
// - Search r
// - Reverse u

// Edge Cases:
// - Empty list (nil head)
// - Single node list
// - Operations at head/tail

// Complexity Analysis

// Create
// TC: O(1) SC: O(1)

// Read (Traversal):
// TC: O(n) SC: O(1)

// Update (known position)
// TC: O(1) SC: O(1)

// Delete (search + delete)
// TC: O(n) SC: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// Create
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 33}
	node3 := &ListNode{Val: 999}

	// Point node1 to node2
	node1.Next = node2
	// Point node2 to node3
	node2.Next = node3

	// Read
	fmt.Printf("node1 Val:\t%v,\tnode1 Memory Address:\t%v,\tnode1 Next:\t%v\n", node1.Val, node1, node1.Next)
	fmt.Printf("node2 Val:\t%v,\tnode2 Memory Address:\t%v,\tnode2 Next:\t%v\n", node2.Val, node2, node2.Next)
	fmt.Println("=============================================================================")

	// Read - Traversal
	fmt.Print("Traversal Linked List: ")
	List := node1
	for List != nil {
		fmt.Printf("%d -> ", List.Val)
		List = List.Next
	}
	fmt.Printf("nil\n")

	// Update
	fmt.Println("=============================================================================")
	curr := node1
	curr.Next.Next.Val++
	fmt.Printf("n3's value 999+1: %v\n", curr.Next.Next.Val)
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("nil\n")
	fmt.Println("=============================================================================")
	// Insertion - Update value between two numbers
	InsertBetween(node1, 33, 1000)

	fmt.Println("=============================================================================")
	// Delete Node: 77
	DeleteNode(node1, 77)

	// Edge Cases:
	// - Empty list (nil head)
	var head *ListNode
	fmt.Printf("Empty List: %v\n", head)

	// - Single node list 能成立
	single := &ListNode{Val: 0}
	fmt.Printf("Single List: %v, Next: %v\n", single.Val, single.Next)

	// - Operations at head/tail
	// Del head
	head_list := node1.Next
	fmt.Printf("new head value: %v\n", head_list.Val)
	// Insert value before head
	fake_head := node1
	newValue := &ListNode{Val: 0}
	newValue.Next = fake_head
	fmt.Printf("Insert before head: %v\n", newValue.Val)

	// Del tail
	del_tail_list := node1
	for del_tail_list != nil {
		if del_tail_list.Next.Next == nil {
			del_tail_list.Next = nil
			fmt.Printf("The last value: %v\n", del_tail_list.Val)
			break
		}
		del_tail_list = del_tail_list.Next
	}

}

// Insert 77 between 33 and 1000
func InsertBetween(node1 *ListNode, prev, next int) *ListNode {

	curr := node1
	newNode := &ListNode{Val: 77}

	for curr != nil {
		if curr.Val == prev {
			originalNext := curr.Next
			curr.Next = newNode
			newNode.Next = originalNext
		}
		curr = curr.Next
	}

	curr = node1
	fmt.Printf("Insert number between %d and %d: ", prev, next)

	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("nil\n")

	return node1
}

func DeleteNode(node1 *ListNode, target int) *ListNode {

	curr := node1

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == target {
			curr.Next = curr.Next.Next
			break
		}
		curr = curr.Next
	}

	fmt.Printf("Delete number between %d and %d: ", curr.Val, curr.Next.Val)

	curr = node1
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("nil\n")

	return node1
}
