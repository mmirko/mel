package mel3program

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/mmirko/mel/pkg/mel"
)

func CreateGenericCreators(c *mel.MelConfig, ep *mel.EvolutionParameters, impls map[uint16]*Mel3Implementation) map[uint16]Mel3VisitorCreator {
	switch c.VisitorCreatorSet {
	case mel.VISDUMP:
		creators := make(map[uint16]Mel3VisitorCreator)
		for libId, _ := range impls {
			creators[libId] = DumpCreator
		}
		return creators
	case mel.VISBASM:
		creators := make(map[uint16]Mel3VisitorCreator)
		for libId, _ := range impls {
			creators[libId] = BasmCreator
		}
		return creators
	}
	return nil
}

type BasmEvaluator struct {
	*Mel3Object
	error
	Result *Mel3Program
}

func (ev *BasmEvaluator) GetName() string {
	return "dump"
}

func (ev *BasmEvaluator) GetMel3Object() *Mel3Object {
	return ev.Mel3Object
}

func (ev *BasmEvaluator) SetMel3Object(mel3o *Mel3Object) {
	ev.Mel3Object = mel3o
}

func (ev *BasmEvaluator) GetError() error {
	return ev.error
}

func (ev *BasmEvaluator) GetResult() *Mel3Program {
	return ev.Result
}

func (ev *BasmEvaluator) Visit(in_prog *Mel3Program) Mel3Visitor {

	debug := ev.Config.Debug

	if debug {
		fmt.Println("basm: Visit: ", in_prog)
	}

	checkEv := ProgMux(ev, in_prog)

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
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("unknown LibraryID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
			return nil
		}
	} else {
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("unknown LibraryID")
			return nil
		}
	}
}

func (ev *BasmEvaluator) Inspect() string {
	obj := ev.GetMel3Object()
	implementations := obj.Implementation
	if ev.error == nil {
		if dump, err := ProgDump(implementations, ev.Result); err == nil {
			return "Evaluation ok: " + dump
		} else {
			return "Result export failed:" + fmt.Sprint(err)
		}
	} else {
		return fmt.Sprint(ev.error)
	}
}

type DumpEvaluator struct {
	*Mel3Object
	error
	Result *Mel3Program
}

func (ev *DumpEvaluator) GetName() string {
	return "basm"
}

func (ev *DumpEvaluator) GetMel3Object() *Mel3Object {
	return ev.Mel3Object
}

func (ev *DumpEvaluator) SetMel3Object(mel3o *Mel3Object) {
	ev.Mel3Object = mel3o
}

func (ev *DumpEvaluator) GetError() error {
	return ev.error
}

func (ev *DumpEvaluator) GetResult() *Mel3Program {
	return ev.Result
}

func (ev *DumpEvaluator) Visit(in_prog *Mel3Program) Mel3Visitor {

	debug := ev.Config.Debug

	if debug {
		fmt.Println("dump: Visit: ", in_prog)
	}

	checkEv := ProgMux(ev, in_prog)

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
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("unknown LibraryID on " + strconv.Itoa(int(libraryId)) + ":" + strconv.Itoa(int(programId)))
			return nil
		}
	} else {
		switch in_prog.LibraryID {
		default:
			ev.error = errors.New("unknown LibraryID")
			return nil
		}
	}
}

func (ev *DumpEvaluator) Inspect() string {
	obj := ev.GetMel3Object()
	implementations := obj.Implementation
	if ev.error == nil {
		if dump, err := ProgDump(implementations, ev.Result); err == nil {
			return "Evaluation ok: " + dump
		} else {
			return "Result export failed:" + fmt.Sprint(err)
		}
	} else {
		return fmt.Sprint(ev.error)
	}
}

func BasmCreator() Mel3Visitor {
	return new(BasmEvaluator)
}

func DumpCreator() Mel3Visitor {
	return new(DumpEvaluator)
}
