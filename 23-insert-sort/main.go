package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.

==============================

Задание 2. Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает (либо сразу сортирует в обратном порядке, как посчитаете нужным).
*/

const size int = 10

func selectSort(a [size]int) [size]int {
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}
	return a
}

func bubbleSort(intSlice []int) []int {
	for i := (len(intSlice) - 1); i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if intSlice[i] > intSlice[j] {
				intSlice[i], intSlice[j] = intSlice[j], intSlice[i]
			}
		}
	}
	return intSlice
}

func userEntry() []int {
	var input string
	fmt.Println("Сортировка пузырьком с переворотом массива заданного пользователем.\nВведите массив, который нужно отсортировать, разделяя значения запятыми без пробелов:")
	fmt.Scan(&input)
	stringSlice := strings.SplitN(input, ",", len(input))
	// fmt.Println(stringSlice)
	intSlice := make([]int, 0)
	for _, s := range stringSlice {
		b, _ := strconv.Atoi(s)
		intSlice = append(intSlice, b)
	}

	return intSlice

}

func main() {
	fmt.Println(bubbleSort(userEntry()))

	var arrTen = [size]int{5, 3, 5, 10, 6, 25, 14, 55, 2, 1}

	fmt.Printf("\n==============\n\nСортировка массива длинной 10.\nНесортированный массив:\n%v\nМассив отсортированный вставкой:\n%v", arrTen, selectSort(arrTen))
}
