package main

import (
	"fmt"
)

/* Задание 4. Предыдущее занятие на if
Что нужно сделать
Есть три переменные со значениями 0. Первую нужно наполнить до 10, вторую — до 100, третью — до 1000. Напишите цикл, в котором эти переменные будут увеличиваться на один.

Рекомендация
Используйте условия для пропуска переменных, которые уже достигли своих лимитов. */

func main() {
	fmt.Println("Предыдущее занятие на if.")
  a := 0
  b := 0
  c := 0
  
  for {
    a++
    b++
    c++
    fmt.Println("a" , a)
    fmt.Println("b" , b)
    fmt.Println("c" , c)
    if a == 10 { 
      a = 0 }
    if b == 100  { 
      b =  0 }
    if c == 1000 { 
      c = 1000
      break }
  }
}
