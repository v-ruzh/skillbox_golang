package main

import (
	"fmt"
  "strings"
)

func generate (open, closing int , queue string, result *[]string){
  if open == 0{
    *result = append(*result, queue + strings.Repeat(")", closing))
    return
  }
  if closing > open {
    generate(open, closing - 1, queue + ")", result)
  }
  generate(open - 1, closing, queue + "(", result)
}

func main() {
	fmt.Println("Задание 5 (по желанию). Комбинации круглых скобок")

  result := make([]string, 0)
  fmt.Println("Введите количество пар скобок:")
  var n int
  fmt.Scan(&n)
  generate(n, n, "", &result)
  fmt.Println(result)
  }