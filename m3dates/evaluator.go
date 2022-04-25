//go:build !DEBUG
// +build !DEBUG

package m3dates

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	m3uint "github.com/mmirko/mel/m3uint"
	mel3program "github.com/mmirko/mel/mel3program"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3Implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3_program
}

func M3datesmux(v mel3program.Visitor, in_prog *mel3program.Mel3_program) mel3program.Visitor {
	libraryid := in_prog.LibraryID

	if libraryid == m3uint.MYLIBID {
		newev := new(m3uint.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	result := new(Evaluator)
	result.Impl = v.Get_Implementations()
	result.Mux = v.GetMux()
	return result
}

func (ev *Evaluator) GetName() string {
	return "m3dates"
}

func (ev *Evaluator) Get_Implementations() map[uint16]*mel3program.Mel3Implementation {
	return ev.Impl
}

func (ev *Evaluator) GetMux() mel3program.Mux {
	return ev.Mux
}

func (ev *Evaluator) GetError() error {
	return ev.error
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
		checkev.Visit(in_prog)
		if checkev.GetError() != nil {
			ev.error = checkev.GetError()
			return nil
		}
		ev.Result = checkev.GetResult()
		return ev
	}

	programid := in_prog.ProgramID
	libraryid := in_prog.LibraryID

	// DEBUG CODE PLACEHOLDER

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
			case DIFFDAYS:
				if arg_num == 2 {
					res0 := evaluators[0].GetResult()
					res1 := evaluators[1].GetResult()

					value0 := ""
					if res0 != nil && res0.LibraryID == libraryid && res0.ProgramID == DATECONST {
						value0 = res0.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument 0 type on " + strconv.Itoa(int(libraryid)) + ":" + strconv.Itoa(int(programid)))
						return nil
					}

					value1 := ""
					if res1 != nil && res1.LibraryID == libraryid && res1.ProgramID == DATECONST {
						value1 = res1.ProgramValue
					} else {
						ev.error = errors.New("Wrong argument 1 type on " + strconv.Itoa(int(libraryid)) + ":" + strconv.Itoa(int(programid)))
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

					result := new(mel3program.Mel3_program)
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
			ev.error = errors.New("Unkwown LibraryID on " + strconv.Itoa(int(libraryid)) + ":" + strconv.Itoa(int(programid)))
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
					result := new(mel3program.Mel3_program)
					result.LibraryID = libraryid
					result.ProgramID = programid
					result.ProgramValue = time.Now().Format(baselayout)
					result.NextPrograms = nil
					ev.Result = result
				default:
					layouts := []string{"2006-01-02", "2006-01-02 15:04:05"}
					oneok := false
					for _, layout := range layouts {
						if t, err := time.Parse(layout, in_prog.ProgramValue); err == nil {
							result := new(mel3program.Mel3_program)
							result.LibraryID = libraryid
							result.ProgramID = programid
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
			ev.error = errors.New("Unkwown LibraryID")
			return nil
		}
	}

	return ev
}

func (ev *Evaluator) Inspect() string {
	if ev.error == nil {
		if ev.Result != nil {
			fmt.Println(ev.Result)
			if dump, err := mel3program.ProgDump(ev.Impl, ev.Result); err == nil {
				return "Evaluation ok: " + dump
			} else {
				return "Result export failed:" + fmt.Sprint(err)
			}
		} else {
			return "Result export failed"
		}
	} else {
		return fmt.Sprint(ev.error)
	}
}
