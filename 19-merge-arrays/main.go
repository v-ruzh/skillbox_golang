package main

import (
	"fmt"
)

/* Задание 1. Слияние отсортированных массивов
Что нужно сделать
Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.

Советы и рекомендации
Обратите внимание на размеры массивов.
В качестве среды разработки может помочь GoLand или VS Code.
Что оценивается
Правильность размеров.
Правильный порядок элементов в конечном массиве.

=====================================================

Задание 2. Сортировка пузырьком
Что нужно сделать
Отсортируйте массив длиной шесть пузырьком. */

// функция bubbleSortThemAll принимает на вход срез и сортирует его данные пузырьковым алгоритмом.
func bubbleSortThemAll(someArray []int) []int {
	for i := 0; i < len(someArray); i++ {
		for j := i; j < len(someArray); j++ {
			if someArray[i] > someArray[j] {
				someArray[i], someArray[j] = someArray[j], someArray[i]
			}
		}
	}
	return someArray
}

// функция mergeArrays принимает два сортированных среза и объединяет их, одновременно также сортируя результат.
func mergeArrays(a []int, b []int) []int {

	var mergedArrays []int

	j, k := 0, 0

	for i := 0; i < (len(a) + len(b)); i++ {
		if j >= len(a) {
			mergedArrays = append(mergedArrays, b[k])
			k++
			continue
		} else if k >= len(b) {
			mergedArrays = append(mergedArrays, a[j])
			j++
			continue
		} else if a[j] > b[k] {
			mergedArrays = append(mergedArrays, b[k])
			k++
		} else {
			mergedArrays = append(mergedArrays, a[j])
			j++
		}
	}
	return mergedArrays
}

// в основной функции заданы значения срезов и формат вывода данных по итогу работы программы.
func main() {
	fmt.Println("Задание 1 и 2. \nСортировка пузырьком и слияние массивов.")

	numArrayL4 := []int{5, 10, 3, 6}
	numArrayL5 := []int{15, 26, 42, 3, 1}
	numArrayL6 := []int{155, 3, 15, 24, 6, 80}
	fmt.Println("Результат сортировки массива длиной шесть пузырьком:", bubbleSortThemAll(numArrayL6))

	bubbleSortThemAll(numArrayL4)
	bubbleSortThemAll(numArrayL5)

	fmt.Println("Результат слияния двух отсортированных массивов длинной 4 и 5", mergeArrays(numArrayL4, numArrayL5))
}
