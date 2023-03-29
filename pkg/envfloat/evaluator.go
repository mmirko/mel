//go:build !DEBUG
// +build !DEBUG

package envfloat

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mmirko/mel/pkg/m3number"
	"github.com/mmirko/mel/pkg/m3uint"
	"github.com/mmirko/mel/pkg/mel3program"
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
	return "envfloat"
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
		fmt.Println("envfloat: Visit: ", in_prog)
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

		envI := *ev.Environment
		env := envI.(*EnvFloat)

		switch in_prog.LibraryID {
		case MYLIBID:
			switch in_prog.ProgramID {
			case READINPUT:
				if arg_num == 1 {
					res0 := evaluators[0].GetResult()
					value0 := ""
					if res0 != nil && res0.LibraryID == m3uint.MYLIBID && res0.ProgramID == m3uint.M3UINTCONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					if value0n, err := strconv.Atoi(value0); err == nil {
						if value0n < len(env.inVars) {
							resultN := env.inVars[value0n]
							resultS := strconv.FormatFloat(float64(resultN), 'f', -1, 32)
							result := new(mel3program.Mel3Program)
							result.LibraryID = m3number.MYLIBID
							result.ProgramID = m3number.M3NUMBERCONST
							result.ProgramValue = resultS
							result.NextPrograms = nil
							ev.Result = result
							return nil
						} else {
							ev.error = errors.New("wrong argument value")
							return nil
						}
					} else {
						ev.error = errors.New("convert to integer failed")
						return nil
					}
				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}
			case WRITEOUTPUT:

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
					if res1 != nil && res1.LibraryID == m3number.MYLIBID && res1.ProgramID == m3number.M3NUMBERCONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					if value0n, err := strconv.Atoi(value0); err == nil {
						if value0n < len(env.outVars) {
							if value1n, err := strconv.ParseFloat(value1, 32); err == nil {
								env.outVars[value0n] = float32(value1n)
								result := new(mel3program.Mel3Program)
								result.LibraryID = m3number.MYLIBID
								result.ProgramID = m3number.M3NUMBERCONST
								result.ProgramValue = value1
								result.NextPrograms = nil
								ev.Result = result
								return nil
							} else {
								ev.error = errors.New("convert to float failed")
								return nil
							}
						} else {
							ev.error = errors.New("wrong argument value")
							return nil
						}
					} else {
						ev.error = errors.New("convert to integer failed")
						return nil
					}
				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}
			default:
				ev.error = errors.New("unknown ProgramID")
				return nil
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
