package merkletree

// xlattice_go/util/merkletree/merkletree.go

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	xu "github.com/jddixon/xlUtil_go"
	xf "github.com/jddixon/xlUtil_go/lfs"
	"golang.org/x/crypto/sha3"
	"hash"
	"path"
	re "regexp"
	"strings"
)

var _ = fmt.Print

// The path to a tree, and the SHA hash of the path and the treehash.

type MerkleDoc struct {
	bound   bool
	exRE    *re.Regexp // exclusions
	matchRE *re.Regexp // must be matched
	tree    *MerkleTree

	path     string
	hash     []byte
	whichSHA int
}

// XXX "MUST ADD matchRE and exRE and test on their values at this level."

func NewMerkleDoc(pathToDir string, whichSHA int, binding bool,
	tree *MerkleTree, exRE, matchRE *re.Regexp) (m *MerkleDoc, err error) {

	if pathToDir == "" {
		err = EmptyPath
	}
	if err == nil {
		if strings.HasSuffix(pathToDir, "/") {
			pathToDir = pathToDir[:len(pathToDir)-1]
		}
		self := MerkleDoc{
			exRE:     exRE,
			matchRE:  matchRE,
			path:     pathToDir,
			whichSHA: whichSHA,
		}
		p := &self
		if tree != nil {
			err = p.SetTree(tree)
		} else if !binding {
			err = NilTreeButNotBinding
		}
		if err == nil && binding {
			var whether bool
			fullerPath := path.Join(pathToDir, tree.name)
			whether, err = xf.PathExists(fullerPath)
			if err == nil && !whether {
				err = DirectoryNotFound
			}
		}
		if err == nil {
			m = p
		}
	}
	return
}
func (md *MerkleDoc) GetTree() *MerkleTree {
	return md.tree
}
func (md *MerkleDoc) SetTree(tree *MerkleTree) (err error) {
	if tree == nil {
		err = NilTree
	} else {
		var digest hash.Hash
		switch md.whichSHA {
		case xu.USING_SHA1:
			digest = sha1.New()
		case xu.USING_SHA2:
			digest = sha256.New()
		case xu.USING_SHA3:
			digest = sha3.New256()
			// XXX DEFAULT = ERROR
		}
		digest.Write(tree.hash)
		digest.Write([]byte(md.path))
		md.tree = tree
		md.hash = digest.Sum(nil)
	}
	return
}
func (md *MerkleDoc) Bound() bool {
	return md.bound
}

func (md *MerkleDoc) Equal(any interface{}) bool {
	if any == md {
		return true
	}
	if any == nil {
		return false
	}
	switch v := any.(type) {
	case *MerkleDoc:
		_ = v
	default:
		return false
	}
	other := any.(*MerkleDoc) // type assertion

	return md.path == other.path &&
		bytes.Equal(md.hash, other.hash) &&
		md.tree.Equal(other.tree)
}

func (md *MerkleDoc) GetHash() []byte {
	return md.hash
}

// XXX SetHash missing

func (md *MerkleDoc) GetPath() string {
	return md.path
}
func (md *MerkleDoc) SetPath(value string) (err error) {
	// XXX STUB: MUST CHECK VALUE
	md.path = value
	return
}
func (md *MerkleDoc) GetExRE() *re.Regexp {
	return md.exRE
}
func (md *MerkleDoc) GetMatchRE() *re.Regexp {
	return md.matchRE
}
func (md *MerkleDoc) WhichSHA() int {
	return md.whichSHA
}

// Create a MerkleDoc based on the information in the directory at
// pathToDir.  The name of the directory will be the last component of
// pathToDir.  Return the MerkleTree.

func CreateMerkleDocFromFileSystem(pathToDir string, whichSHA int,
	exclusions, matches []string) (md *MerkleDoc, err error) {

	if len(pathToDir) == 0 {
		err = NilPath
	}
	if err == nil {
		var found bool
		found, err = xf.PathExists(pathToDir)
		if err == nil && !found {
			err = FileNotFound
		}
	}
	// get the path to the directory, excluding the directory name
	var (
		path string
		// dirName		string
		exRE, matchRE *re.Regexp
		tree          *MerkleTree
	)
	if strings.HasSuffix(pathToDir, "/") {
		pathToDir = pathToDir[:len(pathToDir)-1] // drop trailing slash
	}
	parts := strings.Split(pathToDir, "/")
	if len(parts) == 1 {
		path = "."
		// dirName = pathToDir
	} else {
		partCount := len(parts)
		// dirName = parts[partCount - 1]
		parts = parts[:partCount-1]
		path = strings.Join(parts, "/")
	}
	if exclusions != nil {
		exRE, err = MakeExRE(exclusions)
		if err == nil && matches != nil {
			matchRE, err = MakeMatchRE(matches)
		}
	}
	if err == nil {
		tree, err = CreateMerkleTreeFromFileSystem(
			pathToDir, whichSHA, exRE, matchRE)
		if err == nil {
			// "creates the hash"
			md, err = NewMerkleDoc(path, whichSHA, false, tree, exRE, matchRE)
			if err == nil {
				md.bound = true
			}
		}
	}
	return
}

func ParseMerkleDocFirstLine(line string) (
	docHash []byte, docPath string, err error) {

	line = strings.TrimRight(line, " \t")

	groups := FIRST_LINE_RE_1d.FindStringSubmatch(line)
	if groups == nil {
		groups = FIRST_LINE_RE_3d.FindStringSubmatch(line)
	}
	if groups == nil {
		err = CantParseFirstLine
	} else {
		docHash, err = hex.DecodeString(groups[1])
		if err == nil {
			docPath = groups[2] // includes terminating slash
		}
	}
	return
}

func ParseMerkleDoc(s, deltaIndent string) (md *MerkleDoc, err error) {
	if len(s) == 0 {
		err = EmptySerialization
	} else {
		ss := strings.Split(s, "\r\n")
		md, err = ParseMerkleDocFromStrings(&ss, deltaIndent)
	}
	return
}

func ParseMerkleDocFromStrings(ss *[]string, deltaIndent string) (
	md *MerkleDoc, err error) {

	var (
		docHash  []byte
		path     string
		tree     *MerkleTree
		whichSHA int
	)
	if ss == nil {
		err = NilSerialization
	} else {
		docHash, path, err = ParseMerkleDocFirstLine((*ss)[0])
	}
	if err == nil {
		switch len(docHash) {
		case xu.SHA1_BIN_LEN:
			whichSHA = xu.USING_SHA1
		case xu.SHA2_BIN_LEN:
			whichSHA = xu.USING_SHA2
			// XXX NOT A VALID CASE
			//case SHA3_BIN_LEN:
			//	whichSHA = xu.USING_SHA3
			// XXX otherwise internal error
		}
		rest := (*ss)[1:]
		tree, err = ParseMerkleTreeFromStrings(&rest, deltaIndent)
	}
	if err == nil {
		md, err = NewMerkleDoc(path, whichSHA, false, tree, nil, nil)
	}
	return
}

//    # QUASI-CONSTRUCTORS ############################################
//
//    @staticmethod
//    def createFromStringArray(s):
//        """
//        The string array is expected to follow conventional indentation
//        rules, with zero indentation on the first line and some multiple
//        of two spaces on all successive lines.
//        """
//        if s == None:
//            raise RuntimeError('null argument')
//        # XXX check TYPE - must be array of strings
//        if len(s) == 0:
//            raise RuntimeError("empty string array")
//
//        (docHash, docPath) = \
//                            MerkleDoc.parseFirstLine(s[0].rstrip())
//#       print "DEBUG: doc first line: hash = %s, path = %s" % (
//#                               docHash, docPath)
//        whichSHA = (40 == len(docHash))
//
//        tree = MerkleTree.createFromStringArray( s[1:] )
//
//        #def __init__ (self, path, binding = False, tree = None,
//        #    exRE    = None,    # exclusions, which are Regular Expressions
//        #    matchRE = None):   # matches, also Regular Expressions
//        doc = MerkleDoc( docPath, whichSHA, False, tree )
//        doc.hash = docHash
//        return doc
//
//    # CLASS METHODS #################################################
//    @staticmethod
//    def parseFirstLine(line):
//        line = line.rstrip()
//        m = re.match(MerkleDoc.FIRST_LINE_PAT_1, line)
//        if m == None:
//            m = re.match(MerkleDoc.FIRST_LINE_PAT_3, line)
//        if m == None:
//            raise RuntimeError(
//                    "MerkleDoc first line <%s> does not match expected pattern" %  line)
//        docHash  = m.group(1)
//        docPath  = m.group(2)          # includes terminating slash
//        return (docHash, docPath)

// Given a string array of regular expressions, append a list of standard
// exclusions, and then return the compiled regexp.

func MakeExRE(excl []string) (exRE *re.Regexp, err error) {
	excl = append(excl, "^\\.$")       // .
	excl = append(excl, "^\\.\\.$")    // ..
	excl = append(excl, "^\\.merkle$") // merkletree hidden files
	excl = append(excl, "^\\.git$")    // git control data
	excl = append(excl, "^\\.svn$")    // subversion control data
	// some might disagree with these:
	excl = append(excl, "^junk")
	excl = append(excl, "^\\..*\\.swp$") // vi editor files

	exPat := strings.Join(excl, "|")
	exRE, err = re.Compile(exPat)
	return
}

// Given a possibly empty list of expressions to be matched, return
// nil if the list is empty or a regular expression which matches
// any of the patterns in the list.
func MakeMatchRE(matches []string) (matchRE *re.Regexp, err error) {
	if len(matches) > 0 {
		matchPat := strings.Join(matches, "|")
		matchRE, err = re.Compile(matchPat)
	}
	return
}

// SERIALIZATION ====================================================
// XXX WHY THE INDENT?
func (md *MerkleDoc) ToString(indent, deltaIndent string) (
	s string, err error) {

	hexHash := hex.EncodeToString(md.hash)
	topLine := fmt.Sprintf("%s%s %s/\r\n", indent, hexHash, md.path)
	treeText, err := md.tree.ToString("", deltaIndent)
	if err == nil {
		s = topLine + treeText
	}
	return
}
