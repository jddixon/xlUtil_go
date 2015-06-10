package nlhtree

// xlUtil_go/nlhtree/nlhBase.go

import (
	"fmt"
	"strings"
)

var _ = fmt.Print

type NLHBase struct {
	name      string
	root      *NLHTree
	curTree   *NLHTree // initialized to Root
	usingSHA1 bool
}

func NewNLHBase(name string, usingSHA1 bool) (nb *NLHBase) {
	root := NewNLHTree(name, usingSHA1)
	nb = &NLHBase{
		name:      name,
		root:      root,
		usingSHA1: usingSHA1,
	}
	nb.curTree = nb.root
	return
}

func (nb *NLHBase) Name() string {
	return nb.name
}
func (nb *NLHBase) Root() *NLHTree {
	return nb.root
}
func (nb *NLHBase) UsingSHA1() bool {
	return nb.usingSHA1
}
func (nb *NLHBase) CurTree() *NLHTree {
	return nb.curTree
}
func (nb *NLHBase) SetCurTree(path string) (tree *NLHTree, err error) {
	if path == "" {
		err = CurTreeToEmpty
	} else {
		path = strings.TrimSpace(path)
		if path == "" {
			err = CurTreeToEmpty
		}
		// XXX handle paths with and without internal slashes
		tree = nb.curTree
	}
	return
}
