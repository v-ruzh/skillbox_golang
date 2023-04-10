package main

import (
	"fmt"
)
/* Задание 3. Расчёт суммы скидки
Что нужно сделать
Напишите программу, которая принимает на вход цену товара и скидку. Посчитайте и верните на экран сумму скидки. Скидка должна быть не больше 30% от цены товара и не больше 2000 рублей.
*/

func main() {
	fmt.Println("Расчёт суммы скидки.")
  bill := 0

  for {
    fmt.Println("Введите стоимость товара:")
    var price int
    fmt.Scan(&price)
  
    fmt.Println("Введите размер скидки:")
    var discountP int
    fmt.Scan(&discountP)

    discount := price / 100 * discountP
    bill = price - discount

    if discountP > 30 || discount > 2000 { continue }

    fmt.Println("Ваша скидка" , discount , "рублей.")
    fmt.Println("К оплате" , bill , "рублей.")
  
  }
}
