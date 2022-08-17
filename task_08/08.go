package main

import (
	"fmt"
	"strings"
)

func main(){
	var (
		n, m int
		domen []string
		count int
	)
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++{
		var tmp string
		fmt.Scan(&tmp)
		domen = append(domen, tmp)
	}
	for i:=0; i < m; i++ {
		var prefix string
		var postfix string
		fmt.Scan(&prefix, &postfix)
		for _, v := range domen {
			if strings.HasPrefix(v, prefix) && strings.HasSuffix(v,postfix){
				count++
			}
		}
		fmt.Println(count)
		count = 0
	}
}

// test
// 2 3
// ATSR
// ASR
// S R
// AT R
// A R
