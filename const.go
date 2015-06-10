package xlUtil_go

const (
	MAX_INT  = int(^uint(0) >> 1)
	MIN_INT  = -(MAX_INT - 1)
	MAX_UINT = ^uint(0)
	MIN_UINT = 0

	//               ....x....1....x....2....x....3....x....4
	SHA1_HEX_NONE = "0000000000000000000000000000000000000000"

	//               ....x....1....x....2....x....3....x....4....x....5....x....6....
	SHA2_HEX_NONE = "0000000000000000000000000000000000000000000000000000000000000000"
	SHA3_HEX_NONE = "0000000000000000000000000000000000000000000000000000000000000000"

	// length of hashes in hexadecimal
	SHA1_HEX_LEN = 40 // bytes
	SHA2_HEX_LEN = 64 // bytes
	SHA3_HEX_LEN = 64 // bytes

	// length of binary hashes
	SHA1_BIN_LEN = 20 // bytes
	SHA2_BIN_LEN = 32 // bytes
	SHA3_BIN_LEN = 32 // bytes

	// table of supported hashes
	USING_SHA1 = 1
	USING_SHA2 = 2
	USING_SHA3 = 3

	// we will need a table of string representations ("sha1", etc)
	// XXX STUB

)

var (
	SHA1_BIN_NONE = make([]byte, SHA1_BIN_LEN)
	SHA2_BIN_NONE = make([]byte, SHA2_BIN_LEN)
	SHA3_BIN_NONE = make([]byte, SHA3_BIN_LEN)
)
