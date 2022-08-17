package main

import (
	// "bufio"
	"fmt"
	// "log"
	// "os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("*.txt")
	// if err != nil {
	// 	log.Fatalf("Error when opening file: %s", err)
	// }
	// defer file.Close()
	// fileScanner := bufio.NewScanner(file) // можно подставить os.Stdin вместо file

	mapRes := make(map[string][]string)
	iMap := 0 //счетчик уровней
	// for fileScanner.Scan() {
	// 	s := fileScanner.Text()
	var s string
	for fmt.Scanln(&s); s != ""; fmt.Scanln(&s) {
		

		if s == "{" {
			iMap++
		} else if s == "}" {
			for _, v := range mapRes[strconv.Itoa(iMap)] {
				mapRes[v] = mapRes[v][:len(mapRes[v])-1]
			}
			delete(mapRes, strconv.Itoa(iMap))
			iMap--
		} else {
			res := strings.Split(s, "=")
			var tmp int

			flag := true
			if v, err := strconv.Atoi(res[1]); err == nil {
				tmp = v
				flag = false
			} else if _, ok := mapRes[res[1]]; ok && len(mapRes[res[1]]) != 0 {
				tmp, _ = strconv.Atoi(mapRes[res[1]][len(mapRes[res[1]])-1])
			}

			mapRes[strconv.Itoa(iMap)] = append(mapRes[strconv.Itoa(iMap)], res[0])
			if _, ok := mapRes[res[0]]; ok {
				if len(mapRes[res[0]]) >= iMap+1 {
					mapRes[res[0]][iMap] = strconv.Itoa(tmp)
				} else {
					mapRes[res[0]] = append(mapRes[res[0]], strconv.Itoa(tmp))
				}
			} else {
				mapRes[res[0]] = append(mapRes[res[0]], strconv.Itoa(tmp))
			}
			if flag {
				fmt.Println(tmp)
			}
		}
		s = ""
	}
	fmt.Println(mapRes)
}
