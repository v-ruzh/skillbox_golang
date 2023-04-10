package main

import (
	"fmt"
)

/* Задание 3. Именованные возвращаемые значения
Что нужно сделать
Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его. Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.*/

func multiply(uInt int) (multiplied int) {
  multiplied = uInt * 5
  return
}

func sum(uInt int) (summed int) {
  summed = uInt + 10
  return
}

func main() {
	fmt.Println("Задание 3. Именованные возвращаемые значения")
  var uInt int
  fmt.Scan(&uInt)
  fmt.Println(multiply(uInt))
  fmt.Println(sum(uInt))
}
