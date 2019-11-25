package main

import "fmt"

func max(x0 int, x1 int) int {
	if x0 > x1 { return x0 }
	return x1
}


func lcs(s0 string, s1 string) int {
	l0, l1 := len(s0), len(s1)
	cache := make([][]int, l0+1)
	for i:=0; i!= l0+1; i++ {
		cache[i] = make([]int, l1+1)
	}
	for i, ch0 := range(s0) {
		for j, ch1 := range(s1) {
			if ch0 == ch1 {
				cache[i+1][j+1] = cache[i][j] + 1
			} else {
				cache[i+1][j+1] = max(
					cache[i][j+1],
					cache[i+1][j])
			}
		}
	}
	return cache[l0][l1]
}



func main()  {
	fmt.Println(lcs("ABCDE", "AC"))

}