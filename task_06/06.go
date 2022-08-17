package main

import "fmt"

func main() {
  var N int
  fmt.Scan(&N)
  
  var m, h [][3]int
  var a, b int

  for i := 0; i < N; i++ {
    fmt.Scan(&a, &b)
    m = append(m, [3]int{a, b, i})
  }
  
  var max int
  generateVariants(m, h, &max)
  fmt.Println(max)
}

func generateVariants(a [][3]int, b [][3]int, max *int) {
  for _, s := range a {
    if checkNumbers(b, s) {
      continue
    }
    if len(b) == 0 || b[len(b)-1][1] == s[0] {
      b = append(b, s)
      if *max < len(b) {
        *max = len(b)
      }
      generateVariants(a, b, max)
      b = b[:len(b)-1]
    }
  }
}

func checkNumbers(b [][3]int, f [3]int) bool {
  for _, s := range b {
    if s == f {
      return true
    }
  }
  return false
}

// test
// 7
// 2 6
// 5 6
// 2 5
// 2 2
// 6 8
// 2 2
// 0 2
