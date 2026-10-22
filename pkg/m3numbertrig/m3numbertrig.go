package m3numbertrig

import (
	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

const (
	SIN = uint16(0) + iota
	COS
	TAN
	CSC
	SEC
	COT
)

const MYLIBID = mel3program.LIB_M3NUMBERTRIG

var Implementation = func() mel3program.Mel3Implementation {
	numberType := mel3program.ArgType{LibraryID: m3number.MYLIBID, TypeID: m3number.M3NUMBER, TensorRapr: []uint64{}}

	return mel3program.Mel3Implementation{
		ProgramNames: map[uint16]string{
			SIN: "sin",
			COS: "cos",
			TAN: "tan",
			CSC: "csc",
			SEC: "sec",
			COT: "cot",
		},
		TypeNames: map[uint16]string{},
		ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
			SIN: {numberType},
			COS: {numberType},
			TAN: {numberType},
			CSC: {numberType},
			SEC: {numberType},
			COT: {numberType},
		},
		NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
			SIN: {numberType},
			COS: {numberType},
			TAN: {numberType},
			CSC: {numberType},
			SEC: {numberType},
			COT: {numberType},
		},
		IsVariadic: map[uint16]bool{
			SIN: false,
			COS: false,
			TAN: false,
			CSC: false,
			SEC: false,
			COT: false,
		},
		VariadicType: map[uint16]mel3program.ArgType{
			SIN: {},
			COS: {},
			TAN: {},
			CSC: {},
			SEC: {},
			COT: {},
		},
		ImplName: "m3numbertrig",
	}
}()

type M3numberTrigMe3li struct {
	mel3program.Mel3Object
}

func (prog *M3numberTrigMe3li) MelInit(config *mel.MelConfig, parameters *mel.EvolutionParameters) {
	implementations := map[uint16]*mel3program.Mel3Implementation{
		MYLIBID:          &Implementation,
		m3number.MYLIBID: &m3number.Implementation,
	}

	if prog.Mel3Object.DefaultCreator == nil {
		creators := map[uint16]mel3program.Mel3VisitorCreator{
			MYLIBID:          EvaluatorCreator,
			m3number.MYLIBID: m3number.EvaluatorCreator,
		}
		prog.Mel3Init(config, implementations, creators, parameters)
	} else {
		creators := mel3program.CreateGenericCreators(&prog.Mel3Object, parameters, implementations)
		prog.Mel3Init(config, implementations, creators, parameters)
	}
}

func (prog *M3numberTrigMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
