package main

import (
	"fmt"
  "os"
  "bufio"
  "log"
)

/* Задание 3. Уровни доступа
Что нужно сделать
Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные. 

Рекомендация
Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные. */

func main() {
	fmt.Println("Задание 3. Уровни доступа")

  file, err := os.Create("read_only.txt")
  if err = os.Chmod("read_only.txt", 0444); err != nil {
    fmt.Println("Не удалось создать файл")
    log.Fatal(err)
    return
  } 
  defer file.Close()
  writer := bufio.NewWriter(file)
  if err != nil {
    log.Fatal(err)
    return
  }
  writer.WriteString("Ola")
  writer.WriteString("\n")
  writer.WriteRune('ф') // было ы с полным доступом
  writer.WriteString("\n")
  writer.Flush()
} 
