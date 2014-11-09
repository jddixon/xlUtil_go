package math

import (
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
)

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

	c.Assert(NextExp2_32(0), Equals, uint8(0))
	c.Assert(NextExp2_32(1), Equals, uint8(0))
	c.Assert(NextExp2_32(2), Equals, uint8(1))
	c.Assert(NextExp2_32(7), Equals, uint8(3))
	c.Assert(NextExp2_32(8), Equals, uint8(3))
	c.Assert(NextExp2_32(9), Equals, uint8(4))
	c.Assert(NextExp2_32(1023), Equals, uint8(10))
	c.Assert(NextExp2_32(1024), Equals, uint8(10))
	c.Assert(NextExp2_32(1025), Equals, uint8(11))

	// brute force test of all powers of 2
	n := uint32(1)
	for i := uint8(0); i < uint8(32); i++ {
		c.Assert(NextExp2_32(n), Equals, i)
		n = n << 1
	}
	// quasi-random tests of values in the range [3, 2^32)
	for i := 0; i < 16; i++ {
		exp := uint8(rng.Intn(32)) // so 0..31 inclusive
		flag := uint32(1)          // becomes 1 followed by zero or more zeroes
		for i := uint8(0); i < exp; i++ {
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
