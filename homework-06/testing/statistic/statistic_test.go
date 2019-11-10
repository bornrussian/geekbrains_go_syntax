//
// Задача:
// Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов
// и тестом для этой функции.
//

package statistic

import "testing"

type testpair struct {
	values []float64
	averageResult float64
	summResult float64
}

var tests = []testpair {
	{[]float64{1,2},1.5, 3},
	{[]float64{1,2,3,4,5}, 3, 15},
	{[]float64{-1,1},0, 0},
}

func TestAverageSet(t *testing.T) {
	for _,pair := range tests {
		avg := Average(pair.values)
		if avg != pair.averageResult {
			t.Error(
				"Average() for", pair.values,
				"expected", pair.averageResult,
				"got", avg,
				)
		}
		summ := Summ(pair.values)
		if summ != pair.summResult {
			t.Error(
				"Summ() for", pair.values,
				"expected", pair.summResult,
				"got", avg,
			)
		}
	}
}