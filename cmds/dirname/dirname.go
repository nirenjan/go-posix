// Package dirname returns the dirname of the path
package dirname

import (
	"fmt"
	"os"
	"path"

	"git.nirenjan.com/n/posix/cli"
)

func init() {
	cli.Register("dirname", dirname)
}

func dirname() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: dirname string")
		os.Exit(1)
	}

	fmt.Println(path.Dir(os.Args[1]))
}
