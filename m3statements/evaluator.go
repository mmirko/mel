//go:build !DEBUG
// +build !DEBUG

package m3statements

import (
	"errors"
	"fmt"
	"strconv"

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
	return "m3statements"
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
		fmt.Println("m3statements: Visit: ", in_prog)
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
			if err := evaluators[i].GetError(); err != nil {
				ev.error = err
				return ev
			}
		}

		switch in_prog.LibraryID {
		case MYLIBID:
			switch in_prog.ProgramID {
			case MULTISTMT:

				nextPrograms := make([]*mel3program.Mel3Program, 0)

				for i := 0; i < arg_num; i++ {
					res := evaluators[i].GetResult()
					if res != nil && implementations[res.LibraryID].ProgramTypes[res.ProgramID][0].LibraryID == MYLIBID && implementations[res.LibraryID].ProgramTypes[res.ProgramID][0].TypeID == MULTISTMT {
						nextPrograms = append(nextPrograms, res)
					} else {
						ev.error = errors.New("wrong argument type")
						return nil
					}
				}
				result := new(mel3program.Mel3Program)
				result.LibraryID = libraryId
				result.ProgramID = MULTISTMT
				result.ProgramValue = ""
				result.NextPrograms = nextPrograms
				ev.Result = result
				return nil
			}
		default:
			ev.error = errors.New("unknown library")
			return nil
		}

	} else {

		switch in_prog.LibraryID {
		case MYLIBID:
			switch in_prog.ProgramID {
			case NOP:
				switch in_prog.ProgramValue {
				default:
					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryId
					result.ProgramID = NOP
					result.ProgramValue = in_prog.ProgramValue
					result.NextPrograms = nil
					ev.Result = result
					return nil
				}
			default:
				ev.error = errors.New("unknown ProgramID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
				return nil
			}
		default:
			ev.error = errors.New("unknown LibraryID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
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
