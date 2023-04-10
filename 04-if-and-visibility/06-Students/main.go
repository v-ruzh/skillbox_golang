package main

import (
	"fmt"
)

/* Студенты
Что нужно сделать
Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп. Напишите программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента, а программа определяет, в какую группу он попадёт.

Рекомендация
В одну группу могут попадать студенты из разных частей списка. */

func main() {
	fmt.Println("Студенты.")

  fmt.Println("Введите количество студентов:")
  var studentsCount int
  fmt.Scan(&studentsCount)

  fmt.Println("Введите количество групп:")
  var groupsCount int
  fmt.Scan(&groupsCount)

for studentNumber := 1 ; studentNumber <= studentsCount ; studentNumber++ {
    fmt.Println("Введите номер студента в порядке прибывания, начиная с первого:")
  var studentNumber int
    fmt.Scan(&studentNumber)
    groupNumber := (studentsCount - studentNumber) % groupsCount + 1
  fmt.Println("Студент номер" , studentNumber , "распределён в группу" , groupNumber)
    if studentNumber >= studentsCount { break }
 }
   
  fmt.Println("Все студенты были распределены по группам.")
}


