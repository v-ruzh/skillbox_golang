package main


import (
	"fmt"
  "unicode"
)

func main() {
  fmt.Println("Определение количества слов, начинающихся с большой буквы.")
	s1 := "'Go is an Open source programming Language that makes it \nEasy to build simple, reliable, and efficient Software.'"
  count := 0
  
  for _, symb := range s1 {
      if unicode.IsUpper(symb) {
        count++
      }
  }
  fmt.Printf("В строке %v\nсодержится %v заглавных букв.\n" , s1 , count) 
}