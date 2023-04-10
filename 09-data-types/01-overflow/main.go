package main

import (
	"fmt"
  "math"
)

/* Задание 1. Переполнение
Что нужно сделать
В данном модуле мы рассмотрели примеры по целочисленным типам, их размерам в памяти и то, что происходит при её переполнении. Напишите программу, которая в цикле с использованием встроенных констант (на предельные значения целых чисел, в пакете math) будет подсчитывать, сколько приходится переполнений чисел типа uint8, uint16 в диапазоне от 0 до uint32. 

Советы и рекомендации
Для нахождения количества переполнений используйте цикл. Воспользуйтесь константами максимальных значений из пакета math. */

func main() {
	fmt.Println("Переполнение.")
var (
    overflowUint8 = 0
    overflowUint16 = 0
    maxUint8 = math.MaxUint8
    maxUint16 = math.MaxUint16
  )
  for i := 0; i <= math.MaxUint32; i++ {
    if i == maxUint8 + 1 {
      maxUint8 += math.MaxUint8
      overflowUint8++
      //fmt.Println("overflowUint8:" , overflowUint8) //для проверки
    }
    if i == maxUint16 + 1 {
      maxUint16 += math.MaxUint16
      overflowUint16++
      //fmt.Println("overflowUint16", overflowUint16) //для проверки
    }
  }
  fmt.Println("Переполнений uint8:", overflowUint8)
  fmt.Println("Переполнений uint16:", overflowUint16)
}
