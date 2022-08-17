package main

import (
	"fmt"
	"math"
)
func sumDays(data []int) int {
	var sum float64
	for i := 0; i < len(data); i++ {
		sum += math.Pow(-1, float64(i)) * float64(data[i])
	}
	return int(sum)
}

func minMaxSwap(data []int) []int {
	max := data[1]
	ind_max := 1
	min := data[0]
	ind_min := 0
	for i, v := range data{
		if i % 2 == 0 {
			if v < min{
				min = v
				ind_min = i
			}
		} else {
			if v > max {
				max = v
				ind_max = i
			}
		}
	}
	data[ind_min], data[ind_max] = data[ind_max], data[ind_min]
	return(data)
}

func main() {
	var N, tmp int
	fmt.Scan(&N)
	var data []int
	for i := 0; i < N; i++ {
		fmt.Scan(&tmp)
		data = append(data, tmp)
	}
	minMaxSwap(data)
	fmt.Println(sumDays(data))
}