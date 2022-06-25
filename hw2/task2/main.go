package main

import (
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	slcStr := strings.Split(input, " ")
	max, _ := strconv.Atoi(slcStr[0])

	min := max
	if len(slcStr) == 1 {
		result = slcStr[0]
	} else {

		for i := 1; i < len(slcStr); i++ {
			val, _ := strconv.Atoi(slcStr[i])
			if max < val {
				max = val
			}
			if min > val {
				min = val
			}
		}
		result = strconv.Itoa(max) + " " + strconv.Itoa(min)
	}
	print(result)
}
