//go:build GODEBUG
// +build GODEBUG

// This is the MEL3 implementation of the numbers
package m3number

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

// Program IDs
const (
	M3NUMBERCONST = uint16(0) + iota
	ADD
	SUB
	MULT
	DIV
)

// Program types
const (
	M3NUMBER = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_M3NUMBER
)

// The Mel3 implementation
var Implementation = mel3program.Mel3_implementation{
	ProgramNames: map[uint16]string{
		M3NUMBERCONST: "m3numberconst",
		ADD:           "add",
		SUB:           "sub",
		MULT:          "mult",
		DIV:           "div",
	},
	TypeNames: map[uint16]string{
		M3NUMBER: "uint",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		M3NUMBERCONST: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		ADD:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		SUB:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		MULT:          mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		DIV:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		M3NUMBERCONST: mel3program.ArgumentsTypes{},
		ADD:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		SUB:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		MULT:          mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
		DIV:           mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}, mel3program.ArgType{MYLIBID, M3NUMBER, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		M3NUMBERCONST: false,
		ADD:           false,
		SUB:           false,
		MULT:          false,
		DIV:           false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		M3NUMBERCONST: mel3program.ArgType{},
		ADD:           mel3program.ArgType{},
		SUB:           mel3program.ArgType{},
		MULT:          mel3program.ArgType{},
		DIV:           mel3program.ArgType{},
	},
	Implname: "m3number",
}

// The effective Me3li
type M3number_me3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for M3number_me3li
func (prog *M3number_me3li) Mel_init(ep *mel.Evolution_parameters) {
	impls := make(map[uint16]*mel3program.Mel3_implementation)
	impls[MYLIBID] = &Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *M3number_me3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
