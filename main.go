package main

import (
	"nirenjan.org/posix/cli"

	// Commands
	_ "nirenjan.org/posix/cmds/basename"
	_ "nirenjan.org/posix/cmds/dirname"
	_ "nirenjan.org/posix/cmds/echo"
	_ "nirenjan.org/posix/cmds/false"
	_ "nirenjan.org/posix/cmds/logname"
	_ "nirenjan.org/posix/cmds/sleep"
	_ "nirenjan.org/posix/cmds/true"
)

func main() {
	cli.Parse()
}
