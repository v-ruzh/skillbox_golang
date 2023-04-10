package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

/* Цель задания
Написать программу аналог cat.



Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.



Что оценивается
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать два соединённых файла в результирующий.
*/

func main() {
	flag.Parse()
	flags := flag.Args()
	if len(flags) < 3 {
		for _, flg := range flags {
			f, err := os.Open(flg)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			defer f.Close()
			buf := bufio.NewScanner(f)
			buf.Scan()
		}
	} else {
		str := make([]string, 0)
		for i := 0; i < 2; i++ {
			f, err := os.Open(flags[i])
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			defer f.Close()
			buf := bufio.NewScanner(f)
			buf.Scan()
			str = append(str, buf.Text())
		}
		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		content := strings.Join(str, " ")
		file.WriteString(content)
	}
}
