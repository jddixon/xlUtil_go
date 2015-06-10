package nlhtree

// xlUtil_go/nlhtree/nlhNode.go

import (
	"errors"
	xu "github.com/jddixon/xlUtil_go"
)

type NLHNode struct {
	name      string
	usingSHA1 bool
}

func NewNLHNode(name string, usingSHA1 bool) (nn *NLHNode) {
	nn = &NLHNode{
		name:      name,
		usingSHA1: usingSHA1,
	}
	return
}

func (nn *NLHNode) Name() string {
	return nn.name
}
func (nn *NLHNode) UsingSHA1() bool {
	return nn.usingSHA1
}

// func (nn *NLHNode) IsLeaf() bool {
//		NOT IMPLEMENTED
// |

// utility functions ------------------------------------------------

func CheckHash(hash []byte) (isSHA1 bool, err error) {
	// return True if SHA1, False if SHA2, otherwise raise """
	if hash == nil {
		err = errors.New("hash cannot be nil")
	} else {
		hashLen := len(hash)
		if hashLen == xu.SHA1_BIN_LEN {
			isSHA1 = true
		} else if hashLen == xu.SHA2_BIN_LEN {
			isSHA1 = false
		} else {
			err = errors.New("not a valid SHA hash length")
		}
	}
	return
}

var __SPACES__ []string

func GetSpaces(n int) string {
	if n < 0 {
		n = 0
	}
	// cache strings of N spaces
	k := len(__SPACES__) - 1
	for k < n {
		k++
		if k == 0 {
			__SPACES__ = append(__SPACES__, "")
		} else {
			__SPACES__ = append(__SPACES__, __SPACES__[k-1]+" ")
		}
	}
	return __SPACES__[n]
}
