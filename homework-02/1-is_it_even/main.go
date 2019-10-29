//
// Задача: Написать функцию, которая определяет, четное ли число.
//

package main

import (
    "fmt"
    "strconv"
)

func isItDiv(number, base int64) bool {
    if number%base==0 {
        return true
    } else {
        return false
    }
}

func main() {
    fmt.Println ("Давайте определим, чётное ли Вы введёте число.")
    fmt.Println ("Чётное число — целое число, которое делится на 2 без остатка: …, −4, −2, 0, 2, 4, 6, 8, …")

    fmt.Print ("Введите его: ")
    var input string
    fmt.Scan(&input)
    number, err := strconv.ParseInt(input, 10, 64)
    if err!=nil {
        fmt.Printf ("Ошибка: '%s' не похоже на целое 64-ти битное число.\n", input)
        return
    }

    if isItDiv(number,2) {
        fmt.Printf ("Да! '%d' - это чётное число.\n", number)
    } else {
        fmt.Printf ("Нет! '%d' - это не чётное число.\n", number)
    }
}
