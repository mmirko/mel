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

func (gm *GenerationMatrix) AddTerminalGenerator(progType ProgType, generator TerminalGenerator) {
	for i, typ := range gm.TerminalProgramTypes {
		if SameProg(typ, progType) {
			gm.TerminalGenerators[i] = generator
			return
		}
	}
	gm.TerminalProgramTypes = append(gm.TerminalProgramTypes, progType)
	gm.TerminalGenerators[len(gm.TerminalProgramTypes)-1] = generator
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
			}
		}
		if !found {
			return errors.New("Type " + typ.String(gm.Impls[typ.LibraryID]) + " has no 0 arity program")
		}

	}
	return nil
}

func (gm *GenerationMatrix) GenerateTree(outType ArgType, inTypes []ArgType, levelNode *int64, targetNode int64, params *mel.EvolutionParameters) (*Mel3Program, error) {
	// Compute a desired arity

	// Compute a randomized arity

	// Match with available programs

	// Generate the program

	// Generate the arguments recursively

	// Return the program

	return nil, nil
}
