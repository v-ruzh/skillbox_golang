package main

import (
	"fmt"
  "os"
  "log"
  "time"
)

/* Задание 1. Работа с файлами
Что нужно сделать
Напишите программу, которая на вход получала бы строку, введённую пользователем, а в файл писала № строки, дату и сообщение в формате:

2020-02-10 15:00:00 продам гараж.

При вводе слова exit программа завершает работу. */

func main() {
	fmt.Println("Задание 1. Работа с файлами.")
  fmt.Println("Введите данные")
  var uInput string
  fmt.Scan(&uInput)
  t := time.Now().Format("2006-01-02 15:04:05")
  if uInput == "exit" {return}
  
  file, err := os.Create("uInput.txt")

  if err != nil {
    fmt.Println("Не смогли создать файл")
    log.Fatal(err)
    return
  }
  defer file.Close()
  file.WriteString(t + " " + uInput)
  fmt.Println(file.Name())
}
