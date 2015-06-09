package nlhtree

import (
	e "errors"
)

var (
	CurTreeToEmpty = e.New("curTree may not be set to empty path")
	IsDirectory    = e.New("file is a directory")
)
