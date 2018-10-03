// +build OMIT

package main

import (
	"fmt"
	"math"
)

//%g - %e for large exponents, %f otherwise. Precision is discussed below (%e - scientific notation, e.g. -1.234456e+78; %f - decimal point but no exponent, e.g. 123.456)
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
