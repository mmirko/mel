package m3symbolic

import (
	//"math/rand"
	//"fmt"
	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

// Program IDs
const (
	CONST = uint16(0) + iota
	VAR
	SUM
	MUL
)

// Program types
const (
	NUMBER = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_M3SYMBOLIC
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		CONST: "const",
		VAR:   "var",
		SUM:   "sum",
		MUL:   "mul",
	},
	TypeNames: map[uint16]string{
		NUMBER: "number",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		CONST: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
		VAR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
		SUM:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
		MUL:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		CONST: mel3program.ArgumentsTypes{},
		VAR:   mel3program.ArgumentsTypes{},
		SUM:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
		MUL:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, NUMBER, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		CONST: false,
		VAR:   false,
		SUM:   false,
		MUL:   false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		CONST: mel3program.ArgType{},
		VAR:   mel3program.ArgType{},
		SUM:   mel3program.ArgType{},
		MUL:   mel3program.ArgType{},
	},
	ImplName: "m3symbolic",
}

// The effective Me3li
type Symbolic_math3_me3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for Symbolic_math_me3li
func (prog *Symbolic_math3_me3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3Implementation)
	impls[MYLIBID] = &Implementation
	prog.Mel3Init(impls, ep)
}

func (prog *Symbolic_math3_me3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
