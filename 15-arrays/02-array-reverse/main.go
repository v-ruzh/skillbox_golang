package main

import (
	"fmt"
)

/* Задание 2. Функция, реверсирующая массив
Что нужно сделать
Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
Напишите программу, демонстрирующую работу этого метода.
Что оценивается
При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции, реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.
*/

// Функция reverse инвертирует порядок значений, присвоенных элементам подаваемого на неё массива.
func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		numbers[i], numbers[len(numbers)-1-i] = numbers[len(numbers)-1-i], numbers[i]
	}
	return numbers
}

func main() {
	fmt.Println("Функция, реверсирующая массив.")
	fmt.Println("Введите желаемую длину массива:")
	var l int
	fmt.Scan(&l)
	numbers := make([]int, l)
	for i, _ := range numbers {
		fmt.Printf("Введите %v элемент массива: ", i)
		var n int
		fmt.Scan(&n)
		numbers[i] = n
	}
	fmt.Println(reverse(numbers))
}
