package main

import (
	"flag"
	"fmt"
	"strings"
)

/* Цель задания
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).
*/

func findSubString(s, sub string) bool {
	var a bool
	//fmt.Println(strings.LastIndex(s, sub))
	if strings.LastIndex(s, sub) != -1 {
		a = true
	} else {
		a = false
	}
	return a
}

func main() {
	var s string
	var sub string
	flag.StringVar(&s, "str", "default", "set value")
	flag.StringVar(&sub, "substr", "default", "set value")
	flag.Parse()
	fmt.Printf("Проверка на наличие в строке %v строки %v.\nРезультат:\n", s, sub)
	fmt.Println(findSubString(s, sub))
}
