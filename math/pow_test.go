package math

import (
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
)

// XXX DEPRECATED
//
func (s *XLSuite) TestPow2(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_POW2")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextPow2(0), Equals, uint(1))
	c.Assert(NextPow2(1), Equals, uint(1))
	c.Assert(NextPow2(2), Equals, uint(2))
	c.Assert(NextPow2(7), Equals, uint(8))
	c.Assert(NextPow2(8), Equals, uint(8))
	c.Assert(NextPow2(9), Equals, uint(16))
	c.Assert(NextPow2(1023), Equals, uint(1024))
	c.Assert(NextPow2(1024), Equals, uint(1024))
	c.Assert(NextPow2(1025), Equals, uint(2048))
}

// XXX DEPRECATED
//
func (s *XLSuite) TestExp2(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_EXP2")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextExp2(0), Equals, uint(0))
	c.Assert(NextExp2(1), Equals, uint(0))
	c.Assert(NextExp2(2), Equals, uint(1))
	c.Assert(NextExp2(7), Equals, uint(3))
	c.Assert(NextExp2(8), Equals, uint(3))
	c.Assert(NextExp2(9), Equals, uint(4))
	c.Assert(NextExp2(1023), Equals, uint(10))
	c.Assert(NextExp2(1024), Equals, uint(10))
	c.Assert(NextExp2(1025), Equals, uint(11))

}
func (s *XLSuite) TestExp2_32(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_EXP2_32")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextExp2_32(0), Equals, 0)
	c.Assert(NextExp2_32(1), Equals, 0)
	c.Assert(NextExp2_32(2), Equals, 1)
	c.Assert(NextExp2_32(7), Equals, 3)
	c.Assert(NextExp2_32(8), Equals, 3)
	c.Assert(NextExp2_32(9), Equals, 4)
	c.Assert(NextExp2_32(1023), Equals, 10)
	c.Assert(NextExp2_32(1024), Equals, 10)
	c.Assert(NextExp2_32(1025), Equals, 11)

	// brute force test of all powers of 2
	n := uint32(1)
	for i := 0; i < 32; i++ {
		c.Assert(NextExp2_32(n), Equals, i)
		n = n << 1
	}
	// quasi-random tests of values in the range [3, 2^32)
	for i := 0; i < 16; i++ {
		exp := rng.Intn(32) // so 0.. inclusive
		flag := uint32(1)   // becomes 1 followed by zero or more zeroes
		for i := 0; i < exp; i++ {
			flag <<= 1
		}
		var lowBits uint32
		if flag > uint32(1) {
			lowBits = uint32(rng.Int63n(int64(flag)))
		}
		n = flag + lowBits
		if lowBits == uint32(0) {
			c.Assert(NextExp2_32(n), Equals, exp)
		} else {
			c.Assert(NextExp2_32(n), Equals, exp+1)
		}
	}
}
func (s *XLSuite) TestExp2_64(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_EXP2_64")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextExp2_64(0), Equals, 0)
	c.Assert(NextExp2_64(1), Equals, 0)
	c.Assert(NextExp2_64(2), Equals, 1)
	c.Assert(NextExp2_64(7), Equals, 3)
	c.Assert(NextExp2_64(8), Equals, 3)
	c.Assert(NextExp2_64(9), Equals, 4)
	c.Assert(NextExp2_64(1023), Equals, 10)
	c.Assert(NextExp2_64(1024), Equals, 10)
	c.Assert(NextExp2_64(1025), Equals, 11)

	// brute force test of all powers of 2
	n := uint64(1)
	for i := 0; i < 64; i++ {
		c.Assert(NextExp2_64(n), Equals, i)
		n = n << 1
	}
	// quasi-random tests of values in the range [3, 2^63), restricted
	// to 63 bits because we use Int63n() to generate lowBits
	for i := 0; i < 16; i++ {
		exp := rng.Intn(63) // so 0..62 inclusive
		flag := uint64(1)   // becomes 1 followed by zero or more zeroes
		for i := 0; i < exp; i++ {
			flag <<= 1
		}
		var lowBits uint64
		if flag > uint64(1) {
			lowBits = uint64(rng.Int63n(int64(flag)))
		}
		n = flag + lowBits
		if lowBits == uint64(0) {
			c.Assert(NextExp2_64(n), Equals, exp)
		} else {
			c.Assert(NextExp2_64(n), Equals, exp+1)
		}
	}
}

func (s *XLSuite) TestPow2_32(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_POW2_32")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextPow2_32(0), Equals, uint32(1))
	c.Assert(NextPow2_32(1), Equals, uint32(1))
	c.Assert(NextPow2_32(2), Equals, uint32(2))
	c.Assert(NextPow2_32(7), Equals, uint32(8))
	c.Assert(NextPow2_32(8), Equals, uint32(8))
	c.Assert(NextPow2_32(9), Equals, uint32(16))
	c.Assert(NextPow2_32(1023), Equals, uint32(1024))
	c.Assert(NextPow2_32(1024), Equals, uint32(1024))
	c.Assert(NextPow2_32(1025), Equals, uint32(2048))

	// Test all powers of 2
	n := uint32(4)
	for i := 2; i < 32; i++ {
		c.Assert(NextPow2_32(n), Equals, n)
		n <<= 1
	}

	// Test random values
	n = uint32(4)
	for i := 2; i < 31; i++ {
		lowBits := uint32(rng.Int31n(int32(n)))
		if lowBits != uint32(0) {
			c.Assert(NextPow2_32(n+lowBits), Equals, 2*n)
		}
		n <<= 1
	}
}

func (s *XLSuite) TestPow2_64(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_POW2_64")
	}
	rng := xr.MakeSimpleRNG()
	_ = rng

	c.Assert(NextPow2_64(0), Equals, uint64(1))
	c.Assert(NextPow2_64(1), Equals, uint64(1))
	c.Assert(NextPow2_64(2), Equals, uint64(2))
	c.Assert(NextPow2_64(7), Equals, uint64(8))
	c.Assert(NextPow2_64(8), Equals, uint64(8))
	c.Assert(NextPow2_64(9), Equals, uint64(16))
	c.Assert(NextPow2_64(1023), Equals, uint64(1024))
	c.Assert(NextPow2_64(1024), Equals, uint64(1024))
	c.Assert(NextPow2_64(1025), Equals, uint64(2048))

	// Test all powers of 2
	n := uint64(4)
	for i := 2; i < 64; i++ {
		c.Assert(NextPow2_64(n), Equals, n)
		n <<= 1
	}

	// Test random values
	n = uint64(4)
	for i := 2; i < 63; i++ {
		lowBits := uint64(rng.Int63n(int64(n)))
		if lowBits != uint64(0) {
			c.Assert(NextPow2_64(n+lowBits), Equals, 2*n)
		}
		n <<= 1
	}
}
