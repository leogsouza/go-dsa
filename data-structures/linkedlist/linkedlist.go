package linkedlist

import (
	"fmt"
)

// Node is a Node of LinkedList
type Node struct {
	data int
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.data)
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedList) Head() int {
	return ll.head.data
}

// Clear removes all elements from linked list, O(n)
func (ll *LinkedList) Clear() {
	trav := ll.head
	for trav != nil {
		next := trav.next
		trav.next = nil
		trav.data = 0
		trav = next
	}
	ll.head, ll.tail, trav = nil, nil, nil
	ll.size = 0
}

// Size returns the total the elements of the linked list
func (ll *LinkedList) Size() int {
	return ll.size
}

// IsEmpty checks if the linked list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

// Append puts the element on tail of the linked list
func (ll *LinkedList) Append(data int) {
	n := &Node{data, nil}
	if ll.IsEmpty() {

		ll.head, ll.tail = n, n
		ll.size++
		return
	}

	ll.tail.next = n
	ll.tail = ll.tail.next
	ll.size++

}

// Prepend puts the element on head of the linked list
func (ll *LinkedList) Prepend(data int) {
	n := &Node{data, nil}
	if ll.IsEmpty() {
		ll.head, ll.tail = n, n
		ll.size++
		return
	}

	head := ll.head
	ll.head = n
	ll.head.next = head
	ll.size++
}
