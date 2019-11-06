//
//  Задача: Напишите упрощенный аналог утилиты ​ grep . ​
//

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	flagInvert := flag.Bool("v", false, "invert match")
	flagIgnoreCase := flag.Bool("i", false, "ignore case")

	flag.Parse()

	if flag.NArg()!=2 {
		fmt.Println("Упрощенный аналог утилиты grep")
		fmt.Printf("Использование:\n\t grep [-v] [-i] 'regexp pattern' file\n")
		fmt.Println("Опции:\n")
		fmt.Println("\t -v : исключение строчек, которые соответствую паттерну\n")
		fmt.Println("\t -i : не быть чувствительным к регистру букв\n")
		fmt.Println("Пример:\n\t ./grep -i 'print[f|ln]' main.go\n")
		return
	} else {
		pattern:=flag.Arg(0)
		filename:=flag.Arg(1)

		file,err := os.Open(filename)
		if err!=nil {
			fmt.Println("Error!",err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var fileLine string
			if *flagIgnoreCase {
				fileLine = strings.ToLower(scanner.Text())
			} else {
				fileLine = scanner.Text()
			}

			matched, _ := regexp.MatchString(pattern, fileLine)
			if *flagInvert {
				matched=!matched
			}
			if matched {
				fmt.Println(scanner.Text())
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error!",err)
		}
	}
}