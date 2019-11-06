//
// Задача:
//
// Что бы вы изменили в программе чтения из файла? Приведите исправленный вариант,
// обоснуйте свое решение в комментарии.
//

// Мухачев ВВ:
// Пример хороший.
// Но нужно добавить вывод возможных ошибок, напримерс помощью fmt.Println(err) :

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println("Ошибка открытия файла.", err)
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения размера файла.",err)
		return
	}

	// reading file
	bs := make([]byte,stat.Size())
	_,err = file.Read(bs)
	if err != nil {
		fmt.Println("Ошибка чтения из файла.", err)
		return
	}

	fmt.Println(string(bs))
}