//go:build !DEBUG
// +build !DEBUG

package m3dates

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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
	return "m3dates"
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
		fmt.Println("m3dates: Visit: ", in_prog)
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
			case DIFFDAYS:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()

					value0 := ""
					if res0 != nil && res0.LibraryID == libraryId && res0.ProgramID == DATECONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument 0 type on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == libraryId && res1.ProgramID == DATECONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument 1 type on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
						return nil
					}

					layouts := []string{"2006-01-02", "2006-01-02 15:04:05"}

					oneok := false
					var t0 time.Time
					for _, layout := range layouts {
						if t, err := time.Parse(layout, value0); err == nil {
							oneok = true
							t0 = t
							break
						}
					}

					if !oneok {
						ev.error = errors.New("Date parse error")
						return nil
					}

					oneok = false
					var t1 time.Time
					for _, layout := range layouts {
						if t, err := time.Parse(layout, value1); err == nil {
							oneok = true
							t1 = t
							break
						}
					}

					if !oneok {
						ev.error = errors.New("Date parse error")
						return nil
					}

					op_resultn := 0
					if t0.Before(t1) {
						op_resultn = int(t1.Sub(t0).Hours() / 24)
					} else {
						op_resultn = int(t0.Sub(t1).Hours() / 24)
					}

					op_result := strconv.Itoa(op_resultn)

					result := new(mel3program.Mel3Program)
					result.LibraryID = m3uint.MYLIBID
					result.ProgramID = m3uint.M3UINTCONST
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
			ev.error = errors.New("Unkwown LibraryID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
			return nil
		}
	} else {

		switch in_prog.LibraryID {
		case MYLIBID:
			switch in_prog.ProgramID {
			case DATECONST:
				baselayout := "2006-01-02"
				switch in_prog.ProgramValue {
				case "today":
					result := new(mel3program.Mel3Program)
					result.LibraryID = libraryId
					result.ProgramID = programId
					result.ProgramValue = time.Now().Format(baselayout)
					result.NextPrograms = nil
					ev.Result = result
				default:
					layouts := []string{"2006-01-02", "2006-01-02 15:04:05"}
					oneok := false
					for _, layout := range layouts {
						if t, err := time.Parse(layout, in_prog.ProgramValue); err == nil {
							result := new(mel3program.Mel3Program)
							result.LibraryID = libraryId
							result.ProgramID = programId
							result.ProgramValue = t.Format(baselayout)
							result.NextPrograms = nil
							ev.Result = result
							oneok = true
							break
						}
					}
					if !oneok {
						ev.error = errors.New("Date parse failed")
						return ev
					}
				}
				// TODO
			case TIMESTAMPCONST:
				switch in_prog.ProgramValue {
				case "now":
					ev.Result = nil
				default:
					layout := "2006-01-02 15:04:05"
					if _, err := time.Parse(layout, in_prog.ProgramValue); err == nil {
					} else {
						ev.error = errors.New("Date parse failed")
						return ev
					}
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
