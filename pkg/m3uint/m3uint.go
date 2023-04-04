// This is the MEL3 implementation of the uint numbers
package m3uint

import (
	//"math/rand"
	//"fmt"
	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

// Program IDs
const (
	M3UINTCONST = uint16(0) + iota
	ADD
	SUB
	MULT
	DIV
)

// Program types
const (
	M3UINT = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_M3UINT
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		M3UINTCONST: "m3uintconst",
		ADD:         "add",
		SUB:         "sub",
		MULT:        "mult",
		DIV:         "div",
	},
	TypeNames: map[uint16]string{
		M3UINT: "uint",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		M3UINTCONST: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		ADD:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		SUB:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		MULT:        mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		DIV:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		M3UINTCONST: mel3program.ArgumentsTypes{},
		ADD:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}, mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		SUB:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}, mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		MULT:        mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}, mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
		DIV:         mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}, mel3program.ArgType{MYLIBID, M3UINT, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		M3UINTCONST: false,
		ADD:         false,
		SUB:         false,
		MULT:        false,
		DIV:         false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		M3UINTCONST: mel3program.ArgType{},
		ADD:         mel3program.ArgType{},
		SUB:         mel3program.ArgType{},
		MULT:        mel3program.ArgType{},
		DIV:         mel3program.ArgType{},
	},
	ImplName: "m3uint",
}

// The effective Me3li
type M3uintMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for M3uintMe3li
func (prog *M3uintMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation

	creators := make(map[uint16]mel3program.Mel3VisitorCreator)
	creators[MYLIBID] = EvaluatorCreator

	prog.Mel3Init(c, implementations, creators, ep)
}

func (prog *M3uintMe3li) MelCopy(c *mel.MelConfig) mel.Me3li {
	var result mel.Me3li
	return result
}
