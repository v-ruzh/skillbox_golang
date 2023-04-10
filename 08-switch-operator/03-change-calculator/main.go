package main

import (
	"fmt"
)

func main() {
	fmt.Println("Расчёт сдачи.")
  
  var bills = []int {5 , 5 , 5 , 10 , 20}
  fmt.Println(lemonadeChange(bills))
  bills = []int {10 , 10}
  fmt.Println(lemonadeChange(bills))
  bills = []int {5 , 5 , 10 , 10 , 20}
  fmt.Println(lemonadeChange(bills)) 

  }

  func lemonadeChange(bills []int) bool {
    fives, tens := 0 , 0

    for _, b := range bills {
      switch b {
        case 5:
        fives++
        case 10:
        fives--
        tens++
        case 20:
        if tens > 0 {
          tens--
          fives--
        } else {
          fives -= 3
        }
      }
      if fives < 0 || tens < 0 {
        return false
      }
    }
    return true
}