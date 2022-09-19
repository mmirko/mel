package lisevo

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	m3bool "github.com/mmirko/mel/m3bool"
	m3number "github.com/mmirko/mel/m3number"
	m3uint "github.com/mmirko/mel/m3uint"
	m3uintcmp "github.com/mmirko/mel/m3uintcmp"
	mel3program "github.com/mmirko/mel/mel3program"
	statements "github.com/mmirko/mel/statements"
)

// Program IDs
const ()

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_LISEVO
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames:    map[uint16]string{},
	TypeNames:       map[uint16]string{},
	ProgramTypes:    map[uint16]mel3program.ArgumentsTypes{},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{},
	IsVariadic:      map[uint16]bool{},
	VariadicType:    map[uint16]mel3program.ArgType{},
	ImplName:        "m3redmon",
}

// The effective Me3li
type LisevoMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for LisevoMe3li
func (prog *LisevoMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3Implementation)
	impls[MYLIBID] = &Implementation
	impls[m3bool.MYLIBID] = &m3bool.Implementation
	impls[statements.MYLIBID] = &statements.Implementation
	impls[m3number.MYLIBID] = &m3number.Implementation
	impls[m3uint.MYLIBID] = &m3uint.Implementation
	impls[m3uintcmp.MYLIBID] = &m3uintcmp.Implementation
	prog.Mel3Init(impls, ep)
}

func (prog *LisevoMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
