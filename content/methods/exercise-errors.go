// +build OMIT

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//%g	%e for large exponents, %f otherwise. Precision is discussed below.
	//%e	scientific notation, e.g. -1.234456e+78
	//%f	decimal point but no exponent, e.g. 123.456
	return fmt.Sprintf("cannot Sqrt negative number %g", e)
}

const delta = 1e-6

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := x
		n := 0.0
		for math.Abs(n-z) > delta {
			n, z = z, z-(z*z-x)/(2*z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
