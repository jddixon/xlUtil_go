package xlUtil_go

// xlUtil_go/errors.go

import (
	e "errors"
)

var (
	InvalidName           = e.New("not a valid entity name")
	TooManyPartsInVersion = e.New("too many parts in version")
	WrongLengthForVersion = e.New("wrong length for version")
)
