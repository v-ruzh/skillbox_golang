package main

import (
	"fmt"
)

/* Задание 1. Функция, принимающая аргументы
Что нужно сделать
Напишите функцию, которая принимает в качестве аргументов два числа типа int, вычисляет сумму чётных чисел заданного диапазона и выводит результат в консоль.

Рекомендация
Если первое число, переданное в качестве аргумента, будет больше второго, просто поменяйте их местами.
*/

func evenSum(a, b int) {
	var numbers []int
	sum := 0

	if a > b {
		c := a
		a = b
		b = c
	}
	for i := a; i <= b; i++ {
		numbers = append(numbers, i)
	}
	for _, number := range numbers {
		if number%2 == 0 {
			sum += number
		}
	}
	fmt.Println("Сумма всех чётных чисел заданного диапазона равна", sum)
}

func main() {
	fmt.Println("Задание 1. Функция, принимающая аргументы.")
	var j int
	var y int
	fmt.Println("Введите первое число:")
	fmt.Scan(&j)
	fmt.Println("Введите второе число:")
	fmt.Scan(&y)

	evenSum(j, y)
}
