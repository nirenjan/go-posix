// Package basename returns the basename of the path
package basename

import (
	"fmt"
	"os"
	"path"
	"strings"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("basename", basename)
}

func basename() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Usage: basename string [suffix]")
		os.Exit(1)
	}

	base := path.Base(os.Args[1])

	// Remove suffix if present
	if len(os.Args) == 3 {
		suffix := os.Args[2]
		if strings.HasSuffix(base, suffix) {
			length := len(base) - len(suffix)
			base = base[:length]
		}
	}

	fmt.Println(base)
}
