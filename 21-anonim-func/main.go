package main

import (
	"fmt"
)

/*Задание 1. Расчёт по формуле
Что нужно сделать
Напишите функцию, производящую следующие вычисления.

S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.

Тип S должен быть во float32.

==========================================

Задание 2. Анонимные функции
Что нужно сделать
Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).

Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.
*/

func calc(x int16, y uint8, z float32) (S float32) {
	S = 2*float32(x) + float32(y)*float32(y) - 3/z
	return S
}

func getSomeFunc(A func(int, int) int) {
	i := 2
	j := 4
	fmt.Println("First this:", A(i, j))
	defer func() {
		fmt.Println("Then this:", A(1, 2))
	}()

}

func main() {
	var x int16 = 15
	var y uint8 = 16
	var z float32 = 4.5

	fmt.Printf("\nЗадание 1 - рассчет по формуле.\n\nРезультат рассчёта по формуле 2*x + y^2 − 3/z \nпри x равном %v, y равном %v и z равном %v: %.2f\n", x, y, z, calc(x, y, z))

	a := func(x, y int) int { return x * y }
	b := func(x, y int) int { return x + y }
	c := func(x, y int) int { return x*y - 15 }
	fmt.Printf("\nЗадание 2 - анонимные функции и команда defer.\n")
	getSomeFunc(a)
	getSomeFunc(b)
	getSomeFunc(c)

}
