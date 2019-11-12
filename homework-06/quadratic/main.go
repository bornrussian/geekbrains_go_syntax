//
// Задача:
//
// * Напишите функцию для вычисления корней квадратного уравнения (алгоритм можно найти в
// Википедии​ ) и тесты к ней.
//

package main

import (
	"fmt"
	"geekbrains_go_syntax/homework-06/quadratic/quadratic"
)

func main() {
	x1,x2,err:=quadratic.QuadraticFloat(4,2,0)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println("x1=",x1,"x2=",x2)
	}
}
