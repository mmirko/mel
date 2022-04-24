//go:build DEBUG
// +build DEBUG

package mel3program

import (
	"errors"
	"fmt"
	mel3parser "github.com/mmirko/mel/mel3parser"
)

// Mel string import: Call the Mel3 engine to get a program from a string
func (prog *Mel3_object) MelStringImport(input_string string) error {

	if prog == nil {
		return errors.New("Uninitializated Program Object.")
	}

	impl := prog.Implementation

	if impl == nil {
		return errors.New("Uninitializated Program Implementation")
	}

	if result, _, err := import_engine(impl, input_string); err == nil {
		prog.StartProgram = result
		return nil
	} else {
		return errors.New("Import error: " + err.Error())
	}

	return errors.New("Error unreachable.")

}

// String importer engine: it parse a string and create a program recursivly
func import_engine(implementation map[uint16]*Mel3_implementation, input_string string) (*Mel3_program, *ArgumentsTypes, error) {
	var result Mel3_program

	// Get the program name
	programname, err := mel3parser.Functval(input_string)
	if err != nil {
		return nil, nil, errors.New("Failed to find identifier on " + input_string)
	}

	args, err := mel3parser.Parparser(input_string)
	if err != nil {
		return nil, nil, errors.New("Failed to find arguments on " + input_string)
	}

	// Funcitonall Programs can share names but cannot share the same name with a non funcional one
	libraryids, programids, ok := ids_from_name(implementation, programname)
	if !ok {
		return nil, nil, errors.New("Failed to find program id of " + programname)
	}

	isfunctional := true

	var nonfunctlib uint16

	// If one is not functional all are
	for i, libraryid := range libraryids {
		programid := programids[i]
		impl := implementation[libraryid]
		if len(impl.NonVariadicArgs[programid]) == 0 && !impl.IsVariadic[programid] {
			isfunctional = false
			nonfunctlib = libraryid
			break
		}
	}

	if isfunctional {

		// Make space for the leaves programs
		result.NextPrograms = make([]*Mel3_program, len(args))

		arglist := ArgumentsTypes{}

		for i := 0; i < len(args); i++ {
			if temp_progr, temp_type, err := import_engine(implementation, args[i]); err != nil {
				return nil, nil, err
			} else {
				result.NextPrograms[i] = temp_progr

				// Composition of the argument list
				for _, itype := range *temp_type {
					arglist = append(arglist, itype)
				}
			}
		}

		// Creation of a signature based on computed argument list
		tempsignature := programname + "("

		for i, arg := range arglist {
			if i != 0 {
				tempsignature += ","
			}
			impl := implementation[arg.LibraryID]
			tempsignature += impl.Implname + "." + impl.TypeNames[arg.TypeID]
		}

		tempsignature += ")()"

		pids := make([]uint16, 0)
		libs := make([]uint16, 0)

		// Check for matching signatures (there has to be 1)
		for libid, impl := range implementation {
			for programid, sig := range impl.Signatures {
				fmt.Println(sig)
				if MatchSignature(tempsignature, sig, 0) {
					pids = append(pids, programid)
					libs = append(libs, libid)
				}
			}
		}

		if len(pids) != 1 {
			fmt.Println(pids)
			return nil, nil, errors.New("Argument of different type than expected")
		}

		// The new real programid chosen by
		pid := pids[0]
		lid := libs[0]

		result.LibraryID = lid
		result.ProgramID = pid

		ptype := implementation[lid].ProgramTypes[pid]

		return &result, &ptype, nil

	} else {
		// Non funcional program cannot have name ambiguity
		programid, ok := id_from_name(implementation[nonfunctlib], programname)
		if !ok {
			return nil, nil, errors.New("Failed to find program id of " + programname)
		}

		result.LibraryID = nonfunctlib
		result.ProgramID = programid

		ptype := implementation[nonfunctlib].ProgramTypes[programid]

		// args[0] is "" even if there are no args
		result.ProgramValue = args[0]

		return &result, &ptype, nil
	}

	return nil, nil, errors.New("Wrong import")
}
