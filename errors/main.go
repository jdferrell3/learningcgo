package main

// Random code for experimenting with error handling

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

var ErrNoArgs = fmt.Errorf("zero args")

type argsError struct {
	ArgCount int
}

func (e *argsError) Error() string {
	return fmt.Sprintf("Only %d args, neeed 4", e.ArgCount-1)
}

func processArgs() error {
	if 1 == len(os.Args) {
		return ErrNoArgs
	}
	if len(os.Args) < 5 {
		return &argsError{ArgCount: len(os.Args)}
	}
	return nil
}

func main() {
	err := processArgs()
	if nil != err {
		err = errors.Wrap(err, "wrapped error")
	}

	if nil != err {
		switch err := err.(type) {
		case *argsError:
			fmt.Printf("need more args: %s\n", err.Error())
		default:
			fmt.Printf("Please specify args: %s\n", err.Error())
		}
	}

	switch err := errors.Cause(err).(type) {
	case *argsError:
		fmt.Printf("Cause: %s", err)
	default:
		fmt.Printf("Unknown cause: %s", err)
	}
}
