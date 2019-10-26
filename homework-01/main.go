package main

import (
	"bornrussian/golang0000/bank"
	"bornrussian/golang0000/rub2usd"
	"bornrussian/golang0000/triangle"
)

func main() {
	bank.Deposit()
	rub2usd.Convert()
	triangle.Calc()
}
