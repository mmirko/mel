//go:build !DEBUG
// +build !DEBUG

package m3redmon

import (
	"errors"
	"fmt"
	"strconv"

	bool3 "github.com/mmirko/mel/bool3"
	dates "github.com/mmirko/mel/dates"
	m3uint "github.com/mmirko/mel/m3uint"
	m3uintcmp "github.com/mmirko/mel/m3uintcmp"
	mel3program "github.com/mmirko/mel/mel3program"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3_implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3_program
}

func M3redmonmux(v mel3program.Visitor, in_prog *mel3program.Mel3_program) mel3program.Visitor {
	libraryid := in_prog.LibraryID

	if libraryid == m3uint.MYLIBID {
		newev := new(m3uint.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	if libraryid == m3uintcmp.MYLIBID {
		newev := new(m3uintcmp.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	if libraryid == dates.MYLIBID {
		newev := new(dates.Evaluator)
		newev.Impl = v.Get_Implementations()
		newev.Mux = v.GetMux()
		return newev
	}

	if libraryid == bool3.MYLIBID {
		newev := new(bool3.Evaluator)
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
	return "m3redmon"
}

func (ev *Evaluator) Get_Implementations() map[uint16]*mel3program.Mel3_implementation {
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
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("Unkwown LibraryID on " + strconv.Itoa(int(libraryid)) + ":" + strconv.Itoa(int(programid)))
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
		if ev.Result != nil {
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
