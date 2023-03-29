package envfloat

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
const (
	READINPUT = uint16(0) + iota
	READOUTPUT
	READKEEP
	WRITEOUTPUT
)

// Program types
const ()

const (
	MYLIBID = mel3program.LIB_ENVFLOAT
)

// The Mel3 implementation
var Implementation = mel3program.Mel3Implementation{
	ProgramNames: map[uint16]string{
		READINPUT:   "readinput",
		READOUTPUT:  "readoutput",
		READKEEP:    "readkeep",
		WRITEOUTPUT: "writeoutput",
	},
	TypeNames: map[uint16]string{},
	ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
		READINPUT:   mel3program.ArgumentsTypes{mel3program.ArgType{m3number.MYLIBID, m3number.M3NUMBER, []uint64{}}},
		READOUTPUT:  mel3program.ArgumentsTypes{mel3program.ArgType{m3number.MYLIBID, m3number.M3NUMBER, []uint64{}}},
		READKEEP:    mel3program.ArgumentsTypes{mel3program.ArgType{m3number.MYLIBID, m3number.M3NUMBER, []uint64{}}},
		WRITEOUTPUT: mel3program.ArgumentsTypes{mel3program.ArgType{m3number.MYLIBID, m3number.M3NUMBER, []uint64{}}},
	},
	NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
		READINPUT:   mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		READOUTPUT:  mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		READKEEP:    mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}},
		WRITEOUTPUT: mel3program.ArgumentsTypes{mel3program.ArgType{m3uint.MYLIBID, m3uint.M3UINT, []uint64{}}, mel3program.ArgType{m3number.MYLIBID, m3number.M3NUMBER, []uint64{}}},
	},
	IsVariadic: map[uint16]bool{
		READINPUT: false,
	},
	VariadicType: map[uint16]mel3program.ArgType{
		READINPUT:   mel3program.ArgType{},
		READOUTPUT:  mel3program.ArgType{},
		READKEEP:    mel3program.ArgType{},
		WRITEOUTPUT: mel3program.ArgType{},
	},
	ImplName: "envfloat",
}

// The effective Me3li
type EnvFloatMe3li struct {
	mel3program.Mel3Object
	libs []string
}

func (l *EnvFloatMe3li) Init(c *mel.MelConfig, ep *mel.EvolutionParameters, libs []string) error {

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
func (prog *EnvFloatMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
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
}

func (prog *EnvFloatMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
