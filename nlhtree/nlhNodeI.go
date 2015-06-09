package nlhtree

// xlUtil_go/nlhtree/nlhNodeI.go

// NLHNodeI ---------------------------------------------------------

type NLHNodeI interface {
	Name() string
	UsingSHA1() bool
	IsLeaf() bool
}
