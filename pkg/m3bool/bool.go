package m3bool

import (
	//"math/rand"
	//"fmt"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

// Program IDs
const (
	CONST = uint16(0) + iota
	VAR
	NOT
	AND
	OR
	XOR
	NAND
	NOR
	XNOR
)

// Program types
const (
	BOOL = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_M3BOOL
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		CONST: "m3boolconst",
		VAR:   "m3boolvar",
		NOT:   "not",
		AND:   "and",
		OR:    "or",
		XOR:   "xor",
		NAND:  "nand",
		NOR:   "nor",
		XNOR:  "xnor",
	},
	TypeNames: map[uint16]string{
		BOOL: "bool",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		CONST: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		VAR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		NOT:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		AND:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		OR:    mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		XOR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		NAND:  mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		NOR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		XNOR:  mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		CONST: mel3program.ArgumentsTypes{},
		VAR:   mel3program.ArgumentsTypes{},
		NOT:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		AND:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		OR:    mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		XOR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		NAND:  mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		NOR:   mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
		XNOR:  mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, BOOL, []uint64{}}, mel3program.ArgType{MYLIBID, BOOL, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		CONST: false,
		VAR:   false,
		NOT:   false,
		AND:   false,
		OR:    false,
		XOR:   false,
		NAND:  false,
		NOR:   false,
		XNOR:  false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		CONST: mel3program.ArgType{},
		VAR:   mel3program.ArgType{},
		NOT:   mel3program.ArgType{},
		AND:   mel3program.ArgType{},
		OR:    mel3program.ArgType{},
		XOR:   mel3program.ArgType{},
		NAND:  mel3program.ArgType{},
		NOR:   mel3program.ArgType{},
		XNOR:  mel3program.ArgType{},
	},
	ImplName: "m3bool",
}

// The effective Me3li
type M3boolMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for M3boolMe3li
func (prog *M3boolMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation

	creators := make(map[uint16]mel3program.Mel3VisitorCreator)
	creators[MYLIBID] = EvaluatorCreator

	prog.Mel3Init(c, implementations, creators, ep)
}

func (prog *M3boolMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
