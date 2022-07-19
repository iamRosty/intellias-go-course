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
	for i, v := range n {
		go func(v []int, i int) {
			defer wg.Done()
			sum(v, i)
		}(v, i)
	}
	wg.Wait()
}
func sum(slc []int, i int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	fmt.Printf("slice %v:  %v\n", i+1, sum)
}
