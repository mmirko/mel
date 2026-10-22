package m3numberext

import (
	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/m3uint"
	"github.com/mmirko/mel/pkg/mel"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

const (
	SQRT = uint16(0) + iota
	POW
	EXP
	LOG
	LOG2
	LOG10
	ABS
	CEIL
	FLOOR
	ROUND
	TRUNC
)

const MYLIBID = mel3program.LIB_M3NUMBEREXT

var Implementation = func() mel3program.Mel3Implementation {
	numberType := mel3program.ArgType{LibraryID: m3number.MYLIBID, TypeID: m3number.M3NUMBER, TensorRapr: []uint64{}}
	integerType := mel3program.ArgType{LibraryID: m3uint.MYLIBID, TypeID: m3uint.M3UINT, TensorRapr: []uint64{}}

	return mel3program.Mel3Implementation{
		ProgramNames: map[uint16]string{
			SQRT:  "sqrt",
			POW:   "pow",
			EXP:   "exp",
			LOG:   "log",
			LOG2:  "log2",
			LOG10: "log10",
			ABS:   "abs",
			CEIL:  "ceil",
			FLOOR: "floor",
			ROUND: "round",
			TRUNC: "trunc",
		},
		TypeNames: map[uint16]string{},
		ProgramTypes: map[uint16]mel3program.ArgumentsTypes{
			SQRT:  {numberType},
			POW:   {numberType},
			EXP:   {numberType},
			LOG:   {numberType},
			LOG2:  {numberType},
			LOG10: {numberType},
			ABS:   {numberType},
			CEIL:  {numberType},
			FLOOR: {numberType},
			ROUND: {numberType},
			TRUNC: {numberType},
		},
		NonVariadicArgs: map[uint16]mel3program.ArgumentsTypes{
			SQRT:  {numberType},
			POW:   {numberType, numberType},
			EXP:   {numberType},
			LOG:   {numberType},
			LOG2:  {numberType},
			LOG10: {numberType},
			ABS:   {numberType},
			CEIL:  {numberType},
			FLOOR: {numberType},
			ROUND: {numberType, integerType},
			TRUNC: {numberType},
		},
		IsVariadic: map[uint16]bool{
			SQRT: false, POW: false, EXP: false, LOG: false, LOG2: false, LOG10: false,
			ABS: false, CEIL: false, FLOOR: false, ROUND: false, TRUNC: false,
		},
		VariadicType: map[uint16]mel3program.ArgType{
			SQRT: {}, POW: {}, EXP: {}, LOG: {}, LOG2: {}, LOG10: {},
			ABS: {}, CEIL: {}, FLOOR: {}, ROUND: {}, TRUNC: {},
		},
		ImplName: "m3numberext",
	}
}()

type M3numberExtMe3li struct {
	mel3program.Mel3Object
}

func (prog *M3numberExtMe3li) MelInit(config *mel.MelConfig, parameters *mel.EvolutionParameters) {
	implementations := map[uint16]*mel3program.Mel3Implementation{
		MYLIBID:          &Implementation,
		m3number.MYLIBID: &m3number.Implementation,
		m3uint.MYLIBID:   &m3uint.Implementation,
	}

	if prog.Mel3Object.DefaultCreator == nil {
		creators := map[uint16]mel3program.Mel3VisitorCreator{
			MYLIBID:          EvaluatorCreator,
			m3number.MYLIBID: m3number.EvaluatorCreator,
			m3uint.MYLIBID:   m3uint.EvaluatorCreator,
		}
		prog.Mel3Init(config, implementations, creators, parameters)
	} else {
		creators := mel3program.CreateGenericCreators(&prog.Mel3Object, parameters, implementations)
		prog.Mel3Init(config, implementations, creators, parameters)
	}
}

func (prog *M3numberExtMe3li) MelCopy() mel.Me3li {
	var result mel.Me3li
	return result
}
