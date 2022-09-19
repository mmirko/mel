//go:build !DEBUG
// +build !DEBUG

package lisevo

import (
	"errors"
	"fmt"
	"strconv"

	m3bool "github.com/mmirko/mel/m3bool"
	m3number "github.com/mmirko/mel/m3number"
	m3uint "github.com/mmirko/mel/m3uint"
	m3uintcmp "github.com/mmirko/mel/m3uintcmp"
	mel3program "github.com/mmirko/mel/mel3program"
	statements "github.com/mmirko/mel/statements"
)

type Evaluator struct {
	Impl map[uint16]*mel3program.Mel3Implementation
	Mux  mel3program.Mux
	error
	Result *mel3program.Mel3Program
}

func lisevoMux(v mel3program.Visitor, in_prog *mel3program.Mel3Program) mel3program.Visitor {
	libraryId := in_prog.LibraryID

	switch libraryId {
	case m3uint.MYLIBID:
		newEv := new(m3uint.Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	case m3uintcmp.MYLIBID:
		newEv := new(m3uintcmp.Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	case m3number.MYLIBID:
		newEv := new(m3number.Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	case m3bool.MYLIBID:
		newEv := new(m3bool.Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	case statements.MYLIBID:
		newEv := new(statements.Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	default:
		newEv := new(Evaluator)
		newEv.Impl = v.Get_Implementations()
		newEv.Mux = v.GetMux()
		return newEv
	}
}

func (ev *Evaluator) GetName() string {
	return "lisevo"
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

func (ev *Evaluator) GetResult() *mel3program.Mel3Program {
	return ev.Result
}

func (ev *Evaluator) Visit(in_prog *mel3program.Mel3Program) mel3program.Visitor {

	myMux := ev.GetMux()
	checkEv := myMux(ev, in_prog)

	if ev.GetName() != checkEv.GetName() {
		checkEv.Visit(in_prog)
		if checkEv.GetError() != nil {
			ev.error = checkEv.GetError()
			return nil
		}
		ev.Result = checkEv.GetResult()
		return ev
	}

	programId := in_prog.ProgramID
	libraryId := in_prog.LibraryID

	// DEBUG CODE PLACEHOLDER

	implementation := ev.Impl[libraryId]

	isFunctional := true

	if len(implementation.NonVariadicArgs[programId]) == 0 && !implementation.IsVariadic[programId] {
		isFunctional = false
	}

	if isFunctional {
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("Unkwown LibraryID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
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
