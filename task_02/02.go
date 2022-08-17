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
  
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
	if sc.Text() == "" {
		break
	}
	a := sc.Text()
	a = strings.Trim(a, "\n")
	b := strings.Split(a, " ")
	sort.Strings(b)
	for _, v := range b{
	ind_str += v + " "
	}

	if _, ok := res[ind_str]; !ok{
		res[ind_str] = 1
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
