package main

import "fmt"

func main() {
	a := []int{1,2,3,4,1}
	findPairs(a,5)
}

func findPairs(a []int,sum int){
	fmt.Println("Target sum =",sum)
	var pairs [][] int
	 m:= map[int] int{}
	for i,v := range a{
		complement := sum-v
		if hasKey(m,complement){
			index :=m[complement]
			pairs = append(pairs,[]int{v,a[index]})
		}else{
			m[v]=i
		}
	}

	printSlice(pairs)
}

func printSlice(pairs [][]int){
	for _,v := range pairs{
		fmt.Println(v)
	}
}

func hasKey(m map[int]int,key int) bool {
	_,ok := m[key]
	return ok
}
