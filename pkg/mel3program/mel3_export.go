package mel3program

import (
	"errors"
)

// Mel dump, it prints out the object
func (prog *Mel3Object) String() string {

	result := ""

	if prog != nil {
		impl := prog.Implementation
		if impl != nil {
			startprog := prog.StartProgram
			if startprog != nil {
				if tmpResult, err := exportEngine(impl, startprog); err == nil {
					result = tmpResult
				} else {
					return result
				}
			} else {
				return result
			}
		} else {
			return result
		}
	} else {
		return result
	}

	return result
}

func ProgDump(implementation map[uint16]*Mel3Implementation, program *Mel3Program) (string, error) {
	return exportEngine(implementation, program)
}

// Export engine: it recurse over the program and show it
func exportEngine(implementation map[uint16]*Mel3Implementation, program *Mel3Program) (string, error) {

	result := ""

	if program == nil {
		return result, errors.New("empty program failed to export")
	} else {

		libraryID := program.LibraryID
		programID := program.ProgramID

		impl := implementation[libraryID]

		isFunctional := true

		if len(impl.NonVariadicArgs[programID]) == 0 && !impl.IsVariadic[programID] {
			isFunctional = false
		}

		if isFunctional {
			result = result + impl.ProgramNames[programID] + "("
			for i := range program.NextPrograms {
				if tmpResult, err := exportEngine(implementation, program.NextPrograms[i]); err == nil {
					result = result + tmpResult
					if i != len(program.NextPrograms)-1 {
						result = result + ","
					}
				} else {
					return "", err
				}
			}
			result = result + ")"
		} else {
			result = result + impl.ProgramNames[programID] + "(" + program.ProgramValue + ")"
		}
	}

	return result, nil
}
