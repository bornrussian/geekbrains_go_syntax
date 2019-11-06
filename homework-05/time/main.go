//
// Задача:
//
// Изучите статью ​ Time in Go: A primer ​( https://machiel.me/post/time-in-go-a-primer/ ).
// В письменном  виде кратко изложите свое мнение: что из этой статьи стоило бы добавить в методичку?
// Если считаете, что ничего, — так и напишите, приведя аргументы.
//

// Мухачев ВВ:
// Необходимо добавить в примеры операцию сравнения времени:

package main

import (
	"fmt"
	"time"
)

func main() {
	// Сравниваем дату goReleaseDate и cReleaseDate:
	cReleaseDate := time.Date(1972, time.May, 15, 12, 0, 0, 0, time.UTC)
	goReleaseDate := time.Date(2009, time.November, 10, 12, 0, 0, 0, time.UTC)
	if goReleaseDate.After(cReleaseDate) {
		fmt.Println("Go was released years after C!")
	} else {
		fmt.Println("C was inspired by Go. Go was first.")
	}
}