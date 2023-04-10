package main

import (
	"fmt"
)

/* Проверить, что одно из чисел — положительное
Что нужно сделать
Проверка пользовательского ввода на различные ограничения является частой задачей. Попросите пользователя ввести три числа и проверьте, что хотя бы одно является положительным. Результат проверки необходимо сообщить пользователю.

Рекомендация
Используйте логическое сложение. */

func main() {
	fmt.Println("Является ли хотя бы одно из чисел — положительным.")

  fmt.Println("Введите первое число:")
  var a int
  fmt.Scan(&a)

  fmt.Println("Введите второе число:")
  var b int
  fmt.Scan(&b)

  fmt.Println("Введите третье число:")
  var c int
  fmt.Scan(&c)

  if a > 0 || b > 0 || c > 0 {
    fmt.Println("Как минимум одно из введенных чисел больше нуля.")
  } else {
    fmt.Println("Все числа отрицательны.")
  }
  
}
