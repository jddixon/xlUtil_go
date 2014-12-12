package xlUtil_go

import (
	"fmt"
	"strings"
)

// The 32- and 64-bit versions of the SWAR algorithm.  These are variants of
// the code in Bagwell's "Ideal Hash Trees".  The algorithm seems to have been
// created by the aggregate.org/MAGIC group at the University of Kentucky
// earlier than the fall of 1996.  Illmari Karonen (vyznev.net) explains the
// algorithm at
// stackoverflow.com/questions/22081738/how-variable-precision-swar-algorithm-workds

const (
	OCTO_FIVES  = uint32(0x55555555)
	OCTO_THREES = uint32(0x33333333)
	OCTO_ONES   = uint32(0x01010101)
	OCTO_FS     = uint32(0x0f0f0f0f)

	HEXI_FIVES  = uint64(0x5555555555555555)
	HEXI_THREES = uint64(0x3333333333333333)
	HEXI_ONES   = uint64(0x0101010101010101)
	HEXI_FS     = uint64(0x0f0f0f0f0f0f0f0f)
)

func BitCount32(n uint32) uint {
	n = n - ((n >> 1) & OCTO_FIVES)
	n = (n & OCTO_THREES) + ((n >> 2) & OCTO_THREES)
	return uint((((n + (n >> 4)) & OCTO_FS) * OCTO_ONES) >> 24)
}

func BitCount64(n uint64) uint {
	n = n - ((n >> 1) & HEXI_FIVES)
	n = (n & HEXI_THREES) + ((n >> 2) & HEXI_THREES)
	return uint((((n + (n >> 4)) & HEXI_FS) * HEXI_ONES) >> 56)
}

func dumpByteSlice(sl []byte) string {
	var ss []string
	for i := 0; i < len(sl); i++ {
		ss = append(ss, fmt.Sprintf("%02x ", sl[i]))
	}
	return strings.Join(ss, "")
}

// POP_COUNT AND SUCH ///////////////////////////////////////////////

// XXX THE CODE BELOW IS VESTIGIAL AND FULLY REDUNDANT

// See Wikipedia: http://en.wikipedia.org/wiki/Hamming_weight.

const (
	m1  = 0x5555555555555555
	m2  = 0x3333333333333333
	m4  = 0x0f0f0f0f0f0f0f0f
	h01 = 0x0101010101010101
)

// Code suitable for machines with a fast multiply operation.
func popCount3(x uint64) (count uint) {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return uint((x * h01) >> 56)
}

// Better for cases where few bits are non-zero
func popCount4(x uint64) (count uint) {
	for count = uint(0); x != 0; count++ {
		x &= x - 1
	}
	return
}
