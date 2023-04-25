package m3uint

import (
	"fmt"
	"math/rand"

	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

func M3UintConstGenerator(ev *mel.EvolutionParameters) (*mel3program.Mel3Program, error) {
	// Teporary code: random number generation
	op_result := fmt.Sprint(rand.Intn(10))

	result := new(mel3program.Mel3Program)
	result.LibraryID = MYLIBID
	result.ProgramID = M3UINTCONST
	result.ProgramValue = op_result
	result.NextPrograms = nil

	return result, nil
}
