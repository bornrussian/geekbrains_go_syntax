package triangle

import (
	"fmt"
	"math"
)

// triangleSquare считает площадь прямоугольного треугольника по его двум катетам
func triangleSquare(legA, legB float64) float64 {
	return legA * legB / 2
}

// trianglePerimeter считает периметр прямоугольного треугольника по его двум катетам
func trianglePerimeter(legA, legB float64) float64 {
	var legC float64 = math.Sqrt(legA*legA + legB*legB)
	return legA + legB + legC
}

// triangleHypotenuse находит гипотенузу прямоугольного треугольника по его двум катетам
func triangleHypotenuse(legA, legB float64) float64 {
	return math.Sqrt(legA*legA + legB*legB)
}

func Calc() {
	fmt.Printf("Привет. Давай посчитаем всякие площади, периметры и гипотенузы прямоугольного треугольника :-)\n")
	var legA, legB float64

	fmt.Print("Напиши длину первого катета: ")
	fmt.Scan(&legA)
	fmt.Print("Напиши длину второго катета: ")
	fmt.Scan(&legB)

	fmt.Printf("Площадь такого прямоугольного треугольника = %0.2f\n", triangleSquare(legA, legB))
	fmt.Printf("Периметр такого прямоугольного треугольника = %0.2f\n", trianglePerimeter(legA, legB))
	fmt.Printf("Гипотенуза такого прямоугольного треугольника = %0.2f\n", triangleHypotenuse(legA, legB))
}
