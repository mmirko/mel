package rectangular

import (
	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
	statements "github.com/mmirko/mel/statements"
)

// Program IDs
const (
	RECTCONST = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_RECTANGULAR
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		RECTCONST: "rectconst",
	},
	TypeNames: map[uint16]string{},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		RECTCONST: mel3program.ArgumentsTypes{mel3program.ArgType{statements.MYLIBID, statements.MULTISTMT, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		RECTCONST: mel3program.ArgumentsTypes{},
	},
	IsVariadic: map[uint16]bool{
		RECTCONST: false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		RECTCONST: mel3program.ArgType{},
	},
	ImplName: "rectangular",
}

type RectangularMe3li struct {
	mel3program.Mel3Object
}

// The Mel entry point for Symbolic_math_me3li
func (prog *RectangularMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3Implementation)
	impls[MYLIBID] = &Implementation
	impls[statements.MYLIBID] = &statements.Implementation
	prog.Mel3Init(impls, ep)
}

func (prog *RectangularMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
