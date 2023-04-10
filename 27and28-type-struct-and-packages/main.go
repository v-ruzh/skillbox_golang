package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"student/student"
)

/*
27.8 Практическая работа
Цель задания
Научиться работать с композитными типами данных: структурами и картами



Что нужно сделать
Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру в хранилище map[studentName] *Student.
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent, далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища. Также необходимо реализовать методы put, get.

Зачёт:
при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее ожидает следующую строку или сигнал EOF (Сtrl + Z);
при получении сигнала EOF программа должна вывести имена всех студентов из map.

==========================================================================

28.5 Практическая работа
Цель задания
Научиться декомпозировать и рефакторить код на примере программы, написанной в прошлом уроке.

Что нужно сделать
В прошлом домашнем задании вы написали программу для работы со студентами. Мы указываем в стандартный ввод данные о студенте, а после получения сигнала об окончании работы программы она выводит имена всех студентов на экран.
Вам нужно отрефакторить код прошлого домашнего задания. Декомпозируйте его так, чтобы логике одной сущности соответствовал один пакет.
Для того, чтобы вы могли использовать методы и переменные, которые объявлены в других пакетах, сделайте их экспортируемыми.
*/

func main() {

	fmt.Println("Данная программа хранит данные об оценках студентов.")
	students := make(map[string]*student.Student, 0)

	fmt.Println("Введите имя, возраст и оценку студента через пробел. Для окончания работы с хранилищем нажмите ctrl+z.")
	input := bufio.NewReader(os.Stdin)

	for {
		inputInfo, err := input.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("Хранилище было успешно обновлено. Список студентов:\n")

			break
		}
		inputFields := strings.Fields(inputInfo)

		if len(inputFields) != 3 {
			fmt.Printf("Количество введенной информации не соответствует формату хранения данных. Нужны имя, возраст и оценка студента. Попробуйте ещё раз.\n")
			continue
		}
		studentName := inputFields[0]
		studentAge, err := strconv.Atoi(inputFields[1])
		if err != nil {
			fmt.Println("Ошибка при вводе возраста студента.")
		}
		studentGrade, err := strconv.Atoi(inputFields[2])
		if err != nil {
			fmt.Println("Ошибка при вводе оценки студента.")
		}

		studentIn := student.Student{
			Name:  studentName,
			Age:   studentAge,
			Grade: studentGrade,
		}
		if _, err := student.Get(students, studentIn.Name); err != nil {
			student.Put(students, &studentIn)
		} else {
			fmt.Printf("Отметка этого студента уже сохранена.\n")
		}
	}
	for _, value := range students {
		fmt.Printf(value.Name + " " + strconv.Itoa(value.Age) + " " + strconv.Itoa(value.Grade) + "\n")
	}
}
