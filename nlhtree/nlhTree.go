package nlhtree

// xlUtil_go/nlhtree/nlhTree.go

import (
	"encoding/hex"
	"errors"
	"fmt"
	// xu "github.com/jddixon/xlUtil_go"
	"io/ioutil"
	"path/filepath"
	// re "regexp"
)

type NLHTree struct {
	nodes	[]*NLHNode
	NLHNode
}

func NewNLHTree(name string, usingSHA1 bool) (nt *NLHTree, err error) {
	nn := NewNLHNode(name, usingSHA1)
	nt = &NLHTree{
		usingSHA1:	usingSHA1,
		NLHNode:	&nn,
	}
	return
}

func (nt *NLHTree) IsLeaf() bool {
	return false
}

func (nt *NLHTree) Nodes() *[]NLHNodes {
	return nt.nodes
}


func (nt *NLHTree) Equal(any interface{}) bool {
	if any == nt {
		return true
	}
	if any == nil {
		return false
	}
	switch v := any.(type) {
	case *NLHTree:
		_ = v
	default: 
		return false
	}
	other := any.(*NLHTree)
	if other.Name() != nt.Name() {
		return false
	}
	if other.UsingSHA1() != nt.UsingSHA1() {
		return false
	}
	myNodes = nt.Nodes()
	otherNodes = other.Nodes()
	if len(myNodes) != len(otherNodes) {
		return false
	}
	for i := 0; i < len(myNodes); i++ {
		if ! myNodes[i].Equal(otherNodes[i]) {
			return false
		}
	}
	return true
}

// Delete nodes whose names match the pattern.  This is a glob, as in 
// UNIX-style file name pattern matching.
func (nt *NLHTree) Delete(pat string) (err error) {

	var remainder []*NLHNode
	for i := 0; i < len(nt.nodes); i++ {
		node = nt.nodes[i]
		found, err =  filepath.Match(pat, node.name)
		if err != nil {
			break
		} 
		if ! found {
			append(remainder, node)
		}
	}
	if err == nil {
		if len(remainder) < len(nt.nodes) {
			nt.nodes = remainder
		}
	}
	return
}

// Return a slice of nodes whose names match the pattern.  This is
// a glob, as in UNIX-style file name pattern matching.  The slice
// is guaranteed to be sorted by node name.
//
func (nt *NLHTree) Find(pat string) (err error) {

	var matches []*NLHNode
	for i := 0; i < len(nt.nodes); i++ {
		node = nt.nodes[i]
		found, err =  filepath.Match(pat, node.name)
		if err != nil {
			break
		} 
		if found {
			append(matches, node)
		}
	}
	return
}

// Insert an NLHNode into the tree's list of nodes, maintaining sort order. 
// It is an err if a node with the same name already exists.
//
func (nt *NLHTree) Insert(node NLHNodeI) (err error) {
	if node.usingSHA1 != nt.usingSHA1 {
		err = errors.New("Cannot insert Node of incompatible SHA type")
	} else {
        lenNodes := len(nt.nodes)
        name := node.name
        done := false
		for i := 0; i < lenNodes; i++ {
            iName := nt.nodes[i].name
            if name < iName {
                // insert before
                if i == 0 {
					// prepend	
                    nt.nodes = append([]*NLHNode{node}, nt.nodes...)
                    done = true
                    break
                } else {
                    before = nt.nodes[0:i]
					append(before, node)
					append(before, nt.nodes[i:])
					nt.nodes = before
                    done = true
                    break
				}
            } else if name == iName {
				msg = fmt.Sprintf(
					"attempt to add two nodes with the same name: '%s'", name)
				err = errors.New(msg)
			}
		}
        if ! done {
            nt.nodes.append(node)
		}
	}
	return
}

// Return a sorted list of node names.  If the node is a tree, its name is 
// preceded by '* ', a an asterisk followed by a space.  Otherwise the 
// node's name is preceded by two spaces.
//
func (nt *NLHTree) List(pat string) (names []string, err error) {

	var el []string
	nodeCount := len(nt.nodes)
	for i := 0 ; i < nodeCount; i++ {
		q := nodes[i]
		found, err =  filepath.Match(pat, q.name)
		if err != nil {
			break
		}
		found, err = filepath.Match(pat, q.name)
		if err == nil && found {
            if q.isLeaf {
                append(el, ("  " + q.name))
			} else {
                append(el, ("* " + q.name))
			}
		}
	} 
    return 
}

func (nt *NLHTree) String() (s string) {
    var ss []String
	nt.ToStrings(ss, 0)
    s = "\n".join(ss) + "\n"
    return
}

func (nt *NLHTree) ToStrings(ss []string, indent int) {
    append(ss, fmt.Sprintf("%s%s", GetSpaces(indent), nt.name))
	nodeCount := len(nt.nodes)
	for i := 0 ; i < nodeCount ; i++ {
		q = nt.nodes[i]
        if q.IsLeaf() {
			append(ss,  node._toString(indent+1))
        } else {
            q.ToStrings(ss, indent+1)
		}
	}
}

// Create an NLHTree based on the information in the directory at pathToDir.  
// The name of the directory will be the last component of pathToDir.  Return 
// the NLHTree and possibley an error.
//
// XXX Type of exRe and matchRE is wrong.
//
func createNLHTreeFromFileSystem(pathToDir string, usingSHA1 bool,
	exRE, matchRE string) (nt *NLHTree, err error) {

	var dirName string

    if pathToDir = "" {
        err = errors.New("cannot create a NLHTree, no path set")
	} else {
		_, err = os.Stat(pathToDir) 
	} 
	if err = nil {
        parts := pathToDir.split("/")
		if len(parts) < 2 {
			dirName = pathToDir
		} else {
			dirName = parts[-1]
		}

        nt = NewNLHTree(dirName, usingSHA1)

        // Create data structures for constituent files and subdirectories
        // These are sorted by the bare name
        files, err = ioutil.ReadDir(pathToDir)  // sorted
		if err == nil {
			fileCount = len(files)		// FileInfo objects
			for i := 0; i < fileCount; i++ {
                // exclusions take priority over matches
				info := files[i]
				name := info.Name()
                if exRE != "" {
					found, err = filepath.Match(exRE, name) 
					if err != nil {
						break
					} else if found {
						continue
					}
				}
                if matchRE != "" {
					found, err = filepath.Match(matchRE, name) 
					if err != nil {
						break
					} else if ! found {
						continue
					}
				}
                var node NLHNodeI
                pathToFile = os.path.join(pathToDir, name)

				if info.IsDir() {
                    node, err = CreateNLHTreeFromFileSystem(
                            pathToFile, usingSHA1,exRE, matchRE)
					if err != nil {
						break
					}

				} else if info.IsRegular() {
                    node, err = CreateNLHLeafFromFileSystem(
                            pathToFile, file, usingSHA1)
					if err != nil {
						break
					}

				}
				// otherwise we ignore the file

                if node != nil {
                    append(nt.nodes, node)
				}
			}
		}
	}
    return 
}
	
// Return the name found in the first line or return an error.
//
func ParseNLHTreeFirstLine(s string) (name string, err error) {

	m = DIR_LINE_RE.FindStringSubmatch(s)
    if m == nil {
		err = errors.New("first line doesn't match expected pattern")
	} else {
		if len(m[1]) != 0 {		// group 1, spaces
            err = errors.New("unexpected indent on first line")
		}		
		if err = nil {
			name = m[2]			// group 2
		}
	}
	return
}

// Return the indent (the number of spaces), the name on the line,
// and other "" or the hash found.
//
func ParseNLHTreeOtherLine(s string) (
	indent int, name string, hexHash string, err error) {

    m = DIR_LINE_RE.FindStringSubmatch(s)
	if m != nil {
        return len(m[1]), m[2], "", nil
	}
    m = FILE_LINE_RE_1.FindStringSubmatch(s)
    if m != nil {
		return len(m[1]), m[2], m[3], nil
    }    
	m = FILE_LINE_RE_2.FindStringSubmatch(s)
    if m != nil {
            return len(m[1]), m[2], m[3]
	}
	msg = fmt.Sprintf("can't parse line: '%s'", s)
    err = errors.New(msg)
	return
}

// XXX WORKING HERE ///


// At entry, we don't know whether the string array uses SHA1 or SHA256
//
func createNLHTreeFromStringArray(ss []string, usingSHA1 bool) (
	root []NLHTree, err error) {

	var curLevel *NLHtree
	var stack	[]*NLHTree
	var	depth	int

    if len(ss) == 0 {
		err = errors.New("empty string array")
	}
	if err == nil {
        name, err = ParseNLHTreeFirstLine(ss[0])
	}
	if err == nil {
		var		indent int
		var		name, hexHash	string
		var		binHash []byte
        root    = NLHTree(name, usingSHA1)  
		curLevel = root
        stack   = append(stack, root)			// our first push
        depth   = 0
    
		for i := 1; i < len(ss); i++ {
            indent, name, hexHash, err = NLHTree.parseOtherLine(line)
            if hexHash != "" {
                binHash, err = hex.DecodeString(hexHash)
				if err != nil{
					break
				}
			} 
            if indent > depth+1 {
                // DEBUG
                fmt.Printf("IMPOSSIBLE: indent %d, depth %d\n",  indent, depth)
                // END
                if hexHash != "" {
                    leaf := NewNLHLeaf(name, binHash)
                    stack[depth].Insert(leaf)
				} else {
                    subtree = NLHTree(name, usingSHA1)
                    stack.append(subtree)
                    depth += 1
				}
			} else if indent == depth+1 {
                if hexHash == "" {
                    subtree = NLHTree(name, usingSHA1)
                    stack[depth].insert(subtree)
                    stack.append(subtree)		// another push
                    depth += 1
                } else {
                    leaf = NLHLeaf(name, binHash)
                    stack[depth].insert(leaf)
				}
            } else {
                for indent < depth+1 {
                    stack.pop()					// XXX UNIMPLEMENTED
                    depth -= 1
				}
                if hexHash == "" {
                    subtree = NLHTree(name, usingSHA1)
                    stack[depth].insert(subtree)
                    append(stack, subtree)		// push
                    depth += 1
                } else {
                    leaf = NLHLeaf(name, binHash)
                    stack[depth].insert(leaf)
				}
			}
		}
	}
    return
}

func ParseNLHTree(s string, usingSHA1 bool)(nt *NLHTree, err error) {

    if s == "" {
        err = errors.New("cannot parse an empty string")
	} else {	
		ss := s.split("\n")
		if ss[-1] == "" {
			ss = ss[:-1]
		}
		nt, err = CreateNLHTreeFromStringArray(ss, usingSHA1)
	}
	return
}
