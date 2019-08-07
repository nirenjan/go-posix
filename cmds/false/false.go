package false

import (
	"os"

	"git.nirenjan.com/n/posix/cli"
)

func init() {
	cli.Register("false", func_false)
}

func func_false() {
	os.Exit(1)
}
