package lisevo

import (
	//"math/rand"
	//"fmt"
	mel "github.com/mmirko/mel"
	"github.com/mmirko/mel/m3bool"
	"github.com/mmirko/mel/m3number"
	"github.com/mmirko/mel/m3statements"
	"github.com/mmirko/mel/m3uint"
	"github.com/mmirko/mel/m3uintcmp"
	mel3program "github.com/mmirko/mel/mel3program"
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

// The Mel entry point for M3uintMe3li
func (prog *LisevoMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation
	implementations[m3bool.MYLIBID] = &m3bool.Implementation
	implementations[m3number.MYLIBID] = &m3number.Implementation
	implementations[m3uint.MYLIBID] = &m3uint.Implementation
	implementations[m3uintcmp.MYLIBID] = &m3uintcmp.Implementation
	implementations[m3statements.MYLIBID] = &m3statements.Implementation

	creators := make(map[uint16]mel3program.Mel3VisitorCreator)
	creators[MYLIBID] = EvaluatorCreator
	creators[m3bool.MYLIBID] = m3bool.EvaluatorCreator
	creators[m3number.MYLIBID] = m3number.EvaluatorCreator
	creators[m3uint.MYLIBID] = m3uint.EvaluatorCreator
	creators[m3uintcmp.MYLIBID] = m3uintcmp.EvaluatorCreator
	creators[m3statements.MYLIBID] = m3statements.EvaluatorCreator

	prog.Mel3Init(c, implementations, creators, ep)
}

func (prog *LisevoMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
