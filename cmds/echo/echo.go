// Package echo implements the POSIX `echo` command
package echo

import (
	"fmt"
	"os"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("echo", echo)
}

func echo() {
	for i, v := range os.Args {
		if i == 0 {
			continue
		}

		if printArg(v) != 0 {
			os.Exit(0)
		}

		// Print a space before all but the last argument
		if i != len(os.Args)-1 {
			fmt.Print(" ")
		}
	}

	// Print a final newline
	fmt.Println("")
}

// printArg prints a single argument, and decides if we need to exit early
func printArg(arg string) int {
	var outString string
	var exit int
	var skip int

ARGLOOP:
	for i, c := range arg {
		// If we need to skip this character, then
		// decrement the skip count and continue to the next character
		if skip > 0 {
			skip--
			continue
		}
		s := string(c)
		if i < len(arg)-1 && s == `\` {
			// Check if the next character is a known control character
			s = string(arg[i+1])
			switch s {
			case `a`:
				skip = 1
				outString += "\a"
			case `b`:
				skip = 1
				outString += "\b"
			case `c`:
				exit = 1
				break ARGLOOP
			case `f`:
				skip = 1
				outString += "\f"
			case `n`:
				skip = 1
				outString += "\n"
			case `r`:
				skip = 1
				outString += "\r"
			case `t`:
				skip = 1
				outString += "\t"
			case `v`:
				skip = 1
				outString += "\v"
			case `0`:
				skip, s = parseOctal(arg[i+1:])
				outString += s
			case `\`:
				skip = 1
				fallthrough
			default:
				// If we don't recognize the control character, then
				// simply output the \ as is, and don't skip anything.
				outString += `\`
			}
		} else {
			outString += s
		}
	}

	fmt.Print(outString)
	return exit
}

// parseOctal converts a string of the form \0[0-7]{0,3} into the
// corresponding character
func parseOctal(s string) (int, string) {
	var octal int
	var skip int

	var max int = len(s)
	if max > 4 {
		max = 4
	}

	for i := 0; i < max; i++ {
		c := s[i]
		if c < '0' || c > '7' {
			break
		}
		octal *= 8
		octal += int(c - '0')
		skip++
	}

	return skip, string(octal)
}
