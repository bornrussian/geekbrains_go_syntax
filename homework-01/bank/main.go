package bank

import "fmt"

func Deposit() {
	fmt.Printf("Привет. Я - считатель процентов на банковском вкладе\n")
	var money, percent float64
	fmt.Print("Напиши, сколько денег ты готов положить на счёт: ")
	fmt.Scan(&money)
	fmt.Print("Напиши желаемый готовой процент: ")
	fmt.Scan(&percent)

	for i := 1; i <= 5; i++ {
		money += money / 100 * percent
		fmt.Printf("После %d года на счёте накопится %0.2f денег\n", i, money)
	}
}
