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
[Secure Hash Standard](http://csrc.nist.gov/publications/fips/fips180-4/fips-180-4.pdf)
for hashing.  This is a standard published by the US National Institute of
Standards and Techology (**NIST**).

SHA is a cryptographically secure hash, meaning that for all
practical purposes it is impossible to find two documents with the same hash.
In other words, the SHA hashes are meant to be **one-way**: given a document,
it is very cheap (it requires very little computation)
to determine its SHA hash, but given such a hash the only
practical way to find out what document it corresponds to is to hash all
candidate matches and compare the resultant hash with the one you are searching
for.  Computationally this is extremely expensive or impossible.

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

    Usage: gMerkleize [OPTIONS]
    where the options are:
      -1	use SHA1 hash in building tree
      -P value
        	list of patterns, file name patterns to be matched
      -T	test run
      -V	output package version info
      -X value
        	list of patterns, file names patterns to be excluded
      -i string
        	path to directory being scanned
      -j	display option settings and exit
      -m	output the merkletree
      -o string
        	write serialzed merkletree here
      -t	output UTC timestamp
      -v	be talkative
      -x	output top level hash
	
The default output file name is the UTC timestamp.

**WARNING:** the command line description above may not be current: type

    gMerkleize -h

to confirm the syntax for your release.

## Relationships

Merkletree was implemented as part of the
[XLattice](http://www.xlattice.org)
project.  A Python implementation is available; see
[merkletree](https://jddixon.github.io/xlUtil_go/merkletree).

## Project Status

The library code is well-tested and reliable.  The command line is being
brought into conformance with the Java and Python versions and may change
shortly.

## On-line Documentation
More information on the **merkletree** project can be found
[here](https://jddixon.github.io/merkletree)
