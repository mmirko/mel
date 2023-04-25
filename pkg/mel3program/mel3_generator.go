package mel3program

import (
	"errors"

	"github.com/mmirko/mel/pkg/mel"
)

type GenerationMatrix struct {
	Impls                map[uint16]*Mel3Implementation
	Types                []ArgType                 // The list of types
	TerminalProgramTypes []ProgType                // The list of programs that can be used as terminal
	Programs             map[int]ProgramsTypes     // The list of programs that can output a given type (the key is the Types index))
	TerminalGenerators   map[int]TerminalGenerator // The list of terminal generators for a given terminal program (the key is the TerminalProgramTypes index)
}

type TerminalGenerator func(*mel.EvolutionParameters) (*Mel3Program, error)

func CreateGenerationMatrix(implementation map[uint16]*Mel3Implementation) *GenerationMatrix {
	result := new(GenerationMatrix)
	result.Impls = implementation
	result.Types = make([]ArgType, 0)
	result.TerminalProgramTypes = make([]ProgType, 0)
	result.Programs = make(map[int]ProgramsTypes)
	result.TerminalGenerators = make(map[int]TerminalGenerator)

	for i, impl := range implementation {
		for j, progType := range impl.ProgramTypes {

			var arity int
			if impl.IsVariadic[j] {
				arity = -1
			} else {
				arity = len(impl.NonVariadicArgs[j])
			}

			found := false
			for typeId, arg := range result.Types {
				if SameType(progType[0], arg) {
					result.Programs[typeId] = append(result.Programs[typeId], ProgType{LibraryID: i, ProgramID: j, Arity: arity})
					found = true
					break
				}
			}
			if !found {
				result.Programs[len(result.Types)] = ProgramsTypes{ProgType{LibraryID: i, ProgramID: j, Arity: arity}}
				result.Types = append(result.Types, progType[0])
			}
		}
	}

	return result
}

func (gm *GenerationMatrix) Init() error {
	for i, typ := range gm.Types {
		found := false
		// Check whether each type has at least one program that has 0 arity (a terminal program)
		for _, prog := range gm.Programs[i] {
			if prog.Arity == 0 {
				// fmt.Println("Terminal program found for type " + typ.String(gm.Impls[typ.LibraryID]))
				found = true

				// Check also whether the terminal program has a terminal generator
				foundTerm := false
				for j, progType := range gm.TerminalProgramTypes {
					if SameProg(progType, prog) {
						if _, ok := gm.TerminalGenerators[j]; !ok {
							return errors.New("Type " + typ.String(gm.Impls[typ.LibraryID]) + " has no terminal generator")
						}
						foundTerm = true
						break
					}
				}

				if !foundTerm {
					return errors.New("Type " + typ.String(gm.Impls[typ.LibraryID]) + " has no terminal program")
				}

				break
			}
		}
		if !found {
			return errors.New("Type " + typ.String(gm.Impls[typ.LibraryID]) + " has no terminal program")
		}

	}
	return nil
}

// // Mel dump, it prints out the object
// func (prog *Mel3Object) String() string {

// 	result := ""

// 	if prog != nil {
// 		impl := prog.Implementation
// 		if impl != nil {
// 			startprog := prog.StartProgram
// 			if startprog != nil {
// 				if tmpResult, err := export_engine(impl, startprog); err == nil {
// 					result = tmpResult
// 				} else {
// 					return result
// 				}
// 			} else {
// 				return result
// 			}
// 		} else {
// 			return result
// 		}
// 	} else {
// 		return result
// 	}

// 	return result
// }

// // Export engine: it recurse over the program and show it
// func export_engine(implementation map[uint16]*Mel3Implementation, program *Mel3Program) (string, error) {

// 	result := ""

// 	if program == nil {
// 		return result, errors.New("Empty program failed to export")
// 	} else {

// 		libraryID := program.LibraryID
// 		programID := program.ProgramID

// 		impl := implementation[libraryID]

// 		isFunctional := true

// 		if len(impl.NonVariadicArgs[programID]) == 0 && !impl.IsVariadic[programID] {
// 			isFunctional = false
// 		}

// 		if isFunctional {
// 			result = result + impl.ProgramNames[programID] + "("
// 			for i := range program.NextPrograms {
// 				if tmpResult, err := export_engine(implementation, program.NextPrograms[i]); err == nil {
// 					result = result + tmpResult
// 					if i != len(program.NextPrograms)-1 {
// 						result = result + ","
// 					}
// 				} else {
// 					return "", err
// 				}
// 			}
// 			result = result + ")"
// 		} else {
// 			result = result + impl.ProgramNames[programID] + "(" + program.ProgramValue + ")"
// 		}
// 	}

// 	return result, nil
// }
