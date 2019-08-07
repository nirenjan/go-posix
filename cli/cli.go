// Package cli provides the framework for adding commands to the base image
package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

var commands = make(map[string]func())

const progName string = "posix"

func Register(cmd string, handler func()) error {
	// Make sure there are no duplicates
	_, present := commands[cmd]
	if present {
		err := fmt.Errorf("Duplicate registration of '%v'", cmd)
		return err
	}

	// Register the command
	commands[cmd] = handler
	return nil
}

func Parse() {
	var offset int
	var cmd string
	cmd = filepath.Base(os.Args[0])
	if cmd == progName {
		if len(os.Args) > 1 {
			cmd = os.Args[1]
			offset = 1
		} else {
			fmt.Fprintf(os.Stderr, "usage: %v <command> [args]\n", progName)
			os.Exit(2)
		}
	}

	handler, registered := commands[cmd]
	if !registered {
		fmt.Fprintf(os.Stderr, "%v: Unknown command '%v'\n", progName, cmd)
		os.Exit(2)
	}

	// Update the Args variable so that the individual handlers
	// don't have to deal with offsets
	if offset > 0 {
		os.Args = os.Args[offset:]
	}
	handler()
}
