package m3bool

import (
	"math/rand"

	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

func M3BoolConstGenerator(ev *mel.EvolutionParameters) (*mel3program.Mel3Program, error) {
	// Teporary code: random number generation
	op_result := "true"
	if rand.Intn(2) == 0 {
		op_result = "false"
	}

	result := new(mel3program.Mel3Program)
	result.LibraryID = MYLIBID
	result.ProgramID = BOOL
	result.ProgramValue = op_result
	result.NextPrograms = nil

	return result, nil
}

func M3BoolVarGenerator(ev *mel.EvolutionParameters) (*mel3program.Mel3Program, error) {
	result := new(mel3program.Mel3Program)
	result.LibraryID = MYLIBID
	result.ProgramID = BOOL
	result.ProgramValue = "x"
	result.NextPrograms = nil

	return result, nil
}
