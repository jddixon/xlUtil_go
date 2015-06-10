package nlhtree

// nlhBase_test.go

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
	"hash"
)

var _ = fmt.Print

func doTestConstructor(c *C, rng *xr.PRNG, usingSHA1 bool) {
	name := rng.NextFileName(8)
	b := NewNLHBase(name, usingSHA1)
	c.Assert(b.Name(), Equals, name)
	c.Assert(b.UsingSHA1(), Equals, usingSHA1)
	root := b.Root()
	ct := b.CurTree()
	c.Assert(root.Name(), Equals, ct.Name())
}
func testConstructor(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestConstructor(c, rng, true)
	doTestConstructor(c, rng, false)
}

func doTestWithSimpleTree(c *C, rng *xr.PRNG, usingSHA1 bool) {
	var sha hash.Hash
	if usingSHA1 {
		sha = sha1.New()
	} else {
		sha = sha256.New()
	}

	// XXX WORKING HERE
	_ = sha
}

func (s *XLSuite) TestSimpletTree(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestWithSimpleTree(c, rng, true)
	doTestWithSimpleTree(c, rng, false)
}
