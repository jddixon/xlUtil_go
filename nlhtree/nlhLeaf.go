package nlhtree

// xlUtil_go/nlhtree/nlhLeaf.go

import (
	"bytes"
	"encoding/hex"
	"fmt"
	xu "github.com/jddixon/xlU_go"
	"os"
)

type NLHLeaf struct {
	hash []byte
	NLHNode
}

func NewNLHLeaf(name string, hash []byte) (nl *NLHLeaf, err error) {
	usingSHA1, err := CheckHash(hash)
	if err == nil {
		nn := NewNLHNode(name, usingSHA1)
		nl = &NLHLeaf{
			hash:    hash,
			NLHNode: *nn,
		}
	}
	return
}

func (nl *NLHLeaf) HexHash() string {
	return hex.EncodeToString(nl.hash)
}

func (nl *NLHLeaf) BinHash() []byte {
	return nl.hash // XXX NOT A CLONE
}

func (nl *NLHLeaf) IsLeaf() bool {
	return true
}

func (nl *NLHLeaf) Equal(any interface{}) bool {
	if any == nl {
		return true
	}
	if any == nil {
		return false
	}
	switch v := any.(type) {
	case *NLHLeaf:
		_ = v
	default:
		return false
	}
	other := any.(*NLHLeaf)
	if other.Name() != nl.Name() {
		return false
	}
	if !bytes.Equal(other.BinHash(), nl.BinHash()) {
		return false
	}
	return true
}

func (nl *NLHLeaf) ToString(indent int) string {
	return fmt.Sprintf("%s%s %s",
		GetSpaces(indent),
		nl.Name(),
		nl.HexHash())
}

// Create an NLHLeaf from the contents of the file at **path**.
// The name is part of the path but is passed to simplify the code.
// Returns nil if the file cannot be found.
func CreateNLHLeafFromFileSystem(path, name string, usingSHA1 bool) (
	nl *NLHLeaf, err error) {

	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			err = IsDirectory
		} else {
			var binHash []byte
			if usingSHA1 {
				binHash, err = xu.FileBinSHA1(path)
			} else {
				binHash, err = xu.FileBinSHA2(path)
			}
			if err == nil {
				nl, err = NewNLHLeaf(name, binHash)
			}
		}
	}
	return
}
