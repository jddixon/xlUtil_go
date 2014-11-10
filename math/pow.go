package math

// xlUtil_go/math/pow.go

import (
	"fmt"
	"math"
)

var _ = fmt.Printf

// Return the smallest number k which is a power of two and greater than
// or equal to n.
//
// XXX DEPRECATED
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
//
// XXX DEPRECATED
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
	LOG_TABLE_256 = []int{
		-1, 0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3,
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

// Return the smallest non-negative integer exp where 2^exp is
// greater than or equal to n.
func NextExp2_32(n uint32) (exp int) {

	if n == 0 {
		exp = 0
	} else {
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
		isPow2 := (n & (n - 1)) == 0
		if !isPow2 {
			exp++
		}
	}
	return exp
}

// Return the smallest non-negative integer exp where 2^exp is
// greater than or equal to n.
func NextExp2_64(n uint64) (exp int) {

	if n == 0 {
		exp = 0
	} else {
		test32 := n >> 32
		if test32 > 0 {
			// DEBUG
			// END
			test16 := test32 >> 16
			if test16 > 0 {
				test8 := test16 >> 8
				if test8 > 0 {
					exp = 56 + LOG_TABLE_256[test8]
				} else {
					exp = 48 + LOG_TABLE_256[test16]
				}
			} else {
				test8 := test32 >> 8
				if test8 > 0 {
					exp = 40 + LOG_TABLE_256[test8]
				} else {
					exp = 32 + LOG_TABLE_256[test32]
				}
			}
		} else {
			// high order 32 bits are 0
			test16 := n >> 16
			if test16 > 0 {
				test8 := test16 >> 8
				if test8 > 0 {
					exp = 24 + LOG_TABLE_256[test8]
				} else {
					exp = 16 + LOG_TABLE_256[test16]
				}
			} else {
				// high order 48 bits are 0
				test8 := n >> 8
				if test8 > 0 {
					exp = 8 + LOG_TABLE_256[test8]
				} else {
					exp = LOG_TABLE_256[n]
				}
			}
		}

		isPow2 := (n & (n - 1)) == 0
		if !isPow2 {
			exp++
		}
	}
	// DEBUG
	//fmt.Printf("for n = uint64 %016x, next exp of 2 is %d\n", n, exp)
	// END
	return exp
}

// Return the smallest 32-bit number k which is a power of two and greater
// than // or equal to n.
func NextPow2_32(n uint32) (k uint32) {

	if n == 0 {
		k = uint32(1)
	} else if (n & (n - 1)) == 0 {
		return n // is a power of 2
	} else {
		k = n - 1
		k |= k >> 1
		k |= k >> 2
		k |= k >> 4
		k |= k >> 8
		k |= k >> 16
		k++
	}
	return k
}

// Return the smallest 64-bit number k which is a power of two and greater
// than or equal to n.
func NextPow2_64(n uint64) (k uint64) {

	if n == 0 {
		k = uint64(1)
	} else if (n & (n - 1)) == 0 {
		return n // is a power of 2
	} else {
		k = n - 1
		k |= k >> 1
		k |= k >> 2
		k |= k >> 4
		k |= k >> 8
		k |= k >> 16
		k |= k >> 32
		k++
	}
	return k
}
