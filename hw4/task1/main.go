package main

import "fmt"

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	sum(n)
}

func sum(slc [][]int) {
	firstAddGoroutine(slc[0])
	secondAddGoroutine(slc[1])
	thirdAddGoroutine(slc[2])
}

func firstAddGoroutine(slc []int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	fmt.Println("slice 1: ", sum)
}

func secondAddGoroutine(slc []int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	fmt.Println("slice 2: ", sum)
}

func thirdAddGoroutine(slc []int) {
	sum := 0
	for _, v := range slc {
		sum += v
	}
	fmt.Println("slice 3: ", sum)
}
