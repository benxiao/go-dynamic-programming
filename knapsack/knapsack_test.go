package knapsack

import (
	"fmt"
	"testing"
)

func intArrayEqual(a0 []int, a1 []int) bool {
	if len(a0) != len(a1){
		return false
	}
	for i:=0; i!=len(a0); i++ {
		if a0[i] != a1[i] {
			return false
		}
	}
	return true
}

func TestKnapsack(t *testing.T) {
	values := []int{10, 40, 30, 50}
	weights := []int{5, 4, 6, 3}
	maxVal, items := Knapsack(values, weights, 10)
	groudTruth := []int{1, 3}
	fmt.Println(maxVal)
	if maxVal != 90 || (!intArrayEqual(groudTruth, items)){
		fmt.Println(groudTruth)
		fmt.Println(items)
		t.Error("simple test failed")
	}
}