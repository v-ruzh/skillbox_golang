package main

import (
	"fmt"
)
/*Что нужно сделать
Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число больше пяти.

Рекомендация
Пример работы программы:

Три числа.
Введите первое число:
3
Введите второе число:
6
Введите третье число:
2
Среди введённых чисел есть число больше 5.
*/
func main() {
	
  fmt.Println("Три числа.")
  
  fmt.Println("Введите первое число:")
  var number1 int
  fmt.Scan(&number1)

  fmt.Println("Введите второе число:")
  var number2 int
  fmt.Scan(&number2)

  fmt.Println("Введите третье число:")
  var number3 int
  fmt.Scan(&number3)

  if  ( number1 > 5 && number2 > 5 ) || (number2 > 5 && number3 > 5) || (number1 > 5 && number3 > 5) || ( number1 > 5 && number2 > 5 && number3 > 5) {
    fmt.Println("Два или более введённых числа больше 5")
    } else if number1 > 5 {
      fmt.Println("Среди введённых чисел есть число больше 5, это число" , number1)
    } else if number2 > 5 {
        fmt.Println("Среди введённых чисел есть число больше 5, это число" , number2)
      } else if number3 > 5 {
          fmt.Println("Среди введённых чисел есть число больше 5, это число" , number3)
      } else {
              fmt.Println("Все введённые числа меньше или равны 5")
      } 
    
}
