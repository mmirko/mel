package mel3program

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/mmirko/mel/pkg/mel"
)

const (
	MAXARITY = 5
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
					arityFactor := 0
					for _, pType := range result.Programs[typeId] {
						if pType.Arity == arity {
							arityFactor++
						}
					}
					for k, pType := range result.Programs[typeId] {
						if pType.Arity == arity {
							result.Programs[typeId][k].ArityFactor = arityFactor
						}
					}

					found = true
					break
				}
			}
			if !found {
				result.Programs[len(result.Types)] = ProgramsTypes{ProgType{LibraryID: i, ProgramID: j, Arity: arity, ArityFactor: 1}}
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

func arityPreferred(in float64) float64 {
	m := float64(MAXARITY)
	l := 2.0
	return m / (1 + math.Exp(in/l)) // in is currnode-targetnode
}

func typeCNorm(in float64) float64 {
	l := 1.0
	return l / (l + in)
}

func (gm *GenerationMatrix) GenerateTree(outType ArgType, inTypes []ArgType, levelNode *int64, targetNode int64, params *mel.EvolutionParameters) (*Mel3Program, error) {

	var selected ProgType
	found := false
	for i, tpy := range gm.Types {
		if SameType(outType, tpy) {
			prob := make([]float64, len(gm.Programs[i]))
			sum := float64(0)
			arityP := arityPreferred(float64(*levelNode - targetNode))
			fmt.Println("Arity preferred: ", arityP, "diff", float64(*levelNode-targetNode))
			for j, prog := range gm.Programs[i] {
				if prog.Arity == -1 {
					// Variadic program
					prob[j] = typeCNorm(math.Abs(float64(MAXARITY)-arityP)) / float64(prog.ArityFactor)
				} else {
					prob[j] = typeCNorm(math.Abs(float64(prog.Arity)-arityP)) / float64(prog.ArityFactor)
				}
				sum += prob[j]
			}

			// Normalize the probabilities
			for j, _ := range prob {
				prob[j] /= sum
			}

			fmt.Println("Probabilities: ", prob)

			// Draw a float
			r := rand.Float64()
			sum = float64(0)
			selected = gm.Programs[i][len(gm.Programs[i])-1]
			for j, prog := range gm.Programs[i] {
				sum += prob[j]
				if r < sum {
					selected = prog
					break
				}
			}

			found = true
			break
		}
	}

	if !found {
		return nil, errors.New("type not found")
	}

	var result *Mel3Program

	switch selected.Arity {
	case 0:
		// Terminal program
		for j, progType := range gm.TerminalProgramTypes {
			if SameProg(progType, selected) {
				gen := gm.TerminalGenerators[j]
				if res, err := gen(params); err != nil {
					return nil, err
				} else {
					result = res
				}
			}
		}
		return result, nil
	case -1:
		// Variadic program, draw an arity
		arity := rand.Intn(MAXARITY) + 1
		result = new(Mel3Program)
		result.LibraryID = selected.LibraryID
		result.ProgramID = selected.ProgramID
		result.NextPrograms = make([]*Mel3Program, 0)
		targetDelta := targetNode / int64(arity)
		target := int64(0)
		arg := gm.Impls[selected.LibraryID].VariadicType[selected.ProgramID]
		for i := 0; i < arity; i++ {
			*levelNode++
			target += targetDelta
			prog, err := gm.GenerateTree(arg, inTypes, levelNode, target, params)
			if err != nil {
				return nil, err
			}
			result.NextPrograms = append(result.NextPrograms, prog)
		}
		return result, nil
	default:
		// Functional program
		result = new(Mel3Program)
		result.LibraryID = selected.LibraryID
		result.ProgramID = selected.ProgramID
		result.NextPrograms = make([]*Mel3Program, 0)
		targetDelta := targetNode / int64(selected.Arity)
		target := int64(0)
		for _, arg := range gm.Impls[selected.LibraryID].NonVariadicArgs[selected.ProgramID] {
			*levelNode++
			target += targetDelta
			prog, err := gm.GenerateTree(arg, inTypes, levelNode, target, params)
			if err != nil {
				return nil, err
			}
			result.NextPrograms = append(result.NextPrograms, prog)
		}
		return result, nil
	}
}
