package main

import (
	"git.nirenjan.com/n/posix/cli"

	// Commands
	_ "git.nirenjan.com/n/posix/cmds/basename"
	_ "git.nirenjan.com/n/posix/cmds/dirname"
	_ "git.nirenjan.com/n/posix/cmds/echo"
	_ "git.nirenjan.com/n/posix/cmds/logname"
	_ "git.nirenjan.com/n/posix/cmds/sleep"
)

func main() {
	cli.Parse()
}
