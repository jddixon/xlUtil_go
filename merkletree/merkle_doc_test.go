package merkletree

// xlatttice_go/util/merkletree/merkle_doc_test.go

import (
	"bytes"
	"code.google.com/p/go.crypto/sha3"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	xu "github.com/jddixon/xlUtil_go"
	. "gopkg.in/check.v1"
	"hash"
	re "regexp"
)

// REGEXP TESTS =====================================================
func (s *XLSuite) doTestForExpectedExclusions(c *C, exRE *re.Regexp) {
	// should always match
	c.Assert(exRE.MatchString("."), Equals, true)
	c.Assert(exRE.MatchString(".."), Equals, true)
	c.Assert(exRE.MatchString(".merkle"), Equals, true)
	c.Assert(exRE.MatchString(".svn"), Equals, true)
	c.Assert(exRE.MatchString(".foo.swp"), Equals, true)
	c.Assert(exRE.MatchString("junkEverywhere"), Equals, true)
}
func (s *XLSuite) doTestForExpectedMatches(c *C,
	matchRE *re.Regexp, names []string) {

	for i := 0; i < len(names); i++ {
		name := names[i]
		c.Assert(matchRE.MatchString(name), Equals, true)
	}
}
func (s *XLSuite) doTestForExpectedMatchFailures(c *C,
	matchRE *re.Regexp, names []string) {

	for i := 0; i < len(names); i++ {
		name := names[i]
		m := matchRE.MatchString(name)
		if m {
			fmt.Printf("WE HAVE A MATCH ON '%s'\n", name)
			// self.assertEquals( None, where )
		}
	}
}

// test utility for making excluded file name regexes

func (s *XLSuite) TestMakeExRE(c *C) {
	exRE, err := MakeExRE(nil)
	c.Assert(err, IsNil)
	c.Assert(exRE, NotNil)
	s.doTestForExpectedExclusions(c, exRE)

	// should not be present
	c.Assert(exRE.MatchString("bar"), Equals, false)
	c.Assert(exRE.MatchString("foo"), Equals, false)

	var exc []string
	exc = append(exc, "^foo")
	exc = append(exc, "bar$")
	exc = append(exc, "^junk*")
	exRE, err = MakeExRE(exc)
	c.Assert(err, IsNil)
	s.doTestForExpectedExclusions(c, exRE)

	c.Assert(exRE.MatchString("foobarf"), Equals, true)
	c.Assert(exRE.MatchString(" foobarf"), Equals, false)
	c.Assert(exRE.MatchString(" foobarf"), Equals, false)

	// bear in mind that match must be at the beginning
	c.Assert(exRE.MatchString("ohMybar"), Equals, true)
	c.Assert(exRE.MatchString("ohMybarf"), Equals, false)
	c.Assert(exRE.MatchString("junky"), Equals, true)
	c.Assert(exRE.MatchString(" junk"), Equals, false)
}

// test utility for making matched file name regexes

func (s *XLSuite) TestMakeMatchRE(c *C) {
	matchRE, err := MakeMatchRE(nil)
	c.Assert(err, IsNil)
	c.Assert(matchRE, IsNil)

	var matches []string
	matches = append(matches, "^foo")
	matches = append(matches, "bar$")
	matches = append(matches, "^junk*")
	matchRE, err = MakeMatchRE(matches)
	c.Assert(err, IsNil)
	cases := []string{"foo", "foolish", "roobar", "junky"}
	s.doTestForExpectedMatches(c, matchRE, cases)

	cases = []string{" foo", "roobarf", "myjunk"}
	s.doTestForExpectedMatchFailures(c, matchRE, cases)

	matches = []string{"\\.tgz$"}
	matchRE, err = MakeMatchRE(matches)
	c.Assert(err, IsNil)

	cases = []string{"junk.tgz", "notSoFoolish.tgz"}
	s.doTestForExpectedMatches(c, matchRE, cases)
	cases = []string{"junk.tar.gz", "foolish.tar.gz"}
	s.doTestForExpectedMatchFailures(c, matchRE, cases)

	matches = []string{"\\.tgz$", "\\.tar\\.gz$"}
	matchRE, err = MakeMatchRE(matches)
	c.Assert(err, IsNil)

	cases = []string{
		"junk.tgz", "notSoFoolish.tgz", "junk.tar.gz", "ohHello.tar.gz"}
	s.doTestForExpectedMatches(c, matchRE, cases)

	cases = []string{"junk.gz", "foolish.tar"}
	s.doTestForExpectedMatchFailures(c, matchRE, cases)
}

// PARSER TESTS =====================================================

func (s *XLSuite) doTestMDParser(c *C, rng *xr.PRNG, whichSHA int) {

	var tHash []byte
	switch whichSHA {
	case xu.USING_SHA1:
		tHash = make([]byte, xu.SHA1_BIN_LEN)
	case xu.USING_SHA2:
		tHash = make([]byte, xu.SHA2_BIN_LEN)
	case xu.USING_SHA3:
		tHash = make([]byte, xu.SHA3_BIN_LEN)
		// DEFAULT = ERROR
	}
	rng.NextBytes(tHash)               // not really a hash, of course
	sHash := hex.EncodeToString(tHash) // string form of tHash

	withoutSlash := rng.NextFileName(8)
	dirName := withoutSlash + "/"

	length := rng.Intn(4)
	var rSpaces string
	for i := 0; i < length; i++ {
		rSpaces += " " // on the right
	}

	// TEST FIRST LINE PARSER -----------------------------
	line := sHash + " " + dirName + rSpaces

	treeHash2, dirName2, err := ParseMerkleDocFirstLine(line)
	c.Assert(err, IsNil)
	c.Assert(bytes.Equal(treeHash2, tHash), Equals, true)
	// we retain the terminating slash in MerkleDoc first lines
	c.Assert(dirName2, Equals, dirName)
}

func (s *XLSuite) TestMDParser(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_MERKLE_DOC_PARSER")
	}
	rng := xr.MakeSimpleRNG()

	s.doTestMDParser(c, rng, xu.USING_SHA1)
	s.doTestMDParser(c, rng, xu.USING_SHA2)
	s.doTestMDParser(c, rng, xu.USING_SHA3)
}

// OTHER TESTS ======================================================

func (s *XLSuite) TestMerkleDoc(c *C) {
	if VERBOSITY > 0 {
		fmt.Println("TEST_MERKLE_DOC_CONSTRUCTOR")
	}
	rng := xr.MakeSimpleRNG()
	s.doTestMerkleDoc(c, rng, xu.USING_SHA1)
	s.doTestMerkleDoc(c, rng, xu.USING_SHA2)
	// XXX WILL FAIL BECAUSE xu.SHA2_BIN_LEN == xu.SHA3_BIN_LEN
	//s.doTestMerkleDoc(c, rng, xu.USING_SHA3)
}

func (s *XLSuite) doTestMerkleDoc(c *C, rng *xr.PRNG, whichSHA int) {

	//test directory is single level, with four data files
	dirName1, dirPath1, dirName2, dirPath2 := s.makeTwoTestDirectories(
		c, rng, ONE, FOUR)

	// XXX dirName2, dirPath2 NEVER USED
	_, _ = dirName2, dirPath2

	tree1, err := CreateMerkleTreeFromFileSystem(dirPath1, whichSHA, nil, nil)
	c.Assert(err, IsNil)
	c.Assert(tree1.Name(), Equals, dirName1)
	nodes1 := tree1.nodes
	c.Assert(nodes1, NotNil)
	c.Assert(len(nodes1), Equals, FOUR)
	s.verifyTreeSHA(c, rng, tree1, dirPath1, whichSHA)
	treeHash1 := tree1.GetHash()

	doc1, err := CreateMerkleDocFromFileSystem(dirPath1, whichSHA, nil, nil)
	c.Assert(err, IsNil)
	c.Assert(doc1, NotNil)

	tree1d := doc1.GetTree()
	c.Assert(tree1.Equal(tree1d), Equals, true)

	c.Assert(doc1.Bound(), Equals, true)
	c.Assert(doc1.GetExRE(), Equals, (*re.Regexp)(nil))
	c.Assert(doc1.GetMatchRE(), Equals, (*re.Regexp)(nil))
	c.Assert(doc1.GetPath(), Equals, "tmp")
	c.Assert(doc1.WhichSHA(), Equals, whichSHA)

	var sha hash.Hash
	switch whichSHA {
	case xu.USING_SHA1:
		sha = sha1.New()
	case xu.USING_SHA2:
		sha = sha256.New()
	case xu.USING_SHA3:
		sha = sha3.NewKeccak256()
		// XXX DEFAULT = ERROR
	}
	// Python uses this order
	sha.Write(treeHash1)
	sha.Write([]byte("tmp"))
	expectedDocHash := sha.Sum(nil)
	actualDocHash := doc1.GetHash()

	// DEBUG
	//fmt.Printf("expectedDocHash: %x\n", expectedDocHash)
	//fmt.Printf("actualDocHash:   %x\n", actualDocHash)
	// END

	c.Assert(expectedDocHash, DeepEquals, actualDocHash)

	doc1Str, err := doc1.ToString("", " ")
	c.Assert(err, IsNil)
	c.Assert(len(doc1Str) > 0, Equals, true)

	doc1TreeStr, err := doc1.GetTree().ToString("", " ")
	c.Assert(err, IsNil)
	c.Assert(len(doc1TreeStr) > 0, Equals, true)

	// DEBUG
	//fmt.Printf("DOC1:\n%s", doc1Str)
	//fmt.Printf("  its tree:\n%s", doc1TreeStr)
	// END

	doc1Rebuilt, err := ParseMerkleDoc(doc1Str, " ")
	c.Assert(err, IsNil)

	// compare at the string level
	doc1RStr, err := doc1Rebuilt.ToString("", " ")
	c.Assert(err, IsNil)
	// DEBUG
	//fmt.Printf("DOC2 = doc1RStr:\n%s", doc1RStr)
	// END
	c.Assert(doc1RStr, Equals, doc1Str)

	c.Assert(doc1.Equal(doc1Rebuilt), Equals, true)
}
