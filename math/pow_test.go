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

}
