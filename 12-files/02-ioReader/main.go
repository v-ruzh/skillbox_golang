package main

import (
	"fmt"
  "os"
  "io"
  "log"
  //"time"
)

/* Задание 2. Интерфейс io.Reader
Что нужно сделать
Напишите программу, которая читает и выводит в консоль строки из файла, созданного в предыдущей практике, без использования ioutil. Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.

Рекомендация
Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку. */


func main() {
	fmt.Println("Задание 2. Интерфейс io.Reader")

  /*
  Предыдущее задание для создания файла
  
  fmt.Println("Введите данные")
  var uInput string
  fmt.Scan(&uInput)
  t := time.Now().Format("2006-01-02 15:04:05")

  file, err := os.Create("uInput.txt")

  if err != nil {
    fmt.Println("Не смогли создать файл")
    log.Fatal(err)
    return
  }
  defer file.Close()
  file.WriteString(t + " " + uInput)
  fmt.Println(file.Name()) */

  file, err := os.Open("uInput.txt")
  if err != nil {
    fmt.Println("Не смогли открыть файл")
    log.Fatal(err)
    return
  }
  defer file.Close()

  /* 
  Выясняем размер файла
  
  fileSize, err := file.Stat()
  if err != nil {
    fmt.Println("Не смогли прочитать размер файла")
    log.Fatal(err)
    return
} 
 fmt.Println("Размер файла:",fileSize.Size()) */

  buf := make([] byte, 22)
  if _, err := io.ReadFull(file, buf) ; err != nil {
    fmt.Println("Не смогли прочитать последовательность байтов из файла")
    log.Fatal(err)
    return
  }
  fmt.Printf("%s\n" , buf)
}
