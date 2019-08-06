package main

import (
	"github.com/nirenjan/posix/cli"

	// Commands
	_ "github.com/nirenjan/posix/cmds/logname"
	_ "github.com/nirenjan/posix/cmds/sleep"
)

func main() {
	cli.Parse()
}
