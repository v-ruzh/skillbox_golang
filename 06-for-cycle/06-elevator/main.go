package main

import (
	"fmt"
)
/* Задание 6 (по желанию). Движение лифта
Что нужно сделать
В доме — 24 этажа. Лифт должен ходить вверх-вниз, пока не доставит всех пассажиров на первый этаж. Три пассажира ждут на четвёртом, седьмом и десятом этажах. При движении вверх лифт не должен останавливаться, при движении вниз — должен собирать всех, но не более двух человек в лифте. При этом лифт каждый раз доезжает до самого верхнего этажа и только после этого начинает движение вниз. Напишите программу, которая доставит всех пассажиров на первый этаж. */

func main() {
	fmt.Println("Движение лифта.")

  currentFloor := 1 
  firstFloor := 1
  topFloor := 24 
  floorStep := 1
  peopleInElevator := 0
  passenger1 := 10
  passenger2 := 7
  passenger3 := 4
  
  
  for  passenger1 != firstFloor || passenger2 != firstFloor || passenger3 != firstFloor || currentFloor != firstFloor {
    fmt.Println(currentFloor , "этаж," , peopleInElevator , "человек в лифте.") 
    currentFloor += floorStep
  
    if currentFloor == topFloor  {
      floorStep = -1
    } else if currentFloor == firstFloor {
      floorStep = 1
      peopleInElevator = 0
      /*fmt.Println(currentFloor , "этаж," , peopleInElevator , "человек в лифте.")*/
    } else if floorStep == -1 {
    if passenger1 == currentFloor && peopleInElevator < 2 {
      fmt.Println("Пассажир с", currentFloor , "этажа вошел в лифт.")
      peopleInElevator++ 
      fmt.Println("В лифте" , peopleInElevator , "человек.")
      passenger1 = 1
      }
    if passenger2 == currentFloor && peopleInElevator < 2 {
      fmt.Println("Пассажир с", currentFloor , "этажа вошел в лифт.")
      peopleInElevator++
      passenger2 = 1
      } 
    if passenger3 == currentFloor && peopleInElevator < 2 {
      fmt.Println("Пассажир с", currentFloor , "этажа вошел в лифт.")
      peopleInElevator++
      fmt.Println("В лифте" , peopleInElevator , "человек.")
      passenger3 = 1
      } 
    }
    if passenger1 == 1 && passenger2 == 1 && passenger3 == 1 && currentFloor == 1 {
      fmt.Println("1 этаж. Все пассажиры на первом этаже.") 
      break}
  }
}