//go:build !DEBUG
// +build !DEBUG

package m3uintcmp

import (
	"errors"
	"fmt"
	"strconv"

	m3bool "github.com/mmirko/mel/m3bool"
	m3uint "github.com/mmirko/mel/m3uint"
	mel3program "github.com/mmirko/mel/mel3program"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3Implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3Program
}

func M3uintcmpmux(v mel3program.Visitor, in_prog *mel3program.Mel3Program) mel3program.Visitor {
	libraryid := in_prog.LibraryID

	if libraryid == m3uint.MYLIBID {
		newev := new(m3uint.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	if libraryid == m3bool.MYLIBID {
		newev := new(m3bool.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	result := new(Evaluator)
	result.Impl = v.Get_Implementations()
	result.Mux = v.GetMux()
	return result
}

func (ev *Evaluator) Get_Implementations() map[uint16]*mel3program.Mel3Implementation {
	return ev.Impl
}

func (ev *Evaluator) GetName() string {
	return "m3uintcmp"
}

func (ev *Evaluator) GetError() error {
	return ev.error
}

func (ev *Evaluator) GetMux() mel3program.Mux {
	return ev.Mux
}

func (ev *Evaluator) SetMux(in_mux mel3program.Mux) {
	ev.Mux = in_mux
}

func (ev *Evaluator) GetResult() *mel3program.Mel3Program {
	return ev.Result
}

func (ev *Evaluator) Visit(in_prog *mel3program.Mel3Program) mel3program.Visitor {

	mymux := ev.GetMux()
	checkev := mymux(ev, in_prog)

	if ev.GetName() != checkev.GetName() {
		return checkev.Visit(in_prog)
	}

	programid := in_prog.ProgramID
	libraryid := in_prog.LibraryID

	implementation := ev.Impl[libraryid]

	isfunctional := true

	if len(implementation.NonVariadicArgs[programid]) == 0 && !implementation.IsVariadic[programid] {
		isfunctional = false
	}

	if isfunctional {
		arg_num := len(in_prog.NextPrograms)
		evaluators := make([]mel3program.Visitor, arg_num)
		for i, prog := range in_prog.NextPrograms {
			mymux := ev.GetMux()
			evaluators[i] = mymux(ev, prog)
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
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == m3uint.MYLIBID && res1.ProgramID == m3uint.M3UINTCONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					var op_result bool
					var op_results string

					if value0n, err := strconv.Atoi(value0); err == nil {
						if value1n, err := strconv.Atoi(value1); err == nil {
							if value0n < 0 || value1n < 0 {
								ev.error = errors.New("Convert to integer failed")
								return nil
							}

							switch in_prog.ProgramID {
							case EQ:
								op_result = value0n == value1n
							case NE:
								op_result = value0n != value1n
							case LT:
								op_result = value0n < value1n
							case LE:
								op_result = value0n <= value1n
							case GT:
								op_result = value0n > value1n
							case GE:
								op_result = value0n >= value1n
							}
						} else {
							ev.error = errors.New("Convert to integer failed")
							return nil
						}

					} else {
						ev.error = errors.New("Convert to integer failed")
						return nil
					}

					if op_result {
						op_results = "true"
					} else {
						op_results = "false"
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = m3bool.MYLIBID
					result.ProgramID = m3bool.CONST
					result.ProgramValue = op_results
					result.NextPrograms = nil
					ev.Result = result
					return nil
				} else {
					ev.error = errors.New("Wrong argument number")
					return nil
				}
			}
		default:
			ev.error = errors.New("Unkwown LibraryID")
			return nil
		}
	} else {

		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("Unkwown LibraryID")
			return nil
		}
	}

	return ev
}

func (ev *Evaluator) Inspect() string {
	if ev.error == nil {
		if dump, err := mel3program.ProgDump(ev.Impl, ev.Result); err == nil {
			return "Evaluation ok: " + dump
		} else {
			return "Result export failed:" + fmt.Sprint(err)
		}
	} else {
		return fmt.Sprint(ev.error)
	}
}
