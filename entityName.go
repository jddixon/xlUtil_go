package xlUtil_go

// xlUtil_go/entityName.go

import (
	"regexp"
)

const (
	NAME_STARTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_"
	NAME_CHARS    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
)

var (
	namePat = "^[" + NAME_STARTERS + "]" + "[" + NAME_CHARS + "]*$"
	nameRE  = regexp.MustCompile(namePat)
)

func ValidEntityName(name string) (err error) {
	if !nameRE.MatchString(name) {
		err = InvalidName
	}
	return
}

func NAME_PAT() string        { return namePat }
func NAME_RE() *regexp.Regexp { return nameRE }
func INVALID_NAME() error     { return InvalidName }
