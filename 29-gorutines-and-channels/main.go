package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Задание 1. Конвейер
Цели задания
Научиться работать с каналами и горутинами.
Понять, как должно происходить общение между потоками.
Что нужно сделать
Реализуйте паттерн-конвейер:

Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.
Советы и рекомендации
Воспользуйтесь небуферизированными каналами и waitgroup.
*/

func square(inputChan chan int) chan int {
	a := <-inputChan
	outSq := a * a
	fmt.Printf("Квадрат числа %v равен %v\n", a, outSq)
	squareChan := make(chan int)
	go func() {
		squareChan <- outSq
	}()
	return squareChan
}

func doubleSquare(squareChan chan int) chan int {
	a := <-squareChan
	doubleSqChan := make(chan int)
	go func() {
		doubleSqChan <- a * 2
	}()
	fmt.Printf("Число %v умноженное на 2 равно %v\n", a, <-doubleSqChan)
	return doubleSqChan
}

func main() {

	fmt.Printf("Введите любое число через перенос строки.\nДля окончания работы введите слово 'стоп' или 'stop'.\n")
	inputChan := make(chan int)
	for {
		var inputInfo string
		fmt.Scan(&inputInfo)
		if inputInfo == "стоп" || inputInfo == "stop" {
			fmt.Printf("Работа с программой завершена. Отключаюсь.\n")
			close(inputChan)
			os.Exit(1)
		}

		in, err := strconv.Atoi(inputInfo)
		if err != nil {
			fmt.Println("Для работы программы нужно ввести число, а также слова 'стоп' или 'stop' для её завершения.")
			continue
		}
		go func() {
			inputChan <- in
		}()
		sc := square(inputChan)
		go doubleSquare(sc)
	}
}
