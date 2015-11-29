# merkletree


## What It Does

**merkletree** is a Go package for creating a
[Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree)
for a
directory structure.  A **Merkle tree** is a representation of the contents
of the directory and its subdirectories in terms of hashes.

A file is represented by the hash of its
contents.  A directory is represented by the hash of the hashes
of its members, sorted.  This makes it very easy to verify the
contents of a directory:

	gMerkleize -x -i  .

outputs a single hash, a hexadecimal number.  If any file in the
directory structure has been changed, the output from the above
command will also change.

## SHA, the Secure Hash Algorithm

This package uses hash algorithms specified in the
[Secure Hash Standard](http://nvfpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf)
for hashing.  This is a standard published by the US National Institute of
Standards and Techology (**NIST**).

SHA is a cryptographically secure hash, meaning that for all
practical purposes it is impossible to find two documents with the same hash.
In other words, the SHA hashes are meant to be **one-way**: given a document,
it is very easy to determine its SHA hash, but given such a hash the only
practical way to find out what document it corresponds to is to hash all
candidate matches and compare the resultant hash with the one you are searching
for.

**merkletree** currently uses

* either the older 160 bit/20 byte **SHA1** (aka SHA-1)
* or the more recent and supposedly more secure **SHA2** (aka SHA-256),
  a 256 bit/32 byte hash
* or **SHA3**, the 256-bit/32-byte version of
  [Keccak](https://en.wikipedia.org/wiki/SHA-3), the winner of the 2012
  NIST competition in search of a more secure version of SHA

## What It's Used For

Verifying the integrity of file systems, of directory structures.

## Command Line

	usage: gMerkleize [options]
	where the options are
	  -h, --help           to see this very useful message
	  -1, --usingSHA1      use SHA-1 hash instead of default SHA-256
	  -d  --outDir DIR     write serialized merkletree here
	  -i, --inDir DIR      where DIR names directory being scanned
	  -j, --justShow       list options and exit
	  -m  --showTree       output the merkletree hash/filename pairs
	  -o, --outFile NAME   write output to this file name
	  -P, --match PAT      include ONLY files with matching names
	  -t, --showTimestamp  output UTC timestamp to command line
	  -v, --verbose        verbose: whether the program is chatty
	  -V, --version        show version information
	  -x, --hashOutput     whether to output the top level hash
	  -X, --exclude PAT    don't include files with matching names
	
The default output file name is the UTC timestamp.

**WARNING:** the command line description above may not be current: type

    gMerkleize -h

to confirm the syntax for your release.

## Relationships

Merkletree was implemented as part of the
[XLattice](http://www.xlattice.org)
project.  A Python implementation is available; see
[merkletree](https://gibhub.com/jddixon/merkletree).

## Project Status

The library code is well-tested and reliable.  The command line is being
brought into conformance with the Java and Python versions and may change
shortly.

## On-line Documentation
More information on the **merkletree** project can be found
[here](https://jddixon.github.io/merkletree)