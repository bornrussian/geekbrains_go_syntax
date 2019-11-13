//
// Представленная ниже программа состоит из трех горутин, соединенных двумя каналами.
// Первая горутина генерирует бесконечный поток целых чисел,
// вторая возводит эти числа в квадрат, а третья выводит их в консоль.
//
// Задача:
//
// Перепишите программу-конвейер, ограничив количество передаваемых для обработки
// значений и обеспечив корректное завершение всех горутин.
//

package main

import "fmt"

const VALUE_FIRST int = 1
const VALUE_LAST int = 10

func main() {
	naturals := make(chan int, 1)
	squares := make (chan int, 1)

	// генерация
	go func(out chan<- int) {
		for x := VALUE_FIRST; x <= VALUE_LAST; x++ {
			out <- x
		}
		close(out)
	}(naturals)

	// возведение в квадрат
	go func(out chan<- int, in <-chan int) {
		for x:= range in {
			out <- x*x
		}
		close (out)
	}(squares, naturals)

	// печать
	for value := range squares {
		fmt.Println(value)
	}
}
