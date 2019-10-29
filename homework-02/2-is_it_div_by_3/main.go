//
// Задача: Написать функцию, которая определяет, делится ли число без остатка на 3.
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
    fmt.Println ("Давайте определим, делится ли число, которое Вы введете, без остатка на 3.")

    fmt.Print ("Введите его: ")
    var input string
    fmt.Scan(&input)
    number, err := strconv.ParseInt(input, 10, 64)
    if err!=nil {
        fmt.Printf ("Ошибка: '%s' не похоже на целое 64-ти битное число.\n", input)
        return
    }

    if isItDiv(number,3) {
        fmt.Printf ("Да! '%d' - делится без остатка на 3.\n", number)
    } else {
        fmt.Printf ("Нет! '%d' - не делится без остатка на 3.\n", number)
    }
}
