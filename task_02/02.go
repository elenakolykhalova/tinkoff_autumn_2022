package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main(){
	var N int
	res := make(map[string]int)
	var ind_str string

	fmt.Scan(&N)

	for i := 0; i < N; i++{
		a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		a = strings.Trim(a, "\n")
		b := strings.Split(a, " ")
		sort.Strings(b)
		for _, v := range b{
			ind_str += v + " "
		}

		if _, ok := res[ind_str]; !ok{
			res[ind_str] = 0
		} else {
			res[ind_str]++
		}
		ind_str = ""
	}
	var value []int
	for _, v := range res{
		value = append(value, v)
	}
	sort.Ints(value)
	fmt.Println(value[len(value)-1])
}

//test
// 5
// MIKHAIL VLADISLAV GRIGORY
// VLADISLAV MIKHAIL GRIGORY
// IVAN ILYA VLADIMIR
// ANDREY VLADIMIR ILYA
// VLADIMIR IVAN ANDREY