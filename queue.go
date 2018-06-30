//Queue using linked-list
package main

import (
	"fmt"
)

//Node with data and pointer to the next node
type Node struct {
	data string
	next *Node
}

//Queue has head and tail nodes and the size
type Queue struct {
	head *Node
	tail *Node
}

//Enqueue inserts into tail of the Queue
func (queue *Queue) Enqueue(data string) {
	newNode := NewNode(data)
	if queue.head == nil {
		queue.tail = newNode
		queue.head = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
}

//Dequeue removes from head of the Queue
func (queue *Queue) Dequeue() {
	if queue.head == nil {
		fmt.Println("Queue empty")
	} else {
		queue.head = queue.head.next
	}
}

//Display prints the entire queue
func (queue *Queue) Display() {
	current := queue.head
	fmt.Print("\n")
	if current == nil {
		fmt.Println("Empty queue")
	} else {
		for current != nil {
			fmt.Print(current.data)
			current = current.next
		}
	}
	fmt.Print("\n")
}

//NewNode creates and returns a new node pointer
func NewNode(data string) *Node {
	node := &Node{}
	node.data = data
	return node
}

//NewQueue creates and returns a Queue pointer
func NewQueue() *Queue {
	return &Queue{}
}
