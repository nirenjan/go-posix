package main

import (
	"git.nirenjan.com/n/posix/cli"

	// Commands
	_ "git.nirenjan.com/n/posix/cmds/basename"
	_ "git.nirenjan.com/n/posix/cmds/dirname"
	_ "git.nirenjan.com/n/posix/cmds/echo"
	_ "git.nirenjan.com/n/posix/cmds/false"
	_ "git.nirenjan.com/n/posix/cmds/logname"
	_ "git.nirenjan.com/n/posix/cmds/sleep"
	_ "git.nirenjan.com/n/posix/cmds/true"
)

func main() {
	cli.Parse()
}
