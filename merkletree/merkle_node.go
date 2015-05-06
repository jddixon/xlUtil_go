package merkletree

// xlattice_go/util/merkletree/merkle_node.go

import (
	"bytes"
	"fmt"
	xu "github.com/jddixon/xlUtil_go"
)

var _ = fmt.Print

type MerkleNodeI interface {
	Name() string
	GetHash() []byte
	SetHash([]byte) error
	WhichSHA() int
	IsLeaf() bool

	Equal(any interface{}) bool
	ToString(indent, deltaIndent string) (string, error)
	ToStrings(indent, deltaIndent string, ss *[]string) error
	// XXX DELAY THESE FOR A WHILE
	// GetPath()		        string
	// SetPath(value	string) error
}

type MerkleNode struct {
	name     string
	hash     []byte
	whichSHA int
}

func NewMerkleNode(name string, hash []byte, whichSHA int) (
	mn *MerkleNode, err error) {

	if name == "" {
		err = EmptyName
	}
	if err == nil {
		length := len(hash)
		if length != 0 && length != xu.SHA1_BIN_LEN && length != xu.SHA2_BIN_LEN && length != xu.SHA3_BIN_LEN {
			err = InvalidHashLength
		}
	}
	if err == nil {
		mn = &MerkleNode{
			name:     name,
			hash:     hash,
			whichSHA: whichSHA,
		}
	}
	return
}
func (mn *MerkleNode) Name() string {
	return mn.name
}

// XXX THIS IS A MAJOR CHANGE FROM THE PYTHON, where the hash is a
// hex value
func (mn *MerkleNode) GetHash() []byte {
	return mn.hash
}
func (mn *MerkleNode) SetHash(value []byte) (err error) {
	// XXX SOME VALIDATION NEEDED
	mn.hash = value
	return
}
func (mn *MerkleNode) WhichSHA() int {
	return mn.whichSHA
}
func (mn *MerkleNode) Equal(any interface{}) bool {
	if any == mn {
		return true
	}
	if any == nil {
		return false
	}
	switch v := any.(type) {
	case *MerkleNode:
		_ = v
	default:
		return false
	}
	other := any.(*MerkleNode) // type assertion

	if mn.name != other.name {
		return false
	}
	if !bytes.Equal(mn.hash, other.hash) {
		return false
	}
	if mn.whichSHA != other.whichSHA {
		return false
	}
	return true
}

// SERIALIZATION  for debugging
func (mn *MerkleNode) ToString(indent string) (s string, err error) {
	s = fmt.Sprintf("MerkleNode: name %s hash %x sha%d\n",
		mn.name, mn.hash, mn.whichSHA)
	return
}
