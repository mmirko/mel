//go:build DEBUG
// +build DEBUG

package m3bool

import (
	"errors"
	"fmt"
	mel3program "github.com/mmirko/mel/mel3program"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3Implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3Program
}

func M3boolmux(v mel3program.Visitor, in_prog *mel3program.Mel3Program) mel3program.Visitor {
	result := new(Evaluator)
	result.Impl = v.Get_Implementations()
	result.Mux = v.GetMux()
	return result
}

func (ev *Evaluator) Get_Implementations() map[uint16]*mel3program.Mel3Implementation {
	return ev.Impl
}

func (ev *Evaluator) GetName() string {
	return "m3bool"
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

	fmt.Println("Enter uint ", libraryid, programid)

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
			case NOT:
				if arg_num == 1 {
					res0 := evaluators[0].GetResult()

					value0 := false
					if res0 != nil && res0.LibraryID == libraryid && res0.ProgramID == CONST {
						if res0.ProgramValue == "true" || res0.ProgramValue == "1" {
							value0 = true
						}
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					var op_result bool
					var op_results string

					switch in_prog.ProgramID {
					case NOT:
						op_result = !value0
					}

					if op_result {
						op_results = "true"
					} else {
						op_results = "false"
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryid
					result.ProgramID = CONST
					result.ProgramValue = op_results
					result.NextPrograms = nil
					ev.Result = result
					return nil
				} else {
					ev.error = errors.New("Wrong argument number")
					return nil
				}

			case AND, OR, XOR, NAND, NOR, XNOR:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()
					value0 := false
					if res0 != nil && res0.LibraryID == libraryid && res0.ProgramID == CONST {
						if res0.ProgramValue == "true" || res0.ProgramValue == "1" {
							value0 = true
						}
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					value1 := false
					if res1 != nil && res1.LibraryID == libraryid && res1.ProgramID == CONST {
						if res1.ProgramValue == "true" || res1.ProgramValue == "1" {
							value1 = true
						}
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					var op_result bool
					var op_results string

					switch in_prog.ProgramID {
					case AND:
						op_result = value0 && value1
					case OR:
						op_result = value0 || value1
					case XOR:
						op_result = value0 != value1
					case NAND:
						op_result = !(value0 && value1)
					case NOR:
						op_result = !(value0 || value1)
					case XNOR:
						op_result = value0 == value1
					}

					if op_result {
						op_results = "true"
					} else {
						op_results = "false"
					}

					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryid
					result.ProgramID = CONST
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
		case MYLIBID:
			switch in_prog.ProgramID {
			case CONST:
				switch in_prog.ProgramValue {
				default:
					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryid
					result.ProgramID = programid
					result.ProgramValue = in_prog.ProgramValue
					result.NextPrograms = nil
					ev.Result = result
					return nil
				}
			}
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
