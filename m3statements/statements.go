package m3statements

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
	MYLIBID = mel3program.LIB_M3STATEMENTS
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
	ImplName: "m3statements",
}

// The effective Me3li
type M3statementsMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for M3uintMe3li
func (prog *M3statementsMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation

	creators := make(map[uint16]mel3program.Mel3VisitorCreator)
	creators[MYLIBID] = EvaluatorCreator

	prog.Mel3Init(c, implementations, creators, ep)
}

func (prog *M3statementsMe3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
