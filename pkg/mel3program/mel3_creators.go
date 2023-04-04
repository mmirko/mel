package mel3program

import (
	"fmt"

	"github.com/mmirko/mel/pkg/mel"
)

func CreateGenericCreators(o *Mel3Object, ep *mel.EvolutionParameters, impls map[uint16]*Mel3Implementation) map[uint16]Mel3VisitorCreator {
	if o.DefaultCreator != nil {
		creators := make(map[uint16]Mel3VisitorCreator)
		creators[BUILTINS] = o.DefaultCreator
		for libId, _ := range impls {
			creators[libId] = o.DefaultCreator
		}
		return creators
	}
	return nil
}

type DumpEvaluator struct {
	*Mel3Object
	error
	Result *Mel3Program
	level  int
}

func (ev *DumpEvaluator) GetName() string {
	return "dump"
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

	for i := 0; i < ev.level; i++ {
		fmt.Print(" ")
	}
	fmt.Println(in_prog)

	ev.level = ev.level + 1

	arg_num := len(in_prog.NextPrograms)
	evaluators := make([]Mel3Visitor, arg_num)
	for i, prog := range in_prog.NextPrograms {
		evaluators[i] = ev
		evaluators[i].Visit(prog)
	}

	return nil
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

func DumpCreator() Mel3Visitor {
	return new(DumpEvaluator)
}
