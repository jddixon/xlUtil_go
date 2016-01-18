package merkletree

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	xu "github.com/jddixon/xlUtil_go"
	"golang.org/x/crypto/sha3"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"strings"
)

func (s *XLSuite) doTestSimpleConstructor(c *C, rng *xr.PRNG, whichSHA int) {
	fileName := rng.NextFileName(8)
	leaf1, err := NewMerkleLeaf(fileName, nil, whichSHA)
	c.Assert(err, IsNil)
	c.Assert(leaf1.Name(), Equals, fileName)
	c.Assert(len(leaf1.GetHash()), Equals, 0)
	c.Assert(leaf1.WhichSHA(), Equals, whichSHA)

	fileName2 := rng.NextFileName(8)
	for fileName2 == fileName {
		fileName2 = rng.NextFileName(8)
	}
	leaf2, err := NewMerkleLeaf(fileName2, nil, whichSHA)
	c.Assert(err, IsNil)
	c.Assert(leaf2.Name(), Equals, fileName2)

	c.Assert(leaf1.Equal(leaf1), Equals, true)
	c.Assert(leaf1.Equal(leaf2), Equals, false)
}

func (s *XLSuite) doTestSHA(c *C, rng *xr.PRNG, whichSHA int) {

	var hash, fHash []byte
	var sHash string

	// name guaranteed to be unique
	length, pathToFile := rng.NextDataFile("tmp", 1024, 256)
	data, err := ioutil.ReadFile(pathToFile)
	c.Assert(err, IsNil)
	c.Assert(len(data), Equals, length)
	parts := strings.Split(pathToFile, "/")
	c.Assert(len(parts), Equals, 2)
	fileName := parts[1]

	switch whichSHA {
	case xu.USING_SHA1:
		sha := sha1.New()
		sha.Write(data)
		hash = sha.Sum(nil)
		fHash, err = SHA1File(pathToFile)
	case xu.USING_SHA2:
		sha := sha256.New()
		sha.Write(data)
		hash = sha.Sum(nil)
		fHash, err = SHA2File(pathToFile)
	case xu.USING_SHA3:
		sha := sha3.New256()
		sha.Write(data)
		hash = sha.Sum(nil)
		fHash, err = SHA3File(pathToFile)
		// XXX DEFAULT = ERROR
	}
	c.Assert(err, IsNil)
	c.Assert(bytes.Equal(hash, fHash), Equals, true)

	ml, err := CreateMerkleLeafFromFileSystem(pathToFile, fileName, whichSHA)
	c.Assert(err, IsNil)
	c.Assert(ml.Name(), Equals, fileName)
	c.Assert(bytes.Equal(ml.GetHash(), hash), Equals, true)
	c.Assert(ml.WhichSHA(), Equals, whichSHA)

	// TODO: test ToString
	_ = sHash // TODO

}
func (s *XLSuite) TestMerkleLeaf(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_MERKLE_LEAF")
	}
	rng := xr.MakeSimpleRNG()
	s.doTestSimpleConstructor(c, rng, xu.USING_SHA1)
	s.doTestSimpleConstructor(c, rng, xu.USING_SHA2)
	s.doTestSimpleConstructor(c, rng, xu.USING_SHA3)

}
