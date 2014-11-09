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
	} else if n == uint(1) || n == uint(2) {
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
	} else if n == uint(2) {
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

var (
	LOG_TABLE_256 = []uint8{
		0, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
	}
)

// Return the smallest unsigned integer k where 2^k is greater than
// or equal to n.
func NextExp2_32(n uint32) (exp uint8) {

	if n == 0 || n == uint32(1) {
		exp = uint8(0)
	} else {
		isPow2 := (n & (n - 1)) == 0
		_ = isPow2

		test16 := n >> 16
		if test16 > 0 {
			test8 := test16 >> 8
			if test8 > 0 {
				exp = 24 + LOG_TABLE_256[test8]
			} else {
				exp = 16 + LOG_TABLE_256[test16]
			}
		} else {
			test8 := n >> 8
			if test8 > 0 {
				exp = 8 + LOG_TABLE_256[test8]
			} else {
				exp = LOG_TABLE_256[n]
			}
		}
		if !isPow2 {
			exp++
		}
	}
	// DEBUG
	//fmt.Printf("for n = %d, next power of 2 is %d\n", n, exp)
	// END
	return exp
}
