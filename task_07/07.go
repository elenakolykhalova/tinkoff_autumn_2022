package main

import (
	"fmt"
	"strings"
	"sort"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main(){
	var planMarket string
	var N int
	var tmp string

	fmt.Scan(&planMarket, &N)
	for i := 0; i < N; i++{
		var d1, d2 int
		fmt.Scan(&d1, &d2)
		if d2 == len(planMarket) {
			tmp = planMarket[d1 - 1:]
		} else {
			tmp = planMarket[d1 - 1:d2]
		}
		tmpSort := SortString(tmp)
		count := 0 //счетчик шагов
		iCount := 0 //позиция курсора
		for i := 0; i < len(tmpSort); i++{
			for j := iCount; j < len(tmp); j++{
				if tmpSort[i] == tmp[j]{
					break
				}
				count++
				iCount++
				if j == len(tmp) - 1 {
					iCount = 0
				}
			}
		}
		fmt.Println(count)
	}
}

// test

// hello
// 3
// 1 5
// 1 2
// 2 5