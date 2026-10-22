package m3numbertriginv

import (
	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

const (
	ASIN = uint16(0) + iota
	ACOS
	ATAN
	ACSC
	ASEC
	ACOT
)

const MYLIBID = mel3program.LIB_M3NUMBERTRIGINV

var Implementation = func() mel3program.Mel3Implementation {
	numberType := mel3program.ArgType{LibraryID: m3number.MYLIBID, TypeID: m3number.M3NUMBER, TensorRapr: []uint64{}}

	return mel3program.Mel3Implementation{
		ProgramNames: map[uint16]string{
			ASIN: "asin",
			ACOS: "acos",
			ATAN: "atan",
			ACSC: "acsc",
			ASEC: "asec",
			ACOT: "acot",
		},
		TypeNames: map[uint16]string{},
		ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
			ASIN: {numberType},
			ACOS: {numberType},
			ATAN: {numberType},
			ACSC: {numberType},
			ASEC: {numberType},
			ACOT: {numberType},
		},
		NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
			ASIN: {numberType},
			ACOS: {numberType},
			ATAN: {numberType},
			ACSC: {numberType},
			ASEC: {numberType},
			ACOT: {numberType},
		},
		IsVariadic: map[uint16]bool{
			ASIN: false,
			ACOS: false,
			ATAN: false,
			ACSC: false,
			ASEC: false,
			ACOT: false,
		},
		VariadicType: map[uint16]mel3program.ArgType{
			ASIN: {},
			ACOS: {},
			ATAN: {},
			ACSC: {},
			ASEC: {},
			ACOT: {},
		},
		ImplName: "m3numbertriginv",
	}
}()

type M3numberTrigInvMe3li struct {
	mel3program.Mel3Object
}

func (prog *M3numberTrigInvMe3li) MelInit(config *mel.MelConfig, parameters *mel.EvolutionParameters) {
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

func (prog *M3numberTrigInvMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
