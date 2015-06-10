package nlhtree

// nlhTree_test.go

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
	"hash"
	"os"
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
