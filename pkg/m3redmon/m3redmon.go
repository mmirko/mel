package m3redmon

import (
	//"math/rand"
	//"fmt"
	m3bool "github.com/mmirko/mel/pkg/m3bool"
	m3dates "github.com/mmirko/mel/pkg/m3dates"
	m3uint "github.com/mmirko/mel/pkg/m3uint"
	m3uintcmp "github.com/mmirko/mel/pkg/m3uintcmp"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

// Program IDs
const ()

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_M3REDMON
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
type M3redmonMe3li struct {
	mel3program.Mel3Object
}

// ********* Mel interface

// The Mel entry point for M3redmonMe3li
func (prog *M3redmonMe3li) MelInit(ep *mel.EvolutionParameters) {
	impls := make(map[uint16]*mel3program.Mel3Implementation)
	impls[MYLIBID] = &Implementation
	impls[m3bool.MYLIBID] = &m3bool.Implementation
	impls[m3dates.MYLIBID] = &m3dates.Implementation
	impls[m3uint.MYLIBID] = &m3uint.Implementation
	impls[m3uintcmp.MYLIBID] = &m3uintcmp.Implementation
	prog.Mel3Init(impls, ep)
}

func (prog *M3redmonMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
