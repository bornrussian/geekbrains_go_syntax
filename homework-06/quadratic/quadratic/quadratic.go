package quadratic

import (
	"errors"
	"math"
	"math/cmplx"
)

func QuadraticFloat(A, B, C float64) (x1, x2 float64, err error) {

	if (A == 0) {
		err = errors.New("A==0!")
		return
	}

	D := B*B - 4*A*C
	if (D == 0) {
		x1 = -B / 2 / A
		x2 = x1
		err = nil
		return
	}
	if (D > 0) {
		x1 = -B/2/A + math.Sqrt(D)/2/A
		x2 = -B/2/A - math.Sqrt(D)/2/A
		err = nil
		return
	} else {
		err = errors.New("D<0")
		return
	}
}

func QuadraticComplex(A, B, C float64) (x1, x2 complex128, err error) {
	Ai := complex(A, 0)
	Bi := complex(B, 0)
	Ci := complex(C, 0)
	sqrtD := cmplx.Sqrt(Bi * Bi * -4 * Ai * Ci)
	x1 = (-Bi + sqrtD) / (2 * Ai)
	x2 = (-Bi - sqrtD) / (2 * Ai)
	err = nil
	return
}
