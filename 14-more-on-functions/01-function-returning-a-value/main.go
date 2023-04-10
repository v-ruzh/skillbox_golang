package main

import (
	"fmt"
  "math/rand"
  "time"
)

/* Задание 1. Функция, возвращающая результат
Что нужно сделать
Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечётное.

Рекомендация
Программа запрашивает у пользователя или генерирует случайное число, передает в функцию в качестве аргумента и выводит в консоль результат её работы. */

func evenodd(num int) (even bool) {
  if num%2 == 0 {
  even = true 
  return
  } else {
  even = false
  return 
  }
}

func main() {
	fmt.Println("Задание 1. Функция, возвращающая результат.")
  var uInput int
  fmt.Scan(&uInput)
  var random int
  rand.Seed(time.Now().UnixNano())
  random = rand.Intn(15)

  fmt.Println("user input,", uInput, "=", evenodd(uInput), "; random number,", random, "=", evenodd(random))
}