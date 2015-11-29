# nlhtree

## Tree

An **NLHTree** is a data structure in which each node consists of a pair:
a name and either a list of such pairs or a hash.  The NLHTree is serialized
as an indented list, with each immediate child indented one space more than
the parent

## Hash

All of the hashes
in the tree have the same number of bits, either 160 (for an SHA1 hash)
or 256 (for an SHA2 or SHA3 hash).

## Node Names

All names in the tree must be valid file names.  For the moment, this
will be understood to include letters, both upper and lower case,
digits, the dash ('-'), and the underscore ('_').  Node names may not
include either spaces or line breaks (CR=13 and LF=10).

## Top

The topmost node in the tree is a pair of the first type and belongs to the
**NLHTree** class.  It consists of a name and a list of **NLHNodes**, where an
NLHNode is either an NLHTree or an **NLFLeaf**.  At each level the list of
nodes is sorted by name.

## Intermediate Nodes

All intermediate nodes in the tree are also instances of the NLHTree
class and so consist of a name and a (possibly empty) sorted list.

## Leaf Nodes

Leaf nodes in the tree, instances of the **NLHLeaf** class, consist of a valid
name and a hash.  Once formed an NLHLeaf is immutable in the sense that
its fields (its name and its hash) cannot be changed.  If the leaf node is
in an NLHTree, the tree has a reference to
it.

## Example

	dataDir
	 data1 bea7383743859a81b84cec8fde2ccd1f3e2ff688
	 data2 895c210f5203c48c1e3a574a2d5eba043c0ec72d
	 subDir1
	  data11 cb0ece05cbb91501d3dd78afaf362e63816f6757
	  data12 da39a3ee5e6b4b0d3255bfef95601890afd80709
	 subDir2
	 subDir3
	  data31 8cddeb23f9de9da4547a0d5adcecc7e26cb098c0
	 subDir4
	  subDir41
	   subDir411
	    data41 31c16def9fc4a4b6415b0b133e156a919cf41cc8
	 zData 31c16def9fc4a4b6415b0b133e156a919cf41cc8

In this example, the NLHTree represents the files in the directory `dataDir`.
The directory contains three files (`data1`, `data2`, and `zData` and four
subdirectories (`subDir1`, `subDir2` (which is empty), `subDir3`, and
`subDir4`.  The name of each of the leaf nodes (files) is followed by its
content hash.  In this case these are SHA1 hashes, which are 20 bytes long,
and so written as 40 hex digits.

## Utility

NLHTrees are useful as concise descriptions of file systems.  In particular
they are used in building and editing
[BuildLists.](https://jddixon.github.io/buildList)

A BuildList contains a recursive data structure, its NLHTree.  Each leaf
node has associated with it a content hash which can be used to verify the
integrity of the file.  The BuildList has an associated RSA public key,
title and date and can be signed using the signatory's private RSA key.
The presence of the RSA public key makes it possible for anyone obtaining
a copy of the BuildList to verify that the list is a true representation
of the BuildList.

## Package Status

A good beta.  All tests succeed.

