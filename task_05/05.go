package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	var N, Q int
	fmt.Scan(&N, &Q)

	var abc []string

	for i := 0; i < N ; i++ {
		var tmp string
		fmt.Scan(&tmp)
		abc = append(abc, tmp)
	}
	for i := 0; i < Q; i++{
		var (
			tmp_word string
			k int
			pref string
			abc_temp []string
		)
		fmt.Scan(&k, &pref)
		for _, v :=range abc {
			if strings.HasPrefix(v, pref){
				abc_temp = append(abc_temp, v)
			}
		}
		sort.Strings(abc_temp)
		if k <= len(abc_temp){
			tmp_word = (abc_temp[k - 1])
			for k, v := range abc {
				if v == tmp_word {
					fmt.Println(k + 1)
				}
			}
		} else {
			fmt.Println("-1")
		}
	}
}

// test
// 5 3
// ad
// a
// abc
// aboba
// b
// 3 a
// 2 ab
// 1 b