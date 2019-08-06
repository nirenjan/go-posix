package main

import (
	"git.nirenjan.com/n/posix/cli"

	// Commands
	_ "git.nirenjan.com/n/posix/cmds/logname"
	_ "git.nirenjan.com/n/posix/cmds/sleep"
)

func main() {
	cli.Parse()
}
