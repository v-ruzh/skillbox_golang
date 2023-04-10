package main

import (
	"fmt"
)
/*
Задание 3. Вывод ёлочки
Что нужно сделать
Усложним задачу рисования: попробуйте вывести ёлочку. В первой строке выведите одну звёздочку, во второй — на две больше, в третьей — ещё на две больше, и так до количества строк, указанных пользователем.

Правила вывода ёлочки: она симметрична, количество строк соответствует введённому пользователем. 

Советы и рекомендации
Необходимо сначала выводить пробелы, а затем — звёздочки. Посмотрите, как количество пробелов и звёздочек в каждой строке связано с введённым количеством строк и номером текущей строки. Внутри цикла по строкам, скорее всего, понадобятся два цикла: один — для вывода пробелов, второй — для вывода звёздочек.

Пример работы программы:

Вывод ёлочки.
Введите высоту ёлочки:
5
    *
   ***
  *****
 *******
*********
*/

func main() {
	fmt.Println("Вывод ёлочки.")
  var h int

  fmt.Println("Введите высоту ёлочки:")
  fmt.Scan(&h)
  
  for i := 1; i <= h; i++ {
    for j := 0 ; j < h - i ; j++ {
      fmt.Print(" ")
      }
      for j := 0 ; j < i * 2 - 1 ; j++ {
        fmt.Print("*")
        }
        fmt.Println()
    }
 
}