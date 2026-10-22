package m3numbertriginv

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/mmirko/mel/pkg/m3number"
	mel3program "github.com/mmirko/mel/pkg/mel3program"
)

type Evaluator struct {
	*mel3program.Mel3Object
	error
	Result *mel3program.Mel3Program
}

func EvaluatorCreator() mel3program.Mel3Visitor {
	return new(Evaluator)
}

func (ev *Evaluator) GetName() string {
	return "m3numbertriginv"
}

func (ev *Evaluator) GetMel3Object() *mel3program.Mel3Object {
	return ev.Mel3Object
}

func (ev *Evaluator) SetMel3Object(object *mel3program.Mel3Object) {
	ev.Mel3Object = object
}

func (ev *Evaluator) GetError() error {
	return ev.error
}

func (ev *Evaluator) GetResult() *mel3program.Mel3Program {
	return ev.Result
}

func (ev *Evaluator) Visit(program *mel3program.Mel3Program) mel3program.Mel3Visitor {
	checkEvaluator := mel3program.ProgMux(ev, program)
	if ev.GetName() != checkEvaluator.GetName() {
		return checkEvaluator.Visit(program)
	}

	if program.LibraryID != MYLIBID {
		ev.error = errors.New("unknown LibraryID")
		return nil
	}

	arguments := make([]mel3program.Mel3Visitor, len(program.NextPrograms))
	for index, argument := range program.NextPrograms {
		arguments[index] = mel3program.ProgMux(ev, argument)
		arguments[index].Visit(argument)
	}
	if len(arguments) != 1 {
		ev.error = errors.New("wrong argument number")
		return nil
	}

	argument := arguments[0].GetResult()
	if argument == nil || argument.LibraryID != m3number.MYLIBID || argument.ProgramID != m3number.M3NUMBERCONST {
		ev.error = errors.New("wrong argument type")
		return nil
	}

	value, err := strconv.ParseFloat(argument.ProgramValue, 32)
	if err != nil {
		ev.error = errors.New("convert to number failed")
		return nil
	}
	input := float64(float32(value))
	if math.IsNaN(input) {
		ev.error = errors.New("input outside function domain")
		return nil
	}

	var result float64
	switch program.ProgramID {
	case ASIN:
		if input < -1 || input > 1 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Asin(input)
	case ACOS:
		if input < -1 || input > 1 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Acos(input)
	case ATAN:
		result = math.Atan(input)
	case ACSC:
		if math.Abs(input) < 1 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Asin(1 / input)
	case ASEC:
		if math.Abs(input) < 1 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Acos(1 / input)
	case ACOT:
		result = math.Atan2(1, input)
	default:
		ev.error = errors.New("unknown ProgramID")
		return nil
	}

	ev.Result = &mel3program.Mel3Program{
		LibraryID:    m3number.MYLIBID,
		ProgramID:    m3number.M3NUMBERCONST,
		ProgramValue: strconv.FormatFloat(float64(float32(result)), 'E', -1, 32),
	}
	return nil
}

func (ev *Evaluator) Inspect() string {
	if ev.error != nil {
		return fmt.Sprint(ev.error)
	}
	if dump, err := mel3program.ProgDump(ev.Implementation, ev.Result); err == nil {
		return "Evaluation ok: " + dump
	}
	return "Result export failed"
}
