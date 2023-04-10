package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Задание 1. Подсчёт чисел в массиве
Что нужно сделать
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого. При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.

========================================================

Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.

Что оценивается
Верность индекса.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.
*/

const size = 10

func unsortedSearch(a [size]int, val int) (index int) {
	index = -1
	for i := 0; i < len(a); i++ {
		if a[i] == val {
			index = i
			break
		}
	}
	return
}

func findSorted(a [size]int, value int) (index int) {
	index = -1
	min := 0
	max := size - 1
	for max >= min {
		middle := (max + min) / 2
		if a[middle] == value {
			index = middle
			break
		} else if a[middle] > value {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}
	return
}

func partition(a [size]int, low, high int) ([size]int, int) {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return a, i
}

func quickSort(a [size]int, low, high int) [size]int {
	if low < high {
		var p int
		a, p = partition(a, low, high)
		a = quickSort(a, low, p-1)
		a = quickSort(a, p+1, high)
	}
	return a
}

func sortArray(a [size]int) [size]int {
	return quickSort(a, 0, len(a)-1)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var array [size]int

	for i, _ := range array {
		array[i] = rand.Intn(10 * size)
	}
	fmt.Println(array)
	value := array[rand.Intn(size)]
	index := unsortedSearch(array, value)
	fmt.Printf("Число %v в заданном массиве найдено.\nЕго индекс: %v\n", value, index)
	value = rand.Intn(20 * size)
	index = unsortedSearch(array, value)
	fmt.Printf("Число %v не найдено в заданном массиве.\nИндекс: %v\n", value, index)

	fmt.Println(sortArray(array))
	value = array[rand.Intn(size)]
	index = findSorted(sortArray(array), value)
	fmt.Printf("Число %v в заданном массиве найдено.\nЕго индекс: %v\n", value, index)
	value = rand.Intn(20 * size)
	index = findSorted(array, value)
	fmt.Printf("Число %v не найдено в заданном массиве.\nИндекс: %v\n", value, index)
}
