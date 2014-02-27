package util

import (
	"fmt"
	"strconv"
	"strings"
)

type DecimalVersion uint32

func New(a, b, c, d uint) (dv DecimalVersion) {
	return DecimalVersion(uint((0xff & a) |
		((0xff & b) << 8) |
		((0xff & c) << 16) |
		((0xff & d) << 24)))
}

// Interpret a byte slice as a big-endian uint.
func VersionFromBytes(b []byte) (dv DecimalVersion, err error) {
	if len(b) != 4 {
		err = WrongLengthForVersion
	} else {
		dv = New(uint(b[0]), uint(b[1]), uint(b[2]), uint(b[3]))
	}
	return
}

// Convert a uint32 DecimalVersion to string format.
func (dv DecimalVersion) String() (s string) {
	val := dv
	a := byte(val)
	val >>= 8
	b := byte(val)
	val >>= 8
	c := byte(val)
	val >>= 8
	d := byte(val)

	if c == 0 {
		if d == 0 {
			s = fmt.Sprintf("%d.%d", a, b)
		} else {
			s = fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
		}
	} else if d == 0 {
		s = fmt.Sprintf("%d.%d.%d", a, b, c)
	} else {
		s = fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
	}
	return
}

// Convert a string like a.b.c.d back to a uint32 DecimalVersion.  At
// least one digit must be present.
func ParseDecimalVersion(s string) (dv DecimalVersion, err error) {

	var val uint32
	s = strings.TrimSpace(s)
	parts := strings.Split(s, `.`)
	if len(parts) > 4 {
		err = TooManyPartsInVersion
	}
	if err == nil {
		for i := uint(0); i < uint(len(parts)); i++ {
			var n uint64
			n, err = strconv.ParseUint(parts[i], 10, 8)
			if err != nil {
				break
			}
			val += uint32(n) << (i * 8)
		}
		dv = DecimalVersion(val)
	}
	return
}
