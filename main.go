package main

import (
	"github.com/nirenjan/posix/cli"

	// Commands
	_ "github.com/nirenjan/posix/cmds/logname"
)

func main() {
	cli.Parse()
}
