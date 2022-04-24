package m3redmon

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	m3bool "github.com/mmirko/mel/m3bool"
	m3dates "github.com/mmirko/mel/m3dates"
	m3uint "github.com/mmirko/mel/m3uint"
	m3uintcmp "github.com/mmirko/mel/m3uintcmp"
	mel3program "github.com/mmirko/mel/mel3program"
)

// Program IDs
const ()

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_M3REDMON
)

// The Mel3 implementation
var Implementation = mel3program.Mel3_implementation{
	ProgramNames:    map[uint16]string{},
	TypeNames:       map[uint16]string{},
	ProgramTypes:    map[uint16]mel3program.ArgumentsTypes{},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{},
	IsVariadic:      map[uint16]bool{},
	VariadicType:    map[uint16]mel3program.ArgType{},
	Implname:        "m3redmon",
}

// The effective Me3li
type M3redmonMe3li struct {
	mel3program.Mel3_object
}

// ********* Mel interface

// The Mel entry point for M3redmonMe3li
func (prog *M3redmonMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3_implementation)
	impls[MYLIBID] = &Implementation
	impls[m3bool.MYLIBID] = &m3bool.Implementation
	impls[m3dates.MYLIBID] = &m3dates.Implementation
	impls[m3uint.MYLIBID] = &m3uint.Implementation
	impls[m3uintcmp.MYLIBID] = &m3uintcmp.Implementation
	prog.Mel3_init(impls, ep)
}

func (prog *M3redmonMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
