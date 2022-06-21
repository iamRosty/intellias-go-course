package main

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	uniqChecker := make(map[int]int)
	for _, val := range arr {
		if _, ok := uniqChecker[val]; !ok {
			uniqChecker[val] = 1
			result = append(result, val)
		}
	}
}
