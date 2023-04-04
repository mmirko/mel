package m3melbond

import (
	//"math/rand"
	//"fmt"

	"github.com/mmirko/mel/pkg/m3bool"
	"github.com/mmirko/mel/pkg/m3boolcmp"
	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/m3statements"
	"github.com/mmirko/mel/pkg/m3uint"
	"github.com/mmirko/mel/pkg/m3uintcmp"
	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

// Program IDs
const ()

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_M3MELBOND
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames:    map[uint16]string{},
	TypeNames:       map[uint16]string{},
	ProgramTypes:    map[uint16]mel3program.ArgumentsTypes{},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{},
	IsVariadic:      map[uint16]bool{},
	VariadicType:    map[uint16]mel3program.ArgType{},
	ImplName:        "m3melbond",
}

// The effective Me3li
type M3MelBondMe3li struct {
	mel3program.Mel3Object
	libs []string
}

func (l *M3MelBondMe3li) Init(c *mel.MelConfig, ep *mel.EvolutionParameters, libs []string) error {

	if checked, err := mel3program.LibsCheckAndRequirements(libs); err != nil {
		return err
	} else {
		l.libs = make([]string, len(checked))
		copy(l.libs, checked)
	}

	l.MelInit(c, ep)
	return nil
}

// ********* Mel interface

// The Mel entry point for M3uintMe3li
func (prog *M3MelBondMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
	implementations := make(map[uint16]*mel3program.Mel3Implementation)
	implementations[MYLIBID] = &Implementation

	for _, lib := range prog.libs {
		switch lib {
		case "m3uint":
			implementations[m3uint.MYLIBID] = &m3uint.Implementation
		case "m3uintcmp":
			implementations[m3uintcmp.MYLIBID] = &m3uintcmp.Implementation
		case "m3number":
			implementations[m3number.MYLIBID] = &m3number.Implementation
		case "m3bool":
			implementations[m3bool.MYLIBID] = &m3bool.Implementation
		case "m3boolcmp":
			implementations[m3boolcmp.MYLIBID] = &m3boolcmp.Implementation
		case "m3statements":
			implementations[m3statements.MYLIBID] = &m3statements.Implementation
		}
	}

	if prog.Mel3Object.DefaultCreator == nil {

		creators := make(map[uint16]mel3program.Mel3VisitorCreator)
		creators[MYLIBID] = EvaluatorCreator

		for _, lib := range prog.libs {
			switch lib {
			case "m3uint":
				creators[m3uint.MYLIBID] = m3uint.EvaluatorCreator
			case "m3uintcmp":
				creators[m3uintcmp.MYLIBID] = m3uintcmp.EvaluatorCreator
			case "m3number":
				creators[m3number.MYLIBID] = m3number.EvaluatorCreator
			case "m3bool":
				creators[m3bool.MYLIBID] = m3bool.EvaluatorCreator
			case "m3boolcmp":
				creators[m3boolcmp.MYLIBID] = m3boolcmp.EvaluatorCreator
			case "m3statements":
				creators[m3statements.MYLIBID] = m3statements.EvaluatorCreator
			}
		}
		prog.Mel3Init(c, implementations, creators, ep)
	} else {
		creators := mel3program.CreateGenericCreators(&prog.Mel3Object, ep, implementations)
		prog.Mel3Init(c, implementations, creators, ep)
	}

}

func (prog *M3MelBondMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
