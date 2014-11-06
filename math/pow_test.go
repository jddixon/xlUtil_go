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
		fmt.Println("TEST_POW2")
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
