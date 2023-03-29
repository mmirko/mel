//go:build !DEBUG
// +build !DEBUG

package m3boolcmp

import (
	"errors"
	"fmt"

	m3bool "github.com/mmirko/mel/pkg/m3bool"
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
	return "m3boolcmp"
}

func (ev *Evaluator) GetMel3Object() *mel3program.Mel3Object {
	return ev.Mel3Object
}

func (ev *Evaluator) SetMel3Object(mel3o *mel3program.Mel3Object) {
	ev.Mel3Object = mel3o
}

func (ev *Evaluator) GetError() error {
	return ev.error
}

func (ev *Evaluator) GetResult() *mel3program.Mel3Program {
	return ev.Result
}

func (ev *Evaluator) Visit(in_prog *mel3program.Mel3Program) mel3program.Mel3Visitor {

	debug := ev.Config.Debug

	if debug {
		fmt.Println("m3boolcmp enter:", in_prog)
		defer fmt.Println("m3boolcmp exit")
	}

	checkEv := mel3program.ProgMux(ev, in_prog)

	if ev.GetName() != checkEv.GetName() {
		return checkEv.Visit(in_prog)
	}

	obj := ev.GetMel3Object()
	implementations := obj.Implementation

	programId := in_prog.ProgramID
	libraryId := in_prog.LibraryID

	implementation := implementations[libraryId]

	isFunctional := true

	if len(implementation.NonVariadicArgs[programId]) == 0 && !implementation.IsVariadic[programId] {
		isFunctional = false
	}

	if isFunctional {
		arg_num := len(in_prog.NextPrograms)
		evaluators := make([]mel3program.Mel3Visitor, arg_num)
		for i, prog := range in_prog.NextPrograms {
			evaluators[i] = mel3program.ProgMux(ev, prog)
			evaluators[i].Visit(prog)
		}

		switch in_prog.LibraryID {
		case MYLIBID:
			switch in_prog.ProgramID {
			case EQ, NE:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()
					value0 := ""
					if res0 != nil && res0.LibraryID == m3bool.MYLIBID && res0.ProgramID == m3bool.CONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == m3bool.MYLIBID && res1.ProgramID == m3bool.CONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					var opResult bool
					var opResults string

					opResult = value0 == value1

					if opResult {
						opResults = "true"
					} else {
						opResults = "false"
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = m3bool.MYLIBID
					result.ProgramID = m3bool.CONST
					result.ProgramValue = opResults
					result.NextPrograms = nil
					ev.Result = result
					return nil
				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}
			}
		default:
			ev.error = errors.New("unknown LibraryID")
			return nil
		}
	} else {

		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("unknown LibraryID")
			return nil
		}
	}

	return ev
}

func (ev *Evaluator) Inspect() string {
	obj := ev.GetMel3Object()
	implementations := obj.Implementation
	if ev.error == nil {
		if dump, err := mel3program.ProgDump(implementations, ev.Result); err == nil {
			return "Evaluation ok: " + dump
		} else {
			return "Result export failed:" + fmt.Sprint(err)
		}
	} else {
		return fmt.Sprint(ev.error)
	}
}
