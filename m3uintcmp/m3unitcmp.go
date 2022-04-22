package m3uintcmp

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	bool3 "github.com/mmirko/mel/bool3"
	m3uint "github.com/mmirko/mel/m3uint"
	mel3program "github.com/mmirko/mel/mel3program"
)

// Program IDs
const (
	EQ = uint16(0) + iota
	NE
	LT
	LE
	GT
	GE
)

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_M3UINTCMP
)

// The Mel3 implementation
var Implementation = mel3program.Mel3_implementation{
	ProgramNames: map[uint16]string{
		EQ: "eq",
		NE: "ne",
		LT: "lt",
		LE: "le",
		GT: "gt",
		GE: "ge",
	},
	TypeNames: map[uint16]string{},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		EQ: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
		NE: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
		LT: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
		LE: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
		GT: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
		GE: mel3program.ArgumentsTypes{mel3program.ArgType{bool3.MYLIBID, bool3.BOOL, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		EQ: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		NE: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		LT: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		LE: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		GT: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		GE: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		EQ: false,
		NE: false,
		LT: false,
		LE: false,
		GT: false,
		GE: false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		EQ: mel3program.ArgType{},
		NE: mel3program.ArgType{},
		LT: mel3program.ArgType{},
		LE: mel3program.ArgType{},
		GT: mel3program.ArgType{},
		GE: mel3program.ArgType{},
	},
	Implname: "m3uintcmp",
}

// The effective Me3li
type M3uintcmp_me3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for M3uintcmp_me3li
func (prog *M3uintcmp_me3li) Mel_init(ep *mel.Evolution_parameters) {
	impls := make(map[uint16]*mel3program.Mel3_implementation)
	impls[MYLIBID] = &Implementation
	impls[m3uint.MYLIBID] = &m3uint.Implementation
	impls[bool3.MYLIBID] = &bool3.Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *M3uintcmp_me3li) Mel_copy() mel.Me3li {
	var result mel.Me3li
	return result
}
