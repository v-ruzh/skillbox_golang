package main

import (
	"fmt"
)

/*Задание 2. Дни недели
Что нужно сделать
Пользователь вводит будний день недели в сокращённой форме (пн, вт, ср, чт, пт) и получает развёрнутый список всех последующих рабочих дней, включая пятницу.

Рекомендация
Пример работы программы:

Дни недели.
Введите будний день недели: пн, вт, ср, чт, пт:
вт
вторник
среда
четверг
пятница */


func main() {
	fmt.Println("Дни недели.")
  fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
  var day string
  fmt.Scan(&day)

  switch day {
    case "пн":
    fmt.Println("понедельник")
    fallthrough
    case "вт":
    fmt.Println("вторник")
    fallthrough
    case "ср":
    fmt.Println("среда")
    fallthrough
    case "чт":
    fmt.Println("четверг")
    fallthrough
    case "пт":
    fmt.Println("пятница")
    default:
    fmt.Println("Не знаю такого дня недели")
  }
}
