package xlUtil_go

// popCount_test.go

import (
	"code.google.com/p/intmath/intgr"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
)

var _ = fmt.Print

func slowBitCount32(n uint32) uint {
	var count uint32
	for i := 0; i < 32; i++ {
		count += n & 1
		n >>= 1
	}
	return uint(count)
}

func slowBitCount64(n uint64) uint {
	var count uint64
	for i := 0; i < 64; i++ {
		count += n & 1
		n >>= 1
	}
	return uint(count)
}

func (s *XLSuite) TestSWAR32(c *C) {

	rng := xr.MakeSimpleRNG()
	for i := 0; i < 8; i++ {
		n := uint32(rng.Int63())
		slowCount := slowBitCount32(n)
		intgrCount := uint(intgr.BitCount(int(n)))
		swar32Count := BitCount32(n)
		c.Assert(swar32Count, Equals, intgrCount)
		c.Assert(swar32Count, Equals, slowCount)
	}
}

func (s *XLSuite) TestSWAR64(c *C) {

	rng := xr.MakeSimpleRNG()
	for i := 0; i < 8; i++ {
		n1 := uint64(rng.Int63())
		n2 := uint64(rng.Int63())
		n := (n1 << 32) ^ n2 // we want a full 64 random bits
		slowCount := slowBitCount64(n)
		intgrCount := uint(intgr.BitCount(int(n)))
		swar64Count := BitCount64(n)
		c.Assert(swar64Count, Equals, intgrCount)
		c.Assert(swar64Count, Equals, slowCount)
	}
}

// XXX INCLUDES TEST OF EQUIVALENCE WITH DEPRECATED FUNCTIONS ///////

func (s *XLSuite) TestCounts(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_COUNTS")
	}
	rng := xr.MakeSimpleRNG()

	for i := 0; i < 8; i++ {
		x := uint64(rng.Int63())

		// these operate on uint64s, not BitMap64s
		countX := BitCount64(x)

		// test the BitMap64 operation
		b := NewBitMap64(x)
		c.Assert(b.Count(), Equals, countX)

		// XXX test the deprecated functions
		count3 := popCount3(x)
		count4 := popCount4(x)
		c.Assert(count3, Equals, countX)
		c.Assert(count4, Equals, countX)
	}
}
