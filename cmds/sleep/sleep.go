package sleep

/**********************************************************************
NAME
    sleep - suspend execution for an interval
SYNOPSIS
    sleep time
DESCRIPTION
    The sleep utility shall suspend execution for at least the integral number
    of seconds specified by the time operand.
OPTIONS
    None.
OPERANDS
    The following operand shall be supported:
    time
        A non-negative decimal integer specifying the number of seconds for
        which to suspend execution.
STDIN
    Not used.
INPUT FILES
    None.
ENVIRONMENT VARIABLES
    The following environment variables shall affect the execution of sleep:
    LANG
        Provide a default value for the internationalization variables that are
        unset or null. (See XBD Internationalization Variables for the
        precedence of internationalization variables used to determine the
        values of locale categories.)
    LC_ALL
        If set to a non-empty string value, override the values of all the
        other internationalization variables.
    LC_CTYPE
        Determine the locale for the interpretation of sequences of bytes of
        text data as characters (for example, single-byte as opposed to
        multi-byte characters in arguments).
    LC_MESSAGES
        Determine the locale that should be used to affect the format and
        contents of diagnostic messages written to standard error.
    NLSPATH
        [XSI] Determine the location of message catalogs for the processing of
        LC_MESSAGES.
ASYNCHRONOUS EVENTS
    If the sleep utility receives a SIGALRM signal, one of the following
    actions shall be taken:
    *   Terminate normally with a zero exit status.
    *   Effectively ignore the signal.
    *   Provide the default behavior for signals described in the ASYNCHRONOUS
        EVENTS section of Utility Description Defaults. This could include
        terminating with a non-zero exit status.
    The sleep utility shall take the standard action for all other signals.
STDOUT
    Not used.
STDERR
    The standard error shall be used only for diagnostic messages.
OUTPUT FILES
    None.
EXTENDED DESCRIPTION
    None.
EXIT STATUS
    The following exit values shall be returned:
     0
        The execution was successfully suspended for at least time seconds, or
        a SIGALRM signal was received. See the ASYNCHRONOUS EVENTS section.
    >0
        An error occurred.
CONSEQUENCES OF ERRORS
    Default.
 **********************************************************************
*/

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"nirenjan.org/posix/cli"
)

func init() {
	cli.Register("sleep", sleep)
}

func sleep() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: sleep <seconds>")
		os.Exit(1)
	}

	secs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "sleep: Invalid time value '%v'\n", os.Args[1])
		os.Exit(1)
	}

	time.Sleep(time.Duration(secs) * time.Second)
}
