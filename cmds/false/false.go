package false

import (
	"os"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("false", func_false)
}

func func_false() {
	os.Exit(1)
}
