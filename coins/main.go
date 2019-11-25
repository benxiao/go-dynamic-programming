package main

import "fmt"

func coins(n int, options []int) []int{
	var minCount int
	var minChanges []int
	var currentCount int

	cache := make([][]int, n+1)
	for i:=0; i!= n+1; i++ {
		cache[i] = make([]int, 0)
	}
	for i:= 1; i!= n+1; i++ {
		minCount = n
		for _, o := range options {
			if i >= o {
				currentCount = len(cache[i-o]) + 1
				if minCount > currentCount {
					minCount = currentCount
					minChanges = append(cache[i-o], o)
				}
			}
		}
		cache[i] = minChanges
	}
	//fmt.Println(cache)
	return cache[n]
}


func main(){
	fmt.Println(coins(11, []int{1, 5, 6, 8}))
}