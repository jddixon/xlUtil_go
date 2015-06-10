package nlhtree

// nlhLeaf_test.go

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
	"hash"
)

var _ = fmt.Print

func doTestSimpleConstructor(c *C, rng *xr.PRNG, usingSHA1 bool) {

	var sha hash.Hash
	if usingSHA1 {
		sha = sha1.New()
	} else {
		sha = sha256.New()
	}

	name := rng.NextFileName(8)
	n := rng.SomeBytes(8)
	rng.NextBytes(n)
	sha.Write(n)
	hash0 := sha.Sum(nil)

	leaf0, err := NewNLHLeaf(name, hash0)
	c.Assert(err, IsNil)
	c.Assert(name, Equals, leaf0.Name())
	c.Assert(hash0, Equals, leaf0.BinHash())

	name2 := name
	for name2 == name {
		name2 = rng.NextFileName(8)
	}
	n = rng.SomeBytes(8)
	rng.NextBytes(n)
	sha.Write(n)
	hash1 := sha.Sum(nil)
	leaf1, err := NewNLHLeaf(name2, hash1)
	c.Assert(err, IsNil)
	c.Assert(name2, Equals, leaf1.Name())
	c.Assert(hash1, Equals, leaf1.BinHash())

	c.Assert(leaf0, Equals, leaf0)
	c.Assert(leaf1, Equals, leaf1)
	c.Assert(leaf0.Equal(leaf1), Equals, false)
}
func testSimplestConstructor(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestSimpleConstructor(c, rng, true)
	doTestSimpleConstructor(c, rng, false)
}
