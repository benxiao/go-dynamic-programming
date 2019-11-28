package main

import (
	"fmt"
	"github/ranxiao/go-dynamic-programming/knapsack"
)

func main(){
	values := []int{10, 40, 30, 50}
	weights := []int{5, 4, 6, 3}
	fmt.Println(knapsack.Knapsack(values, weights, 10))
}