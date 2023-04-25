package mel3program

type GenerationMatrix struct {
	Types    []ArgType
	Programs map[int]ProgramsTypes
}

func CreateGenerationMatrix(implementation map[uint16]*Mel3Implementation) *GenerationMatrix {
	result := new(GenerationMatrix)
	result.Types = make([]ArgType, 0)
	result.Programs = make(map[int]ProgramsTypes)

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
