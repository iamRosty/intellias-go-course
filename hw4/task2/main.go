package main

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	var result int
	chanSum := make(chan int, 3)
	for _, v := range n {
		go func(v []int) {

			sum(v, chanSum)
		}(v)
	}
	for i := 0; i < len(n); i++ {
		result += <-chanSum
	}
	close(chanSum)
	print("result: ", result)
}

func sum(slc []int, chanSum chan int) {
	var sum int
	for _, v := range slc {
		sum += v
	}
	chanSum <- sum
}
