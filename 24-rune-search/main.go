package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.

================================================

Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

chars := [5]rune{'H','E','L','П','М'}

Пример вывода результата в первом элементе массива:

'H' position 0
'E' position 1
'L' position 9
*/

func evenOdd(a []int) (even, odd []int) {
	for _, v := range a {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return even, odd
}

func parseTest(sentences [4]string, chars [5]rune) (index [4][2]string) {
	for i := 0; i < len(sentences); i++ {
		ss := strings.Split(sentences[i], " ")
		lastWord := strings.ToUpper(ss[len(ss)-1])
		if strings.IndexRune(lastWord, chars[i]) > -1 {
			index[i][0] = "Буква: " + strconv.QuoteRune(chars[i]) + " Индекс: "
			index[i][1] = strconv.Itoa(strings.IndexRune(lastWord, chars[i]))
		}
	}

	return
}

func main() {
	a := []int{5, 10, 15, 3, 4, 8, 6}
	even, odd := evenOdd(a)
	fmt.Printf("Изначальный массив:\n%v\nЧётные:\n%v\nНечётные:\n%v\n", a, even, odd)

	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'O', 'L', 'И', 'O', 'S'}

	fmt.Printf("\n===================\n\n")
	fmt.Printf("Массив фраз:\n%v\n\nМассив рун:\n%c\n\nРасположение заданных букв в последних словах в каждой фразе:\n", sentences, chars)
	fmt.Println(parseTest(sentences, chars))

}
