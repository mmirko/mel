package mel3parser

import (
	"errors"
	"strings"
)

// It returns the functional value of an expression of comma fail if none
func Functval(instr string) (string, error) {

	instr = strings.TrimSpace(instr)

	for i := 0; i < len(instr); i++ {
		if instr[i] == '(' {
			return instr[0:i], nil
		}
	}

	return "", errors.New("Unknown format")
}

// It returns a string slice of the comma separated arguments (it take note of parentesis)
func Parparser(instr string) ([]string, error) {

	// Up to 10 arguments
	var result = make([]string, 0)

	currlevel := uint64(0)
	currargn := uint64(0)
	currarg := ""

	instr = strings.TrimSpace(instr)

	for i := 0; i < len(instr); i++ {
		if instr[i] == '(' {
			if currlevel != 0 {
				currarg = currarg + instr[i:i+1]
			}
			currlevel++
			continue
		}

		if instr[i] == ')' {
			if currlevel == 0 {
				break
			}

			if currlevel == 1 {
				result = append(result, currarg)
				break
			} else {
				currarg = currarg + instr[i:i+1]
				currlevel--
				continue
			}
		}

		if instr[i] == ',' && currlevel == 1 {
			result = append(result, currarg)
			currargn++
			currarg = ""
			continue
		}

		if currlevel >= 1 {
			currarg = currarg + instr[i:i+1]
		}

	}

	return result, nil
}
