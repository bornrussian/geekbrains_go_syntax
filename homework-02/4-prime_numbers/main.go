//
// Задача: * Заполнить массив из 100 элементов различными простыми числами.
//

package main

import (
    "fmt"
    "os"
    "strconv"
)

const isVerboseMode = true
const howMuchPrimeNumbersToFind = 100

func main() {
    var numbersFound []int

    for try:=2;;try++ {
        if isVerboseMode {
            fmt.Printf("Попытка определить, является ли число %d простым...\n",try)
        }

        isPrime:=true
        for prev:=2;prev<try;prev++ {
            if isVerboseMode {
                fmt.Printf("  пытаемся делить %d на %d...\n",try,prev)
            }
            if try%prev==0 {
                if isVerboseMode {
                    fmt.Printf("    %d делится на %d, т.е. не является простым числом\n",try,prev)
                }
                isPrime=false
                break
            }
        }
        if isPrime {
            if isVerboseMode {
                fmt.Printf("  %d не делится ни на одно из предыдущих, т.е. это простое число, запоминаем его в массив\n",try)
            }
            numbersFound=append(numbersFound,try)
        }

        if len(numbersFound)==howMuchPrimeNumbersToFind {
            if isVerboseMode {
                fmt.Printf("Найдено %d простых чисел. Нам хватит.\n",howMuchPrimeNumbersToFind)
            }
            break
        }
    }

    fmt.Printf("Массив из %d элементов различных простых чисел:\n",howMuchPrimeNumbersToFind)
    fmt.Println(numbersFound)

    filename:=strconv.Itoa(howMuchPrimeNumbersToFind)+" простых чисел.txt"
    fmt.Printf("Сохраняем массив в файл '%s' ...\n",filename)
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Не получилось открыть файл на запись :-(\n")
        return
    }
    defer file.Close()
    for _,value:=range numbersFound {
        file.WriteString(strconv.Itoa(value)+"\n")
    }
}
