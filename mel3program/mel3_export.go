package mel3program

import (
	"errors"
)

// Mel dump, it prints out the object
func (prog *Mel3_object) String() string {

	result := ""

	if prog != nil {
		impl := prog.Implementation
		if impl != nil {
			startprog := prog.StartProgram
			if startprog != nil {
				if tmpresult, err := export_engine(impl, startprog); err == nil {
					result = tmpresult
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

func ProgDump(implementation map[uint16]*Mel3Implementation, program *Mel3_program) (string, error) {
	return export_engine(implementation, program)
}

// Export engine: it recurse over the program and show it
func export_engine(implementation map[uint16]*Mel3Implementation, program *Mel3_program) (string, error) {

	result := ""

	if program == nil {
		return result, errors.New("Empty program failed to export")
	} else {

		libraryid := program.LibraryID
		programid := program.ProgramID

		impl := implementation[libraryid]

		isfunctional := true

		if len(impl.NonVariadicArgs[programid]) == 0 && !impl.IsVariadic[programid] {
			isfunctional = false
		}

		if isfunctional {
			result = result + impl.ProgramNames[programid] + "("
			for i := range program.NextPrograms {
				if tmpresult, err := export_engine(implementation, program.NextPrograms[i]); err == nil {
					result = result + tmpresult
					if i != len(program.NextPrograms)-1 {
						result = result + ","
					}
				} else {
					return "", err
				}
			}
			result = result + ")"
		} else {
			result = result + impl.ProgramNames[programid] + "(" + program.ProgramValue + ")"
		}
	}

	return result, nil
}
