//go:build !DEBUG
// +build !DEBUG

package m3uintcmp

import (
	"errors"
	"fmt"
	"strconv"

	m3bool "github.com/mmirko/mel/pkg/m3bool"
	m3uint "github.com/mmirko/mel/pkg/m3uint"
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
	return "m3uintcmp"
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
		fmt.Println("m3uintcmp enter:", in_prog)
		defer fmt.Println("m3uintcmp exit")
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
			case EQ, NE, LT, LE, GT, GE:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()
					value0 := ""
					if res0 != nil && res0.LibraryID == m3uint.MYLIBID && res0.ProgramID == m3uint.M3UINTCONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == m3uint.MYLIBID && res1.ProgramID == m3uint.M3UINTCONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					var opResult bool
					var opResults string

					if value0n, err := strconv.Atoi(value0); err == nil {
						if value1n, err := strconv.Atoi(value1); err == nil {
							if value0n < 0 || value1n < 0 {
								ev.error = errors.New("convert to integer failed")
								return nil
							}

							switch in_prog.ProgramID {
							case EQ:
								opResult = value0n == value1n
							case NE:
								opResult = value0n != value1n
							case LT:
								opResult = value0n < value1n
							case LE:
								opResult = value0n <= value1n
							case GT:
								opResult = value0n > value1n
							case GE:
								opResult = value0n >= value1n
							}
						} else {
							ev.error = errors.New("convert to integer failed")
							return nil
						}

					} else {
						ev.error = errors.New("convert to integer failed")
						return nil
					}

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
