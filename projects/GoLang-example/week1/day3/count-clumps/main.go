package main

import "fmt"

func countClumps(xs []int) int {
	count := 0
	for i := 1; i < len(xs); i++ {
		curr, prev := xs[i], xs[i-1]
		if curr == prev {
			count++
			for ; i < len(xs); i++ {
				curr, prev = xs[i], xs[i-1]
				if curr != prev {
					break
				}
			}
		}
	}
	return count
}
func main() {
	clumps := countClumps([]int{1, 1, 2, 3, 1, 1})
	fmt.Println(clumps)
}
