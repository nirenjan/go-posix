package true

import (
	"os"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("true", func_true)
}

func func_true() {
	os.Exit(0)
}
