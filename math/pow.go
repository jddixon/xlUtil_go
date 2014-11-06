package math

// xlUtil_go/math/pow.go

import (
	"fmt"
	"math"
)

var _ = fmt.Printf

// Return the smallest number k which is a power of two and greater than
// or equal to n.
func NextPow2(n uint) (k uint) {

	if n == 0 {
		k = uint(1)
	} else if n == uint(1) ||  n == uint(2) {
		k = n
	} else {
		f := float64(n)
		frac, exp := math.Frexp(f)
		//we are guaranteed that frac is in [1/2, 1)
		if frac == float64(0.5) {
			k = n
		} else {
			k = uint(math.Pow(float64(2), float64(exp)))
		}
	}
	return k
}

// Return the smallest number integer k where 2^k is greater than
// or equal to n.
func NextExp2(n uint) (k uint) {

	if n == 0 || n == uint(1) {
		k = uint(0)
	} else if  n == uint(2) {
		k = uint(1)
	} else {
		f := float64(n)
		frac, exp := math.Frexp(f)
		//we are guaranteed that frac is in [1/2, 1)
		if frac == float64(0.5) {
			k = uint(exp - 1)
		} else {
			k = uint(exp)
		}
	}
	return k
}
