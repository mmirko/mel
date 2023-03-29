package m3uintcmp

import (
	//"math/rand"
	//"fmt"

	m3bool "github.com/mmirko/mel/pkg/m3bool"
	m3uint "github.com/mmirko/mel/pkg/m3uint"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
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
var Implementation = mel3program.Mel3Implementation{
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
		EQ: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
		NE: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
		LT: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
		LE: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
		GT: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
		GE: mel3program.ArgumentsTypes{mel3program.ArgType{m3bool.MYLIBID, m3bool.BOOL, []uint64{}}},
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
	ImplName: "m3uintcmp",
}

// The effective Me3li
type M3uintcmpMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for M3uintcmpMe3li
func (prog *M3uintcmpMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation
	implementations[m3uint.MYLIBID] = &m3uint.Implementation
	implementations[m3bool.MYLIBID] = &m3bool.Implementation

	creators := make(map[uint16]mel3program.Mel3VisitorCreator)
	creators[MYLIBID] = EvaluatorCreator
	creators[m3uint.MYLIBID] = m3uint.EvaluatorCreator
	creators[m3bool.MYLIBID] = m3bool.EvaluatorCreator

	prog.Mel3Init(c, implementations, creators, ep)
}

func (prog *M3uintcmpMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
