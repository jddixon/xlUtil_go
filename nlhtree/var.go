package nlhtree

import (
	re "regexp"
)

var (
	// notice the terminating forward slash and lack of newlines or CR-LF
	DIR_LINE_RE = re.MustCompile(
		"^( *)([A-Za-z0-9_$+-.~]+/?)$")
	FILE_LINE_RE_1 = re.MustCompile(
		"^( *)([A-Za-z0-9_$+-.:~]+/?) ([0-9A-Za-f]{40})$")
	FILE_LINE_RE_2 = re.MustCompile(
		"^( *)([A-Za-z0-9_$+-.:~]+/?) ([0-9A-Za-f]{64})$")
)
