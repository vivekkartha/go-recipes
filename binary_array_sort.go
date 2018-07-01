package main

//Sort array of 0s and 1s in linear time
import (
	"fmt"
)

func main() {
	a := []int{0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0}
	var zeroes int
	b := []int{}
	for _, v := range a {
		if v == 0 {
			zeroes++
		}
	}

	//Append 0s
	for i := 0; i < zeroes; i++ {
		b = append(b, 0)
	}

	//Append the rest with 1s
	for i := 0; i < len(a)-zeroes; i++ {
		b = append(b, 1)
	}

	fmt.Println(b)
}
