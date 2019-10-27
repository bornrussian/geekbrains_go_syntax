package rub2usd

import "fmt"

const rub2usdRate float64 = 63.63

func Convert() {
	fmt.Printf("Привет. Я - конвертер рублей в доллары по курсу %0.2f\n", rub2usdRate)
	var rub float64
	fmt.Print("Напиши, сколько у тебя рублей на обмен?: ")
	fmt.Scan(&rub)
	fmt.Printf("%0.2f рублей можно обменять на %0.2f долларов.\n", rub, rub/rub2usdRate)
}
