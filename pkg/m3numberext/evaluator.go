package m3numberext

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/m3uint"
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
	return "m3numberext"
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

	expectedArguments := 1
	if program.ProgramID == POW || program.ProgramID == ROUND {
		expectedArguments = 2
	}
	if len(program.NextPrograms) != expectedArguments {
		ev.error = errors.New("wrong argument number")
		return nil
	}

	arguments := make([]mel3program.Mel3Visitor, len(program.NextPrograms))
	for index, argument := range program.NextPrograms {
		arguments[index] = mel3program.ProgMux(ev, argument)
		arguments[index].Visit(argument)
	}

	numberArgument := arguments[0].GetResult()
	if numberArgument == nil || numberArgument.LibraryID != m3number.MYLIBID || numberArgument.ProgramID != m3number.M3NUMBERCONST {
		ev.error = errors.New("wrong argument type")
		return nil
	}
	numberValue, err := strconv.ParseFloat(numberArgument.ProgramValue, 32)
	if err != nil {
		ev.error = errors.New("convert to number failed")
		return nil
	}
	input := float64(float32(numberValue))

	var result float64
	switch program.ProgramID {
	case SQRT:
		if input < 0 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Sqrt(input)
	case POW:
		powerArgument := arguments[1].GetResult()
		if powerArgument == nil || powerArgument.LibraryID != m3number.MYLIBID || powerArgument.ProgramID != m3number.M3NUMBERCONST {
			ev.error = errors.New("wrong argument type")
			return nil
		}
		powerValue, parseErr := strconv.ParseFloat(powerArgument.ProgramValue, 32)
		if parseErr != nil {
			ev.error = errors.New("convert to number failed")
			return nil
		}
		result = math.Pow(input, float64(float32(powerValue)))
	case EXP:
		result = math.Exp(input)
	case LOG:
		if input <= 0 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Log(input)
	case LOG2:
		if input <= 0 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Log2(input)
	case LOG10:
		if input <= 0 {
			ev.error = errors.New("input outside function domain")
			return nil
		}
		result = math.Log10(input)
	case ABS:
		result = math.Abs(input)
	case CEIL:
		result = math.Ceil(input)
	case FLOOR:
		result = math.Floor(input)
	case ROUND:
		precisionArgument := arguments[1].GetResult()
		if precisionArgument == nil || precisionArgument.LibraryID != m3uint.MYLIBID || precisionArgument.ProgramID != m3uint.M3UINTCONST {
			ev.error = errors.New("wrong argument type")
			return nil
		}
		precision, parseErr := strconv.ParseUint(precisionArgument.ProgramValue, 10, 32)
		if parseErr != nil {
			ev.error = errors.New("convert to integer failed")
			return nil
		}
		if precision > 45 {
			result = input
		} else {
			scale := math.Pow10(int(precision))
			result = math.Round(input*scale) / scale
		}
	case TRUNC:
		result = math.Trunc(input)
	default:
		ev.error = errors.New("unknown ProgramID")
		return nil
	}

	if math.IsNaN(result) {
		ev.error = errors.New("input outside function domain")
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
