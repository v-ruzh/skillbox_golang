package main

import (
	"fmt"
  "os"
  "log"
  "time"
  "io/ioutil"
  "bytes"
)

func main() {
	fmt.Println("Задание 4. Пакет ioutil")
  fmt.Println("Введите данные")
  var uInput string
  fmt.Scan(&uInput)
  t := time.Now().Format("2006-01-02 15:04:05")
  if uInput == "exit" {return}

  var b bytes.Buffer
  fileName := "uInputIoutil.txt"
  text := t + " " + uInput
  textBytes := []byte (text) 
  if err := ioutil.WriteFile(fileName, textBytes, 0777); err != nil {
    log.Fatal(err)
    return
  }
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Println("Не смогли открыть файл")
    log.Fatal(err)
    return
  }
  defer file.Close()

  b.WriteString(text)
  
  resultBytes, err := ioutil.ReadAll(file)
  if err != nil {
    log.Fatal(err)
    return
    }
  fmt.Println("Результат чтения из файла:")
  fmt.Println(string(resultBytes))
}
