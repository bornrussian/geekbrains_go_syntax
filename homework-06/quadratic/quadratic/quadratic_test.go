package quadratic

import (
	"testing"
)

func TestQuadraticFloat(t *testing.T) {
	_, _, err := QuadraticFloat(2, 2, 2)
	if err == nil {
		t.Error("QuadraticFloat(2,2,2): Should be error: D<0, but got:", err)
	}

	x1, x2, _ := QuadraticFloat(1, 2, 1)
	x1correct := -1.0
	x2correct := -1.0
	if x1 != x1correct || x2 != x2correct {
		t.Error("QuadraticFloat(1,2,1): Should be", x1correct, x2correct, "but got:", x1, x2)
	}
}
func TestQuadraticComplex(t *testing.T) {
	x1, x2, err := QuadraticComplex(2, 2, 2)
	x1correct := complex(-0.5, +2)
	x2correct := complex(-0.5, -2)
	_ = err
	if x1 != x1correct || x2 != x2correct {
		t.Error("QuadraticComplex(2,2,2): Should be", x1correct, x2correct, "but got:", x1, x2)
	}
}
