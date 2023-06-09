package main

import (
	"fmt"
)

/*
	Задание 2. Функция, принимающая значение по ссылке

Что нужно сделать
Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.

Рекомендация
В методе main создайте и присвойте значения двум переменным типа int, выведите значения этих переменных в консоль до вызова функции и после.
*/
func pointerValue(a, b *int) {
	c := *a
	*a = *b
	*b = c
}
func main() {
	fmt.Println("Задание 2. Функция, принимающая значение по ссылке")
	a := 10
	b := 5
	fmt.Printf("Орининальные значения: a = %v, b = %v \n", a, b)
	pointerValue(&a, &b)
	fmt.Printf("Обновленные значения: a = %v, b = %v \n", a, b)
}
