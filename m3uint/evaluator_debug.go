//go:build DEBUG
// +build DEBUG

package m3uint

import (
	"errors"
	"fmt"
	mel3program "github.com/mmirko/mel/mel3program"
	"strconv"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3_implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3_program
}

func M3uintmux(v mel3program.Visitor, in_prog *mel3program.Mel3_program) mel3program.Visitor {
	result := new(Evaluator)
	result.Impl = v.Get_Implementations()
	result.Mux = v.GetMux()
	return result
}

func (ev *Evaluator) Get_Implementations() map[uint16]*mel3program.Mel3_implementation {
	return ev.Impl
}

func (ev *Evaluator) GetName() string {
	return "m3uint"
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

func (ev *Evaluator) GetResult() *mel3program.Mel3_program {
	return ev.Result
}

func (ev *Evaluator) Visit(in_prog *mel3program.Mel3_program) mel3program.Visitor {

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
			case ADD, SUB, MULT, DIV:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()
					fmt.Println(res0, res1)
					value0 := ""
					if res0 != nil && res0.LibraryID == libraryid && res0.ProgramID == M3UINTCONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == libraryid && res1.ProgramID == M3UINTCONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument type")
						return nil
					}

					op_result := ""

					if value0n, err := strconv.Atoi(value0); err == nil {
						if value1n, err := strconv.Atoi(value1); err == nil {
							if value0n < 0 || value1n < 0 {
								ev.error = errors.New("Convert to integer failed")
								return nil
							}

							var op_resultn uint

							switch in_prog.ProgramID {
							case ADD:
								op_resultn = uint(value0n + value1n)
							case SUB:
								// TODO Check for the ovwrflow
								op_resultn = uint(value0n - value1n)
							case MULT:
								op_resultn = uint(value0n * value1n)
							case DIV:
								op_resultn = uint(value0n / value1n)
							}

							op_result = strconv.Itoa(int(op_resultn))

						} else {
							ev.error = errors.New("Convert to integer failed")
							return nil
						}
					} else {
						ev.error = errors.New("Convert to integer failed")
						return nil
					}

					result := new(mel3program.Mel3_program)
					result.LibraryID = libraryid
					result.ProgramID = M3UINTCONST
					result.ProgramValue = op_result
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
			case M3UINTCONST:
				switch in_prog.ProgramValue {
				default:
					result := new(mel3program.Mel3_program)
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
