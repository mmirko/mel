package mel3program

import (
	"errors"
)

const (
	B_IN_DECLARE = uint16(0) + iota
)

func isBuiltin(programName string) bool {
	switch programName {
	case "decl":
		return true
	}
	return false
}

func processBuiltin(programName string, args []string) (*Mel3Program, *ArgumentsTypes, error) {

	switch programName {
	case "decl":

	}

	return nil, nil, errors.New("built-in processing failed")
}
