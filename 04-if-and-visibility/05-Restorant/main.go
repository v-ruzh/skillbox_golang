package main

import (
	"fmt"
)

/* Ресторан
Что нужно сделать
Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате. В ресторане действуют следующие правила:

по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый;
по пятницам, если сумма чека превышает 10 000 рублей, включается дополнительная скидка в размере 5%;
если число гостей в одной компании превышает пять человек, автоматически включается надбавка на обслуживание 10%.
Рекомендация
Пример работы программы:

Введите день недели:
5
Введите число гостей:
7
Введите сумму чека:
12000
Скидка по пятницам: 600
Надбавка на обслуживание: 1200
Сумма к оплате: 12600 */


func main() {
	fmt.Println("Ресторан.")

  fmt.Println("Введите день недели:")
  var day int
  fmt.Scan(&day)

  fmt.Println("Введите число гостей:")
  var guestsCount int
  fmt.Scan(&guestsCount)

  fmt.Println("Введите сумму чека:")
  var bill int
  fmt.Scan(&bill)

  serviceFee := bill / 100 * 10
  mondayDiscount := bill / 100 * 10
  fridayDiscount := bill / 100 * 5
  
  if day == 1 {
    fmt.Println("Скидка по понедельникам:" , mondayDiscount)
    fmt.Println("Сумма к оплате:" , bill - mondayDiscount)
  } else if day == 5 && bill > 10000 {
      fmt.Println("Скидка по пятницам за заказ от 10000 рублей:" , fridayDiscount)
      fmt.Println("Сумма к оплате:" , bill - fridayDiscount)
  } else if guestsCount > 5 && day != 1 && ( day != 5 && bill > 10000 ){
        fmt.Println("Надбавка за обслуживание:" , serviceFee)
        fmt.Println("Сумма к оплате:" , bill + serviceFee)
  } else if guestsCount > 5 && day == 1 {
          fmt.Println("Скидка по понедельникам:" , mondayDiscount)
          fmt.Println("Надбавка за обслуживание:" , serviceFee)
          fmt.Println("Сумма к оплате:" , bill + serviceFee - mondayDiscount)
  } else if guestsCount > 5 && ( day == 5 && bill < 10000 ) {
            fmt.Println("Надбавка за обслуживание:" , serviceFee)
            fmt.Println("Сумма к оплате:" , bill + serviceFee )
  } else if guestsCount > 5 && ( day == 5 && bill > 10000 ) {
              fmt.Println("Скидка по пятницам за заказ от 10000 рублей:" , fridayDiscount)
              fmt.Println("Надбавка за обслуживание:" , serviceFee)
              fmt.Println("Сумма к оплате:" , bill + serviceFee - fridayDiscount)
  } else if guestsCount > 5 && ( day != 5 && day != 1) {
                fmt.Println("Надбавка за обслуживание:" , serviceFee)
                fmt.Println("Сумма к оплате:" , bill + serviceFee )
  } else {
                  fmt.Println("Сумма к оплате" , bill)
  }
}
