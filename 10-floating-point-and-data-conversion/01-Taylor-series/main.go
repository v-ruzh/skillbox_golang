package main

import (
	"fmt"
  "math"
)

/* Задание 1. Разложение ex в ряд Тейлора
Что нужно сделать
Напишите программу, вычисляющую ex посредством разложения в ряд Тейлора с заданной пользователем точностью. Пользователь вводит значение x и с точностью до какого знака после запятой необходимо вычислить значение функции.

Советы и рекомендации 
Получить значение точности (эпсилон) можно, разделив 1 на 10, возведённую в степень, равную количеству знаков после запятой.

Для x = 0 и n = 1 вывод должен быть 1

Для x = 1 и n = 3 вывод должен быть 2,7182539682539684

Для x = 1 и n = 5 вывод должен быть 2,7182815255731922 */

func main() {
	fmt.Println("Разложение ex в ряд Тейлора.")
  fmt.Println("Введите значение x:")
  var x float64
  fmt.Scan(&x)
  fmt.Println("С точностью до какого знака после запятой необходимо вычислить значение функции?")
  var n float64
  fmt.Scan(&n)

  var epsilon float64 = 1 /  math.Pow (10, n)
  var result, lastResult float64
  fact := 1 //изначальное значение факториала
  y := 0 //счетчик для факториала
  for {
    
    if y > 0 {
      fact *= y
    }
    
    result += math.Pow(x, float64(y)) / float64(fact)
    
    if math.Abs(result - lastResult) < epsilon {
      fmt.Println(result)
      break
      }
    
    y++
    lastResult = result  
  }
}
