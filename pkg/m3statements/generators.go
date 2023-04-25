package m3statements

import (
	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

func M3StmtNopGenerator(ev *mel.EvolutionParameters) (*mel3program.Mel3Program, error) {
	result := new(mel3program.Mel3Program)
	result.LibraryID = MYLIBID
	result.ProgramID = STMT
	result.ProgramValue = "nop"
	result.NextPrograms = nil

	return result, nil
}
