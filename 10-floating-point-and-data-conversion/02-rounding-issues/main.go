package main

import (
	"fmt"
  "math"
)

/* Задание 2. Проблемы округления процентов
Что нужно сделать
В связи с особенностями хранения дробных чисел они не очень подходят для хранения денежных значений (может потеряться значимая часть при переполнении, да и дробная часть больше двух знаков не нужна). Но мы попробуем решить задачу начисления процентов, используя именно их. 

Пусть пользователь вводит сумму, которую он кладёт в банк, ежемесячный процент капитализации и количество лет, в течение которых будет открыт вклад.
Необходимо ежемесячно пересчитывать сумму, округляя её до целого количества копеек в меньшую сторону. Например, если после начисления процентов остаток включает 35,78 копейки, то оставляем только 35 копеек, а дробную часть отбрасываем. 
По окончании работы программы необходимо вывести, какую сумму получит вкладчик на руки и какая сумма будет зачислена в пользу банка за счёт округления копеек.
Советы и рекомендации
Для округления копеек можно умножить получившееся после капитализации процентов число на 100, округлить в меньшую сторону, затем поделить опять на 100. Отбрасываемую часть можно получить вычитанием текущего значения остатка на счёте и того, который был до округления.

Для 1000 рублей, 1% и одного года программа должна вывести 1126,78 и 0,04350000000022192 (возможно меньше знаков).

Для 1000 рублей, 1% и десяти лет программа должна вывести 3299,41 и 0,5216000000013992 (возможно меньше знаков). */

func main() {
	fmt.Println("Проблемы округления процентов.")
  fmt.Println("Введите сумму:")
  var dep  float64
  fmt.Scan(&dep)
  fmt.Println("Введите ежемесячный процент капитализации (в %):")
  var perc float64
  fmt.Scan(&perc)
  fmt.Println("Введите длительность хранения вклада:")
	var dur int
  fmt.Scan(&dur)

	roundingTotal := dep
	totalNotRounded := dep
	roundAmount := 0.0

	for i := 0; i < dur * 12; i++ {
		roundingTotal += roundingTotal * perc / 100
		roundingTotal = math.Floor(roundingTotal * 100) / 100
		totalNotRounded += totalNotRounded * perc / 100
	}

	roundAmount = totalNotRounded - roundingTotal
  benefit := roundingTotal - dep


	fmt.Printf("Сумма в %v рублей принесёт вам %f дохода за %v лет. За счёт округления банк в свою очередь получит %f рублей. \n" , dep , benefit , dur , roundAmount)

}
