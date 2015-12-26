# xlUtil_go

Utilities useful or necessary for the
[xlattice_go](https://jddixon.github.io/xlattice_go)
project.  These currently include

* **BitMap64**, functions for creating and manipulating 64-bit bit maps
* **DecimalVersion**, a class supporting the representation of version
  numbers for XLattice project components
* **entityName**, functions specifying what is or is not a valid XLattice name
* **lfs**, functions standardizing the handling of the **local file system**
* **math**, routines returning the next power of 2 and the exponent of the
  next power of 2
* **MerkleTree**, a standard way of representing a directory structure
  in terms of file names and content hashes
* **NLHTree**, a somewhat simpler way of representing directory structures
  and files
* **popCount**, 32- and 64-bit versions of the SWAR algorithm used for
  counting the number of bits set in a machine word
* **TimeStamp**, the XLattice standard form for a timestamp

## BitMap64

As the term is used here, a `bitmap` is an array of bits, specifically
a 64-bit word, where the bit being set (having a value of 1) means that
something of interest is present whereas the bit being clear (having a
value of 0) means that it is absent.

	type BitMap64 struct {
		Bits uint64
	}
	
	func NewBitMap64(bits uint64) (bm *BitMap64)
	
	// An array of a bit maps with the low order N bits set.
	// In a better world this would be a constant.
	var lowNMap = [...]uint64
	
	// Returns a bit map with the low order N bits set.  If N is 0, the map
	// is empty - that is, no bits are set.
	func LowNMap(n uint) (bm *BitMap64)
	
	// Return true if all bits are set
	func (bm *BitMap64) All() bool
	
	// Return true if any bits are set
	func (bm *BitMap64) Any() bool
	
	// Clear bit N, setting it to zero.
	func (bm *BitMap64) Clear(n uint) *BitMap64
	
	// Set the low order N bits to zero.
	func (bm *BitMap64) ClearLowN(n uint) *BitMap64
	
	// Return a clone of this bit map.
	func (bm *BitMap64) Clone() *BitMap64
	
	// Returns a bit map in which all of the bits in this map have been flipped.
	func (bm *BitMap64) Complement() *BitMap64
	
	// Returns a count of the bits set (that is, equal to 1) in the bit map.
	func (bm *BitMap64) Count() uint
	
	// Returns the difference between two bit maps.
	func (bm *BitMap64) Difference(other *BitMap64) *BitMap64
	
	// Whether 'other' is a bit map and has the same bits set.
	func (bm *BitMap64) Equal(any interface{}) bool
	
	// Flip the Nth bit in the map, where 0 <= n <= 63
	func (bm *BitMap64) Flip(n uint) *BitMap64
	
	// Return the intersection of the two bit maps -- that is, a map
	// in which all of the bits set in both input sets are set.
	func (bm *BitMap64) Intersection(other *BitMap64) *BitMap64
	
	// Return whether none of the bits in the map is set.
	func (bm *BitMap64) None() bool
	
	// Return a map identical to this one except that the Nth bit is set.
	func (bm *BitMap64) Set(n uint) *BitMap64
	
	// Return a map which is the XOR of the two inputs.
	func (bm *BitMap64) SymmetricDifference(other *BitMap64) *BitMap64
	
	// Test the Nth bit in the map
	func (bm *BitMap64) Test(n uint) bool
	
	// Return a map which is the union of the two inputs -- that is,
	// where all of the bits which are set in either of the two inputs
	// is set in the output.
	func (bm *BitMap64) Union(other *BitMap64) *BitMap64

## DecimalVersion

These functions translate unsigned 32-bit integers to and from the
string serialization used by XLattice for version numbers.

	type DecimalVersion uint32
	
	func New(a, b, c, d uint) (dv DecimalVersion)
	
	// Interpret a byte slice as a big-endian uint.
	func VersionFromBytes(b []byte) (dv DecimalVersion, err error)
	
	// Convert a uint32 DecimalVersion to string format.
	func (dv DecimalVersion) String() (s string)
	
	// Convert a string like a.b.c.d back to a uint32 DecimalVersion.  At
	// least one digit must be present.
	func ParseDecimalVersion(s string) (dv DecimalVersion, err error)

## EntityName

Functions concerned with defining whether a sequence of charactrs
is or is not a valid XLattice name.

	func ValidEntityName(name string) (err error)
	func NAME_PAT() string
	func NAME_RE() *regexp.Regexp
	func INVALID_NAME() error

## LFS

Convenience functions for dealing with the local file system.

	// If the directory named does not exist, create it.  The permisssions
	// passed are ORed with 0700.  If the directory name is empty, call it
	// "lfs", that is, ./lfs/
	//
	// If the directory named exists, permissions are not inspected.
	
	func CheckLFS(lfs string, perm os.FileMode) (err error)
	
	// Given a path to a file, create any missing intermediate directories.
	func MkdirsToFile(pathToFile string, perm os.FileMode) (err error)
## Math

	// Return the smallest non-negative integer exp where 2^exp is
	// greater than or equal to n.
	func NextExp2_32(n uint32) (exp int)
	
	// Return the smallest non-negative integer exp where 2^exp is
	// greater than or equal to n.
	func NextExp2_64(n uint64) (exp int)
	
	// Return the smallest 32-bit number k which is a power of two and greater
	// than // or equal to n.
	func NextPow2_32(n uint32) (k uint32)
	
	// Return the smallest 64-bit number k which is a power of two and greater
	// than or equal to n.
	func NextPow2_64(n uint64) (k uint64)

### MerkleTree

For further information on this package click
[here](https://jddixon.github.io/xlUtil_go/merkletree.html)

### NLHTree

For further information on this package click
[here](https://jddixon.github.io/xlUtil_go/nlhtree.html)

### popCount

This is an implementation in software of `popCount` (also called `CTPOP`)
which on more
modern machine architectures is a machine instruction for counting
the number of bits set in a word.

### Timestamp

	/**
	 * Convenience class handling YYYY-MM-DD HH:MM:SS formatted dates.
	 */
	type Timestamp int64
	
	func (t Timestamp) String() (x string)
	
	func ParseTimestamp(s string) (t Timestamp, err error)

## Project Status

Except as noted above, the code in this project is stable and well-tested.

## On-line Documentation

More information on the **xlUtil_go** project can be found
[here](https://jddixon.github.io/xlUtil_go)
