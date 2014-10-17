package lfs

// xlUtil_go/fs/lfs.go

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// If the directory named does not exist, create it.  The permisssions
// passed are ORed with 0700.  If the directory name is empty, call it
// "lfs", that is, ./lfs/
//
// XXX If the directory named exists, permissions are no inspected.

func CheckLFS(lfs string, perm os.FileMode) (err error) {
	perm |= 0700
	if lfs == "" {
		lfs = "lfs"
	}
	fileInfo, err := os.Stat(lfs)
	if os.IsNotExist(err) {
		err = os.MkdirAll(lfs, perm)
	} else if err == nil {
		if !fileInfo.IsDir() {
			errMsg := fmt.Sprintf("%s is not a directory", lfs)
			err = errors.New(errMsg)
		}
	}
	return
}

// Given a path to a file, create any missing intermediate directories.
func MkdirsToFile(pathToFile string, perm os.FileMode) (err error) {

	parts := strings.Split(pathToFile, "/")
	if len(parts) > 1 {
		var pathToDir string
		if len(parts) == 2 {
			// just drop the file name
			pathToDir = parts[0]
		} else {
			pathToDir = strings.Join(parts[:len(parts)-1], "/")
		}
		err = os.MkdirAll(pathToDir, perm)
	}
	return
}
