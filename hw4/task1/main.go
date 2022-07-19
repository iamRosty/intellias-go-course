package main

import (
	"fmt"
	"sync"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(n))
	for i := 0; i < len(n); i++ {
		s := n[i]
		go func([]int) {
			defer wg.Done()
			sum(s)
		}(s)
	}
	wg.Wait()
}
func sum(slc []int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	fmt.Println("slice elements sum: ", sum)
}
