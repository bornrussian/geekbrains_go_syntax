//
// Задача:
// Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов
// и тестом для этой функции.
//

package statistic

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Summ(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total
}
