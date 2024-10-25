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

	envI := ev.Environment
	env := envI.(*EnvFloat)
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
			case READINPUT, READOUTPUT, READKEEP:
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
						var resultN float32
						switch in_prog.ProgramID {
						case READINPUT:
							if val, err := env.ReadInput(value0n); err == nil {
								resultN = val
							} else {
								ev.error = err
								return nil
							}
						case READOUTPUT:
							if val, err := env.ReadOutput(value0n); err == nil {
								resultN = val
							} else {
								ev.error = err
								return nil
							}
						case READKEEP:
							if val, err := env.ReadKeep(value0n); err == nil {
								resultN = val
							} else {
								ev.error = err
								return nil
							}
						}

						resultS := strconv.FormatFloat(float64(resultN), 'f', -1, 32)
						result := new(mel3program.Mel3Program)
						result.LibraryID = m3number.MYLIBID
						result.ProgramID = m3number.M3NUMBERCONST
						result.ProgramValue = resultS
						result.NextPrograms = nil
						ev.Result = result
						return nil
					} else {
						ev.error = errors.New("convert to integer failed")
						return nil
					}
				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}
			case WRITEOUTPUT, WRITEKEEP:

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

					value0n := 0
					if res, err := strconv.Atoi(value0); err == nil {
						value0n = res
					} else {
						ev.error = errors.New("convert to integer failed")
						return nil
					}

					value1n := float32(0.0)
					if res, err := strconv.ParseFloat(value1, 32); err == nil {
						value1n = float32(res)
					} else {
						ev.error = errors.New("convert to float failed")
						return nil
					}

					switch in_prog.ProgramID {
					case WRITEOUTPUT:
						if err := env.WriteOutput(value0n, value1n); err != nil {
							ev.error = err
							return nil
						}
					case WRITEKEEP:
						if err := env.WriteKeep(value0n, value1n); err != nil {
							ev.error = err
							return nil
						}
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = m3number.MYLIBID
					result.ProgramID = m3number.M3NUMBERCONST
					result.ProgramValue = value1
					result.NextPrograms = nil
					ev.Result = result
					return nil

				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}
			case PUSHKEEP:
				if arg_num == 1 {
					res0 := evaluators[0].GetResult()

					value0 := ""
					if res0 != nil && res0.LibraryID == m3number.MYLIBID && res0.ProgramID == m3number.M3NUMBERCONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					value0n := float32(0.0)
					if res, err := strconv.ParseFloat(value0, 32); err == nil {
						value0n = float32(res)
					} else {
						ev.error = errors.New("convert to float failed")
						return nil
					}

					if err := env.PushKeep(value0n); err != nil {
						ev.error = err
						return nil
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = m3number.MYLIBID
					result.ProgramID = m3number.M3NUMBERCONST
					result.ProgramValue = value0
					result.NextPrograms = nil
					ev.Result = result
					return nil

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
		case MYLIBID:
			switch in_prog.ProgramID {
			case POPKEEP:
				if val, err := env.PopKeep(); err == nil {
					result := new(mel3program.Mel3Program)
					result.LibraryID = m3number.MYLIBID
					result.ProgramID = m3number.M3NUMBERCONST
					result.ProgramValue = strconv.FormatFloat(float64(val), 'f', -1, 32)
					result.NextPrograms = nil
					ev.Result = result
					return nil
				} else {
					ev.error = err
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
