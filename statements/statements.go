package statements

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

// Program IDs
const (
	MULTISTMT = uint16(0) + iota
	NOP
)

// Program types
const (
	STMT = uint16(0) + iota
)

const (
	MYLIBID = mel3program.LIB_STATEMENTS
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		MULTISTMT: "multistmt", // Multiline statement
		NOP:       "nop",       // void instruction
	},
	TypeNames: map[uint16]string{
		STMT: "stmt",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		MULTISTMT: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, STMT, []uint64{}}},
		NOP:       mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, STMT, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		MULTISTMT: mel3program.ArgumentsTypes{},
		NOP:       mel3program.ArgumentsTypes{},
	},
	IsVariadic: map[uint16]bool{
		MULTISTMT: true,
		NOP:       false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		MULTISTMT: mel3program.ArgType{MYLIBID, STMT, []uint64{}},
		NOP:       mel3program.ArgType{},
	},
	Implname: "statements",
}

// The effective Me3li
type StatementsMe3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for Symbolic_math_me3li
func (prog *StatementsMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3Implementation)
	impls[MYLIBID] = &Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *StatementsMe3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
