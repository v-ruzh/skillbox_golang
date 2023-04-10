package main

import (
	"fmt"
)

/*Что нужно сделать
Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше или равны пяти.

Рекомендация
Пример работы программы:

Три числа.
Введите первое число:
3
Введите второе число:
5
Введите третье число:
7
Среди введённых чисел 2 больше или равны 5.*/

func main() {
	fmt.Println("Три числа: ещё попытка.")

  fmt.Println("Введите первое число:")
  var number1 int
  fmt.Scan(&number1)

  fmt.Println("Введите второе число:")
  var number2 int
  fmt.Scan(&number2)

  fmt.Println("Введите третье число:")
  var number3 int
  fmt.Scan(&number3)

  moreThan5 := 0

  if number1 >= 5 {
    moreThan5++
  } 

  if number2 >= 5 {
    moreThan5++
  }

  if number3 >= 5 {
    moreThan5++
  }

  if moreThan5 == 0 {
    fmt.Println("Среди введённых чисел нет ни одного больше или равного 5")
  } else {
    fmt.Println("Среди введённых чисел" , moreThan5 , "больше или равно 5")
  }
  
}
