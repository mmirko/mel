//go:build !DEBUG
// +build !DEBUG

package mel3program

import (
	"errors"
	"io/ioutil"

	mel3parser "github.com/mmirko/mel/pkg/mel3parser"
)

func (p *Mel3Object) LoadProgramFromFile(filename string) error {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return p.MelStringImport(string(fileContent))
}

// Mel string import: Call the Mel3 engine to get a program from a string
func (prog *Mel3Object) MelStringImport(input_string string) error {

	if prog == nil {
		return errors.New("uninitialized program object")
	}

	impl := prog.Implementation

	if impl == nil {
		return errors.New("uninitialized program implementation")
	}

	if result, _, err := import_engine(impl, input_string); err == nil {
		prog.StartProgram = result
		return nil
	} else {
		return errors.New("Import error: " + err.Error())
	}

	return errors.New("Error unreachable.")

}

// String importer engine: it parse a string and create a program recursively
func import_engine(implementation map[uint16]*Mel3Implementation, input_string string) (*Mel3Program, *ArgumentsTypes, error) {
	var result Mel3Program

	// Get the program name
	programName, err := mel3parser.FunctionalValue(input_string)
	if err != nil {
		return nil, nil, errors.New("Failed to find identifier on " + input_string)
	}

	args, err := mel3parser.ParParser(input_string)
	if err != nil {
		return nil, nil, errors.New("Failed to find arguments on " + input_string)
	}

	// Check for built-ins
	if isBuiltin(programName) {
		return processBuiltin(implementation, programName, args)
	}

	// Functional Programs can share names but cannot share the same name with a non functional one
	libraryIds, programIds, ok := ids_from_name(implementation, programName)
	if !ok {
		return nil, nil, errors.New("Failed to find program id of " + programName)
	}

	isFunctional := true

	var nonFunctionalLib uint16

	// If one is not functional all are
	for i, libraryId := range libraryIds {
		programId := programIds[i]
		impl := implementation[libraryId]
		if len(impl.NonVariadicArgs[programId]) == 0 && !impl.IsVariadic[programId] {
			isFunctional = false
			nonFunctionalLib = libraryId
			break
		}
	}

	if isFunctional {

		// Make space for the leaves programs
		result.NextPrograms = make([]*Mel3Program, len(args))

		argList := ArgumentsTypes{}

		for i := 0; i < len(args); i++ {
			if tempProgr, tempType, err := import_engine(implementation, args[i]); err != nil {
				return nil, nil, err
			} else {
				result.NextPrograms[i] = tempProgr

				// Composition of the argument list
				for _, itype := range *tempType {
					argList = append(argList, itype)
				}
			}
		}

		// Creation of a signature based on computed argument list
		tempSignature := programName + "("

		for i, arg := range argList {
			if i != 0 {
				tempSignature += ","
			}
			impl := implementation[arg.LibraryID]
			tempSignature += impl.ImplName + "." + impl.TypeNames[arg.TypeID]
		}

		tempSignature += ")()"

		pids := make([]uint16, 0)
		libs := make([]uint16, 0)

		// Check for matching signatures (there has to be 1)
		for libId, impl := range implementation {
			for programId, sig := range impl.Signatures {
				if MatchSignature(tempSignature, sig, 0) {
					pids = append(pids, programId)
					libs = append(libs, libId)
				}
			}
		}

		if len(pids) != 1 {
			return nil, nil, errors.New("argument of different type than expected")
		}

		// The new real programid chosen by
		pid := pids[0]
		lid := libs[0]

		result.LibraryID = lid
		result.ProgramID = pid

		pType := implementation[lid].ProgramTypes[pid]

		return &result, &pType, nil

	} else {
		// Non functional program cannot have name ambiguity
		programId, ok := id_from_name(implementation[nonFunctionalLib], programName)
		if !ok {
			return nil, nil, errors.New("failed to find program id of " + programName)
		}

		result.LibraryID = nonFunctionalLib
		result.ProgramID = programId

		pType := implementation[nonFunctionalLib].ProgramTypes[programId]

		// args[0] is "" even if there are no args
		result.ProgramValue = args[0]

		return &result, &pType, nil
	}
}
