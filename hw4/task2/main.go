package main

// Конкурентно порахувати суму усіх слайсів int, та роздрукувати результат.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// “result: 26”
func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	for i, v := range n {
		go func(v []int, i int) {
			sum(v, i)
		}(v, i)
	}
	// Ваша реалізація
}

// func sum(slc []int, i int) {
// 	sum := 0
// 	for _, v := range slc {
// 		sum += v
// 	}
// 	fmt.Printf("slice %v:  %v\n", i+1, sum)
// }
