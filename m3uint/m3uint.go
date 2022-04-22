// This is the MEL3 implementation of the uint numbers
package m3uint

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
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
var Implementation = mel3program.Mel3_implementation{
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
	Implname: "m3uint",
}

// The effective Me3li
type M3uint_me3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for M3uint_me3li
func (prog *M3uint_me3li) Mel_init(ep *mel.Evolution_parameters) {
	impls := make(map[uint16]*mel3program.Mel3_implementation)
	impls[MYLIBID] = &Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *M3uint_me3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
