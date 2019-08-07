package true

import (
	"os"

	"git.nirenjan.com/n/posix/cli"
)

func init() {
	cli.Register("true", func_true)
}

func func_true() {
	os.Exit(0)
}
