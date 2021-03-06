package merkletree

// xlattice_go/util/merkletree/merkletree.go

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	xu "github.com/jddixon/xlUtil_go"
	xf "github.com/jddixon/xlUtil_go/lfs"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
)

var _ = fmt.Print

type MerkleLeaf struct {
	MerkleNode
}

// Creates a MerkleTree leaf node.

func NewMerkleLeaf(name string, hash []byte, whichSHA int) (
	ml *MerkleLeaf, err error) {

	mn, err := NewMerkleNode(name, hash, whichSHA)
	if err == nil {
		ml = &MerkleLeaf{*mn}
	}
	return
}

// Create a MerkleTree leaf node corresponding to a file in the file
// system.  To simplify programming, the base name of the file, which is
// part of the path, is also passed as a separate argument.

func CreateMerkleLeafFromFileSystem(pathToFile, name string, whichSHA int) (
	ml *MerkleLeaf, err error) {

	var hash []byte
	if whichSHA == xu.USING_SHA1 {
		hash, err = SHA1File(pathToFile)
	} else if whichSHA == xu.USING_SHA2 {
		hash, err = SHA2File(pathToFile)
	} else {
		hash, err = SHA3File(pathToFile)
	}
	if err == nil {
		ml, err = NewMerkleLeaf(name, hash, whichSHA)
	}
	return
}
func (ml *MerkleLeaf) IsLeaf() bool {
	return true
}
func (ml *MerkleLeaf) Equal(any interface{}) bool {
	if any == ml {
		return true
	}
	if any == nil {
		return false
	}
	switch v := any.(type) {
	case *MerkleLeaf:
		_ = v
	default:
		return false
	}
	other := any.(*MerkleLeaf) // type assertion
	myNode := &ml.MerkleNode
	return myNode.Equal(&other.MerkleNode)
}

// Serialize the leaf node, prefixing it with 'indent', which should
// conventionally be a number of spaces.

func (ml *MerkleLeaf) ToString(indent, deltaIndent string) (
	str string, err error) {

	var shash string
	hash := ml.hash
	if len(hash) == 0 {
		switch ml.whichSHA {
		case xu.USING_SHA1:
			shash = SHA1_NONE
		case xu.USING_SHA2:
			shash = SHA2_NONE
		case xu.USING_SHA3:
			shash = SHA3_NONE
			// XXX DEFAULT => ERROR
		}
	} else {
		shash = hex.EncodeToString(hash)
	}
	str = fmt.Sprintf("%s%s %s", indent, shash, ml.name)
	return
}

func (ml *MerkleLeaf) ToStrings(indent, deltaIndent string, ss *[]string) (
	err error) {

	str, err := ml.ToString(indent, deltaIndent)
	if err == nil {
		*ss = append(*ss, str)
	}
	return
}

// Return the SHA1 hash of a file.  This is a sequence of 20 bytes.

func SHA1File(pathToFile string) (hash []byte, err error) {
	var data []byte
	found, err := xf.PathExists(pathToFile)
	if err == nil && !found {
		err = FileNotFound
	}
	if err == nil {
		data, err = ioutil.ReadFile(pathToFile)
		if err == nil {
			digest := sha1.New()
			digest.Write(data)
			hash = digest.Sum(nil)
		}
	}
	return
}

// Return the SHA256 hash of a file.  This is a sequence of 32 bytes.

func SHA2File(pathToFile string) (hash []byte, err error) {
	var data []byte
	found, err := xf.PathExists(pathToFile)
	if err == nil && !found {
		err = FileNotFound
	}
	if err == nil {
		data, err = ioutil.ReadFile(pathToFile)
		if err == nil {
			digest := sha256.New()
			digest.Write(data)
			hash = digest.Sum(nil)
		}
	}
	return
}

// Return the SHA3-256 hash of a file.  This is a sequence of 32 bytes.

func SHA3File(pathToFile string) (hash []byte, err error) {
	var data []byte
	found, err := xf.PathExists(pathToFile)
	if err == nil && !found {
		err = FileNotFound
	}
	if err == nil {
		data, err = ioutil.ReadFile(pathToFile)
		if err == nil {
			digest := sha3.New256()
			digest.Write(data)
			hash = digest.Sum(nil)
		}
	}
	return
}
