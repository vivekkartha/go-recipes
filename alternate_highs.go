package main

//Rearrange to make alternate numbers higher than its neighbors
import "fmt"

func main() {
	a := []int{9, 6, 8, 3, 7}
	reArrange(a)
}
func reArrange(a []int) {
	for i := 1; i < len(a); i += 2 {
		if a[i] < a[i-1] {
			swap(a, i, i-1)
		} else if i+1 < len(a) && a[i] < a[i+1] {
			swap(a, i, i+1)
		}
	}
	fmt.Println(a)
}

func swap(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
