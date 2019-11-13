//
// Задача:
//
// Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа
// Фибоначчи. Как теперь заставить спиннер вращаться в течение некоторого времени? 10
// секунд?
//

package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	//const n = 42
	//fibN := fibonacci(n)
	//fmt.Printf("\rFibonacci(%d) = %d", n, fibN)
	time.Sleep(10 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

//func fibonacci(x int) int {
//	if x < 2 {
//		return x
//	}
//	return fibonacci(x-1) + fibonacci(x-2)
//}