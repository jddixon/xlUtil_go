package nlhtree

// nlhTree_test.go

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	xu "github.com/jddixon/xlUtil_go"
	. "gopkg.in/check.v1"
	"hash"
	"os"
	"strings"
)

var _ = fmt.Print

// UTILITIES FOR TEST 1 /////////////////////////////////////////////

func makeLeaf(c *C, rng *xr.PRNG, namesSoFar map[string]bool, usingSHA1 bool) (
	leaf *NLHLeaf) {

	var err error
	var name string
	for {
		name = rng.NextFileName(8)
		if !namesSoFar[name] {
			namesSoFar[name] = true
			break
		}
	}
	n := rng.SomeBytes(8) // 8 quasi-random bytes
	var sha hash.Hash
	if usingSHA1 {
		sha = sha1.New()
	} else {
		sha = sha256.New()
	}
	sha.Write(n)
	leaf, err = NewNLHLeaf(name, sha.Sum(nil))
	c.Assert(err, IsNil)
	return
}

// UTILITIES FOR TEST 2 /////////////////////////////////////////////

func getTwoUniqueDirectoryNames(c *C, rng *xr.PRNG) (dirName1, dirName2 string) {

	dirName1 = rng.NextFileName(8)
	dirName2 = dirName1
	for dirName2 == dirName1 {
		dirName2 = rng.NextFileName(8)
	}
	c.Assert(len(dirName1) > 0, Equals, true)
	c.Assert(len(dirName2) > 0, Equals, true)
	c.Assert(dirName1 != dirName2, Equals, true)
	return
}
func makeOneNamedTestDirectory(c *C, rng *xr.PRNG, name string,
	depth, width int) (dirPath string) {

	dirPath = fmt.Sprintf("tmp/%s", name)
	if _, err := os.Stat(dirPath); err == nil {
		err = os.RemoveAll(dirPath)
		c.Assert(err, IsNil)
	}
	//                                     maxLen, minLen of files (bytes)
	rng.NextDataDir(dirPath, depth, width, 4096, 32)
	return
}
func makeTwoTestDirectories(c *C, rng *xr.PRNG, depth, width int) (
	dirName1, dirPath1, dirName2, dirPath2 string) {

	dirName1 = rng.NextFileName(8)
	dirPath1 = makeOneNamedTestDirectory(c, rng, dirName1, depth, width)

	dirName2 = dirName1
	for dirName2 == dirName1 {
		dirName2 = rng.NextFileName(8)
	}
	dirPath2 = makeOneNamedTestDirectory(c, rng, dirName2, depth, width)

	return
}

// UNIT TESTS 1 /////////////////////////////////////////////////////

func testSimpleTreeConstructor(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestSimpleTreeConstructor(c, rng, true)
	doTestSimpleTreeConstructor(c, rng, false)
}
func doTestSimpleTreeConstructor(c *C, rng *xr.PRNG, usingSHA1 bool) {
	name := rng.NextFileName(8)
	tree := NewNLHTree(name, usingSHA1)
	c.Assert(tree.name, Equals, name)
	c.Assert(tree.usingSHA1, Equals, usingSHA1)
	c.Assert(len(tree.nodes), Equals, 0)
}

// Create 4 leaf nodes with random but unique names.  Insert
// them into a tree, verifying that the resulting sort is correct.
func doTestInsert4Leafs(c *C, rng *xr.PRNG, usingSHA1 bool) {

	var sha hash.Hash
	if usingSHA1 {
		sha = sha1.New()
	} else {
		sha = sha256.New()
	}
	_ = sha // XXX NOT USED

	name := rng.NextFileName(8)
	tree := NewNLHTree(name, usingSHA1)
	leafNames := make(map[string]bool)
	aL := makeLeaf(c, rng, leafNames, usingSHA1)
	bL := makeLeaf(c, rng, leafNames, usingSHA1)
	cL := makeLeaf(c, rng, leafNames, usingSHA1)
	dL := makeLeaf(c, rng, leafNames, usingSHA1)
	c.Assert(len(tree.nodes), Equals, 0)
	tree.Insert(aL)
	c.Assert(len(tree.nodes), Equals, 1)
	tree.Insert(bL)
	c.Assert(len(tree.nodes), Equals, 2)
	tree.Insert(cL)
	c.Assert(len(tree.nodes), Equals, 3)
	tree.Insert(dL)
	c.Assert(len(tree.nodes), Equals, 4)
	// we expect the nodes to be sorted
	for i := 0; i < 3; i++ {
		c.Assert(tree.nodes[i].Name() < tree.nodes[i+1].Name(), Equals, true)
	}

	matches, err := tree.List("*")
	c.Assert(err, IsNil)
	for i := 0; i < len(tree.Nodes()); i++ {
		q := tree.Nodes()[i]
		c.Assert(matches[i], Equals, "  "+q.Name())
	}
	c.Assert(tree.Equal(tree), Equals, true)
}
func testInsert4Leafs(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestInsert4Leafs(c, rng, true)
	doTestInsert4Leafs(c, rng, false)
}

// UNIT TESTS 2 /////////////////////////////////////////////////////

func testPathlessUnboundConstructor1(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestPathlessUnboundConstructor1(c, rng, true)
	doTestPathlessUnboundConstructor1(c, rng, false)
}
func doTestPathlessUnboundConstructor1(c *C, rng *xr.PRNG, usingSHA1 bool) {

	dirName1, dirName2 := getTwoUniqueDirectoryNames(c, rng)

	tree1 := NewNLHTree(dirName1, usingSHA1)
	c.Assert(dirName1, Equals, tree1.name)
	c.Assert(tree1.usingSHA1, Equals, usingSHA1)

	tree2 := NewNLHTree(dirName2, usingSHA1)
	c.Assert(dirName2, Equals, tree2.name)
	c.Assert(tree2.usingSHA1, Equals, usingSHA1)

	c.Assert(tree1.Equal(tree1), Equals, true)
	c.Assert(tree1.Equal(tree2), Equals, false)
	c.Assert(tree1.Equal(nil), Equals, false)
}

func testBoundFlatDirs(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestBoundFlatDirs(c, rng, true)
	doTestBoundFlatDirs(c, rng, false)
}

// test directory is single level, with four data files
func doTestBoundFlatDirs(c *C, rng *xr.PRNG, usingSHA1 bool) {

	dirName1, dirPath1, dirName2, dirPath2 :=
		makeTwoTestDirectories(c, rng, 1, 4)
		// no exRE, matchRE
	tree1, err := CreateNLHTreeFromFileSystem(dirPath1, usingSHA1, "", "")
	c.Assert(err, IsNil)
	c.Assert(dirName1, Equals, tree1.name)
	nodes1 := tree1.Nodes()
	c.Assert(nodes1, NotNil)
	c.Assert(4, Equals, len(nodes1))

	tree2, err := CreateNLHTreeFromFileSystem(dirPath2, usingSHA1, "", "")
	c.Assert(dirName2, Equals, tree2.Name())
	nodes2 := tree2.Nodes()
	c.Assert(nodes2, NotNil)
	c.Assert(4, Equals, len(nodes2))

	c.Assert(tree1.Equal(tree1), Equals, true)
	c.Assert(tree1.Equal(tree2), Equals, false)
	c.Assert(tree1.Equal(nil), Equals, false)
}

func testBoundNeedleDirs1(c *C) {
	rng := xr.MakeSimpleRNG()
	doTestBoundNeedleDirs(c, rng, true)
	doTestBoundNeedleDirs(c, rng, false)
}

// test directories four deep with one data file at the lowest level
func doTestBoundNeedleDirs(c *C, rng *xr.PRNG, usingSHA1 bool) {

	dirName1, dirPath1, dirName2, dirPath2 :=
		makeTwoTestDirectories(c, rng, 4, 1)
	tree1, err := CreateNLHTreeFromFileSystem(dirPath1, usingSHA1, "", "")
	c.Assert(err, IsNil)

	c.Assert(dirName1, Equals, tree1.name)
	nodes1 := tree1.nodes
	c.Assert(nodes1, NotNil)
	c.Assert(1, Equals, len(nodes1))

	tree2, err := CreateNLHTreeFromFileSystem(dirPath2, usingSHA1, "", "")
	c.Assert(err, IsNil)
	c.Assert(dirName2, Equals, tree2.name)
	nodes2 := tree2.nodes
	c.Assert(nodes2, NotNil)
	c.Assert(1, Equals, len(nodes2))

	c.Assert(tree1.Equal(tree1), Equals, true)
	c.Assert(tree1.Equal(tree2), Equals, false)
}

// UNIT TESTS 3 /////////////////////////////////////////////////////

//  adapted from the buildList example 2015-05-22
var (
	EXAMPLE1 = []string{
		"dataDir",
		" data1 bea7383743859a81b84cec8fde2ccd1f3e2ff688",
		" data2 895c210f5203c48c1e3a574a2d5eba043c0ec72d",
		" subDir1",
		"  cb0ece05cbb91501d3dd78afaf362e63816f6757",
		"  da39a3ee5e6b4b0d3255bfef95601890afd80709",
		" subDir2",
		" subDir3",
		"  data31 8cddeb23f9de9da4547a0d5adcecc7e26cb098c0",
		" subDir4",
		"  subDir41",
		"   subDir411",
		"    data41 31c16def9fc4a4b6415b0b133e156a919cf41cc8",
		" zData 31c16def9fc4a4b6415b0b133e156a919cf41cc8",
	}
	//  this is just a hack but ...
	EXAMPLE2 = []string{
		"dataDir",
		" data1 012345678901234567890123bea7383743859a81b84cec8fde2ccd1f3e2ff688",
		" data2 012345678901234567890123895c210f5203c48c1e3a574a2d5eba043c0ec72d",
		" subDir1",
		"  data11 012345678901234567890123cb0ece05cbb91501d3dd78afaf362e63816f6757",
		"  data12 012345678901234567890123da39a3ee5e6b4b0d3255bfef95601890afd80709",
		" subDir2",
		" subDir3",
		"  data31 0123456789012345678901238cddeb23f9de9da4547a0d5adcecc7e26cb098c0",
		" subDir4",
		"  subDir41",
		"   subDir411",
		"    data41 01234567890123456789012331c16def9fc4a4b6415b0b133e156a919cf41cc8",
		" zData 01234567890123456789012331c16def9fc4a4b6415b0b133e156a919cf41cc8",
	}
)

func testSpaces(c *C) {
	rng := xr.MakeSimpleRNG()
	for i := 0; i < 4; i++ {
		j := rng.Intn(32)
		spaces := GetSpaces(j)
		c.Assert(len(spaces), Equals, j)
		for k := 0; k < len(spaces); i++ {
			c.Assert(spaces[k], Equals, " ") // XXX ???
		}
	}
}

func doTestPatternMatching(c *C, usingSHA1 bool) {
	var ss []string
	if usingSHA1 {
		ss = EXAMPLE1
	} else {
		ss = EXAMPLE2
	}
	//  first line --------------------------------------
	m := DIR_LINE_RE.FindStringSubmatch(ss[0])
	c.Assert(m, NotNil)
	c.Assert(len(m[1]), Equals, 0)
	c.Assert(m[2], Equals, "dataDir")

	//  simpler approach ----------------------
	name, err := ParseNLHTreeFirstLine(ss[0])
	c.Assert(err, IsNil)
	c.Assert(name, Equals, "dataDir")

	//  file with indent of 1 ---------------------------
	if usingSHA1 {
		m = FILE_LINE_RE_1.FindStringSubmatch(ss[1])
	} else {
		m = FILE_LINE_RE_2.FindStringSubmatch(ss[1])
	}
	c.Assert(m, NotNil)
	c.Assert(len(m[1]), Equals, 1)
	c.Assert(m[2], Equals, "data1")

	//  that simpler approach -----------------
	indent, name, hash, err := ParseNLHTreeOtherLine(ss[1])
	c.Assert(err, IsNil)
	c.Assert(indent, Equals, 1)
	c.Assert(name, Equals, "data1")
	if usingSHA1 {
		c.Assert(len(hash), Equals, xu.SHA1_HEX_LEN)
	} else {
		c.Assert(len(hash), Equals, xu.SHA2_HEX_LEN)
	}
	//  subdirectory ------------------------------------
	m = DIR_LINE_RE.FindStringSubmatch(ss[3])
	c.Assert(m, NotNil)
	c.Assert(len(m[1]), Equals, 1)
	c.Assert(m[2], Equals, "subDir1")

	//  that simpler approach -----------------
	indent, name, hash, err = ParseNLHTreeOtherLine(ss[3])
	c.Assert(err, IsNil)
	c.Assert(indent, Equals, 1)
	c.Assert(name, Equals, "subDir1")
	c.Assert(hash, Equals, nil)

	//  lower level file ----------------------
	if usingSHA1 {
		m = FILE_LINE_RE_1.FindStringSubmatch(ss[12])
	} else {
		m = FILE_LINE_RE_2.FindStringSubmatch(ss[12])
	}
	c.Assert(m, NotNil)
	c.Assert(len(m[1]), Equals, 4)
	c.Assert(m[2], Equals, "data41")

	//  that simpler approach -----------------
	indent, name, hash, err = ParseNLHTreeOtherLine(ss[12])
	c.Assert(err, IsNil)
	c.Assert(indent, Equals, 4)
	c.Assert(name, Equals, "data41")
	if usingSHA1 {
		c.Assert(len(hash), Equals, xu.SHA1_HEX_LEN)
	} else {
		c.Assert(len(hash), Equals, xu.SHA2_HEX_LEN)
	}
}

func testPatternMatching(c *C) {
	doTestPatternMatching(c, true)
	doTestPatternMatching(c, false)
}

func doTestSerialization(c *C, usingSHA1 bool) {
	var tree, tree2, tree3 *NLHTree
	var err error
	if usingSHA1 {
		tree, err = CreateNLHTreeFromStringArray(EXAMPLE1, usingSHA1)
	} else {
		tree, err = CreateNLHTreeFromStringArray(EXAMPLE2, usingSHA1)
	}
	c.Assert(err, IsNil)
	c.Assert(tree.usingSHA1, Equals, usingSHA1)

	var ss []string
	tree.ToStrings(ss, 0)

	tree2, err = CreateNLHTreeFromStringArray(ss, usingSHA1)
	c.Assert(err, IsNil)
	c.Assert(tree.Equal(tree2), Equals, true)

	s := strings.Join(ss, "\n") + "\n"
	tree3, err = ParseNLHTree(s, usingSHA1)
	c.Assert(err, IsNil)
	s3 := tree3.String()

	c.Assert(s3, Equals, s)
	c.Assert(tree3.Equal(tree), Equals, true)
}
func testSerialization(c *C) {
	doTestSerialization(c, true)
	doTestSerialization(c, false)
}
