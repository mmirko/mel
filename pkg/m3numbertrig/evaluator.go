package m3numbertrig

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
	return "m3numbertrig"
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

	var result float64
	switch program.ProgramID {
	case SIN:
		result = math.Sin(input)
	case COS:
		result = math.Cos(input)
	case TAN:
		result = math.Tan(input)
	case CSC:
		denominator := math.Sin(input)
		if denominator == 0 {
			ev.error = errors.New("division by zero")
			return nil
		}
		result = 1 / denominator
	case SEC:
		denominator := math.Cos(input)
		if denominator == 0 {
			ev.error = errors.New("division by zero")
			return nil
		}
		result = 1 / denominator
	case COT:
		denominator := math.Sin(input)
		if denominator == 0 {
			ev.error = errors.New("division by zero")
			return nil
		}
		result = math.Cos(input) / denominator
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
