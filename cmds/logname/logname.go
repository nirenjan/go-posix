// Package logname returns the user's login name
package logname

/**********************************************************************
NAME
    logname - return the user's login name
SYNOPSIS
    logname
DESCRIPTION
    The logname utility shall write the user's login name to standard output. The login name shall be the string that would be returned by the getlogin() function defined in the System Interfaces volume of POSIX.1-2017. Under the conditions where the getlogin() function would fail, the logname utility shall write a diagnostic message to standard error and exit with a non-zero exit status.
OPTIONS
    None.
OPERANDS
    None.
STDIN
    Not used.
INPUT FILES
    None.
ENVIRONMENT VARIABLES
    The following environment variables shall affect the execution of logname:
    LANG
        Provide a default value for the internationalization variables that are unset or null. (See XBD Internationalization Variables for the precedence of internationalization variables used to determine the values of locale categories.)
    LC_ALL
        If set to a non-empty string value, override the values of all the other internationalization variables.
    LC_CTYPE
        Determine the locale for the interpretation of sequences of bytes of text data as characters (for example, single-byte as opposed to multi-byte characters in arguments).
    LC_MESSAGES
        Determine the locale that should be used to affect the format and contents of diagnostic messages written to standard error.
    NLSPATH
        [XSI] Determine the location of message catalogs for the processing of LC_MESSAGES.
ASYNCHRONOUS EVENTS
    Default.
STDOUT
    The logname utility output shall be a single line consisting of the user's login name:
    "%s\n", <login name>
STDERR
    The standard error shall be used only for diagnostic messages.
OUTPUT FILES
    None.
EXTENDED DESCRIPTION
    None.
EXIT STATUS
    The following exit values shall be returned:
     0
        Successful completion.
    >0
        An error occurred.
CONSEQUENCES OF ERRORS
    Default.
 **********************************************************************
*/

import (
	"fmt"
	"os"
	"os/user"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("logname", logname)
}

func logname() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: logname")
		os.Exit(1)
	}

	user, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(user.Username)
}
