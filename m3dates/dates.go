package m3dates

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	m3uint "github.com/mmirko/mel/m3uint"
	mel3program "github.com/mmirko/mel/mel3program"
)

// Program IDs
const (
	DATECONST = uint16(0) + iota
	TIMESTAMPCONST
	DIFFDAYS
)

// Program types
const (
	DATE = uint16(0) + iota
	TIMESTAMP
)

const (
	MYLIBID = mel3program.LIB_M3DATES
)

// The Mel3 implementation
var Implementation = mel3program.Mel3_implementation{
	ProgramNames: map[uint16]string{
		DATECONST:      "dateconst",
		TIMESTAMPCONST: "timestampconst",
		DIFFDAYS:       "diffdays",
	},
	TypeNames: map[uint16]string{
		DATE:      "date",
		TIMESTAMP: "timestamp",
	},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		DATECONST:      mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, DATE, []uint64{}}},
		TIMESTAMPCONST: mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, TIMESTAMP, []uint64{}}},
		DIFFDAYS:       mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		DATECONST:      mel3program.ArgumentsTypes{},
		TIMESTAMPCONST: mel3program.ArgumentsTypes{},
		DIFFDAYS:       mel3program.ArgumentsTypes{mel3program.ArgType{MYLIBID, DATE, []uint64{}}, mel3program.ArgType{MYLIBID, DATE, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		DATECONST:      false,
		TIMESTAMPCONST: false,
		DIFFDAYS:       false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		DATECONST:      mel3program.ArgType{},
		TIMESTAMPCONST: mel3program.ArgType{},
		DIFFDAYS:       mel3program.ArgType{},
	},
	Implname: "m3dates",
}

// The effective Me3li
type M3datesMe3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for M3datesMe3li
func (prog *M3datesMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3_implementation)
	impls[MYLIBID] = &Implementation
	impls[m3uint.MYLIBID] = &m3uint.Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *M3datesMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
