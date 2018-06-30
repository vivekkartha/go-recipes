//Queue using linked-list
package queue

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main() {
	fmt.Println("Type 'exit' to quit")
	fmt.Println("1: Push")
	fmt.Println("2: Pop")
	fmt.Println("3: Display All")

	b := bufio.NewReader(os.Stdin)
	queue := NewQueue()

	for {
		input, _ := b.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		if input == "1" {
			fmt.Println("Type string to push:")
			data, _ := b.ReadString('\n')
			data = strings.TrimRight(data, "\n")
			queue.Enqueue(data)
		} else if input == "2" {
			queue.Dequeue()
		} else if input == "3" {
			queue.Display()
		} else if input == "exit" {
			os.Exit(0)
		}
	}
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
