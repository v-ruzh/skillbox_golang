package main

import (
	"fmt"
  "strings"
  "strconv"
)
/*
Задание 2. Вывод чисел, содержащихся в строке
Что нужно сделать
Напишите программу, которая выведет все части строки

a10 10 20b 20 30c30 30 dd, 

которые можно привести к числу в десятичном формате.

Рекомендация
Пример работы программы:

Исходная строка:
a10 10 20b 20 30c30 30 dd
В строке содержатся числа в десятичном формате:
10 20 30 */
func main() {
	fmt.Println("Вывод чисел, содержащихся в строке.")
  s := "a10 10 20b 20 30c30 30 dd"
  space := " "
  for strings.Contains(s, space) {
    spaceIndx := strings.Index(s, space)
    word := s[:spaceIndx]
    i , err := strconv.Atoi(word)
    if err == nil {
      fmt.Println(i)
    }
    s = s[spaceIndx+1:]
    s = strings.Trim(s , space)
  }
  i, err := strconv.Atoi(s)
  if err == nil {
    fmt.Println(i)
  }
}
