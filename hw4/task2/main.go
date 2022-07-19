package main

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	result := 0
	chanSum := make(chan int)
	for _, v := range n {
		go func(v []int) {
			sum(v, chanSum)
		}(v)
		result += <-chanSum
	}
	print("result: ", result)
}

func sum(slc []int, chanSum chan int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	chanSum <- sum
}
