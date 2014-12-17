package merkletree

const (

	// version number no longer consistent with the Python version info

	//           ....x....1....x....2....x....3....x....4....x....5....x....6....
	SHA1_NONE = "0000000000000000000000000000000000000000"
	SHA2_NONE = "0000000000000000000000000000000000000000000000000000000000000000"
	SHA3_NONE = "0000000000000000000000000000000000000000000000000000000000000000"
	SHA1_LEN  = 20 // bytes
	SHA2_LEN  = 32 // bytes
	SHA3_LEN  = 32 // bytes

	USING_SHA1 = 1
	USING_SHA2 = 2
	USING_SHA3 = 3
)
