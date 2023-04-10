package main

import (
	"fmt"
)

/* Задание 4. Область видимости переменных
Что нужно сделать
Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные. Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат. Затем вызовите по очереди три функции, передавая результат из одной в другую. */

var first int = 5
var second int = 10
var third int = 2

func addFirstNum(integer int) int {
  return integer + first
}
func addSecondNum(integer int) int {
  return integer + second
}

func addThirdNum(integer int) int {
  return integer + third
}

func main() {
	fmt.Println("Задание 4. Область видимости переменных")

  fmt.Println("Результат сложения числа 10 с первой переменной:", addFirstNum(10))
  fmt.Println("Результат сложения числа 10 со второй переменной:", addSecondNum(10))
  fmt.Println("Результат сложения числа 10 с третьей переменной:", addThirdNum(10))
  fmt.Println("Результат последовательного выполнения всех трех функций с числом 10:", addThirdNum(addSecondNum(addFirstNum(10))))
}
