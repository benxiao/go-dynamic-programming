package knapsack


func max(i0 int, i1 int) int {
	if i0 > i1 { return i0}
	return i1
}


/*

 */
func Knapsack(values []int, weights []int, weightLimit int) (int, []int) {
	nItems := len(values)
	cache := make([][]int, nItems+1)
	for i:=0; i!=nItems+1; i++ {
		cache[i] = make([]int, weightLimit+1)
	}

	cacheItem := make([][][]int, nItems+1)
	for i:=0; i!=nItems+1; i++ {
		row := make([][]int, weightLimit+1)
		cacheItem[i] = row
		for j:=0; j!=weightLimit+1; j++ {
			row[j] = make([]int, 0)
		}
	}

	for i:=0; i!=nItems; i++ {
		for j:=1; j!=weightLimit+1; j++ {
			dontTake := cache[i][j]
			dontTakeItem := cacheItem[i][j]
			if j < weights[i] {
				cache[i+1][j] = dontTake
				cacheItem[i+1][j] = cacheItem[i][j]
			} else {
				take := cache[i][j-weights[i]] + values[i]
				takeItem := append(cacheItem[i][j-weights[i]], i)
				if take > dontTake {
					cache[i+1][j] = take
					cacheItem[i+1][j] = takeItem
				} else {
					cache[i+1][j] = dontTake
					cacheItem[i+1][j] = dontTakeItem
				}
			}
		}
	}
	//fmt.Println(cache)
	//fmt.Println(cacheItem)
	return cache[nItems][weightLimit], cacheItem[nItems][weightLimit]
}

