//go:build !DEBUG
// +build !DEBUG

package m3bool

import (
	"errors"
	"fmt"

	mel3program "github.com/mmirko/mel/mel3program"
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
	return "m3bool"
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
		fmt.Println("m3bool: Visit: ", in_prog)
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
			case NOT:
				if arg_num == 1 {
					res0 := evaluators[0].GetResult()

					value0 := false
					if res0 != nil && res0.LibraryID == libraryId && res0.ProgramID == CONST {
						if res0.ProgramValue == "true" || res0.ProgramValue == "1" {
							value0 = true
						}
					} else {
						ev.error = errors.New("wrong argument type")
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
					result.LibraryID = libraryId
					result.ProgramID = CONST
					result.ProgramValue = op_results
					result.NextPrograms = nil
					ev.Result = result
					return nil
				} else {
					ev.error = errors.New("wrong argument number")
					return nil
				}

			case AND, OR, XOR, NAND, NOR, XNOR:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()
					value0 := false
					if res0 != nil && res0.LibraryID == libraryId && res0.ProgramID == CONST {
						if res0.ProgramValue == "true" || res0.ProgramValue == "1" {
							value0 = true
						}
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}

					value1 := false
					if res1 != nil && res1.LibraryID == libraryId && res1.ProgramID == CONST {
						if res1.ProgramValue == "true" || res1.ProgramValue == "1" {
							value1 = true
						}
					} else {
						ev.error = errors.New("wrong argument type")
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
					result.LibraryID = libraryId
					result.ProgramID = CONST
					result.ProgramValue = op_results
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
		case MYLIBID:
			switch in_prog.ProgramID {
			case CONST:
				switch in_prog.ProgramValue {
				default:
					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryId
					result.ProgramID = programId
					result.ProgramValue = in_prog.ProgramValue
					result.NextPrograms = nil
					ev.Result = result
					return nil
				}
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
