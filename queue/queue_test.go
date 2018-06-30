package main

import (
	"strconv"
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	inputs := createSampleList(5)
	for _, v := range inputs {
		q.Enqueue(v)
	}
	//Check if head points to first input and tail to last input
	if q.head.data != inputs[0] && q.tail.data == inputs[len(inputs)-1] {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	inputs := createSampleList(5)
	for _, v := range inputs {
		q.Enqueue(v)
	}

	q.Dequeue()
	//Check if head now points to second input
	if q.head.data != inputs[1] {
		t.Fail()
	}
}

func createSampleList(size int) []string {
	var inputs []string
	for i := 0; i < size; i++ {
		inputs = append(inputs, strconv.Itoa(i))
	}
	return inputs
}
