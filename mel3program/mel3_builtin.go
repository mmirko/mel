package mel3program

import (
	"errors"
)

const (
	B_IN_DECLARE = uint16(0) + iota
)

func isBuiltin(programname string) bool {
	switch programname {
	case "decl":
		return true
	}
	return false
}

func processBuiltin(programname string, args []string) (*Mel3_program, *ArgumentsTypes, error) {

	switch programname {
	case "decl":

	}

	return nil, nil, errors.New("Builtin processing failed")
}
