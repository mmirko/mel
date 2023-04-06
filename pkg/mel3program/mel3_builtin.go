package mel3program

import (
	"fmt"
)

const (
	B_IN_INPUT = uint16(0) + iota
	B_IN_OUTPUT
	B_IN_GROUP
	B_IN_UNGROUP
)

func isBuiltin(programName string) bool {
	switch programName {
	case "input", "output", "group", "ungroup":
		return true
	}
	return false
}

func processBuiltin(implementation map[uint16]*Mel3Implementation, programName string, args []string) (*Mel3Program, *ArgumentsTypes, error) {

	// All the built-in programs are defined here. They process some number of arguments and use the rest as their parameters.
	// Their "value" parameter is a reference to their structures.

	switch programName {
	case "output", "group", "ungroup", "input":

		// Unary built-ins
		var result Mel3Program
		result.NextPrograms = make([]*Mel3Program, 1)
		argList := ArgumentsTypes{}

		// The last argument is the value of the program
		nextProg := args[len(args)-1]

		if tempProgr, tempType, err := import_engine(implementation, nextProg); err != nil {
			return nil, nil, err
		} else {
			result.NextPrograms[0] = tempProgr

			// Composition of the argument list
			for _, itype := range *tempType {
				argList = append(argList, itype)
			}
		}

		result.LibraryID = BUILTINS
		switch programName {
		case "group":
			// Group can have an argument to specify the group name
			if len(args) > 2 {
				return nil, nil, fmt.Errorf("Too many arguments for group")
			} else if len(args) == 2 {
				result.ProgramValue = args[0]
			} else {
				result.ProgramValue = ""
			}
			result.ProgramID = B_IN_GROUP
		case "ungroup":
			if len(args) > 1 {
				return nil, nil, fmt.Errorf("Too many arguments for ungroup")
			}
			result.ProgramID = B_IN_UNGROUP
			result.ProgramValue = ""
		case "output":
			if len(args) > 2 {
				return nil, nil, fmt.Errorf("Too many arguments for output")
			} else if len(args) == 2 {
				result.ProgramValue = args[0]
			} else {
				result.ProgramValue = ""
			}
			result.ProgramID = B_IN_OUTPUT
		case "input":
			if len(args) > 2 {
				return nil, nil, fmt.Errorf("Too many arguments for input")
			} else if len(args) == 2 {
				result.ProgramValue = args[0]
			} else {
				result.ProgramValue = ""
			}
			result.ProgramID = B_IN_INPUT
		}

		return &result, &argList, nil
	}
	return nil, nil, fmt.Errorf("Unknown builtin: %s", programName)
}

type BuiltInEvaluator struct {
	*Mel3Object
	error
	Result *Mel3Program
}

func (ev *BuiltInEvaluator) GetName() string {
	return "builtin"
}

func (ev *BuiltInEvaluator) GetMel3Object() *Mel3Object {
	return ev.Mel3Object
}

func (ev *BuiltInEvaluator) SetMel3Object(mel3o *Mel3Object) {
	ev.Mel3Object = mel3o
}

func (ev *BuiltInEvaluator) GetError() error {
	return ev.error
}

func (ev *BuiltInEvaluator) GetResult() *Mel3Program {
	return ev.Result
}

func (ev *BuiltInEvaluator) Visit(in_prog *Mel3Program) Mel3Visitor {

	debug := ev.Config.Debug

	if debug {
		fmt.Println("builtin: Visit: ", in_prog)
	}

	checkEv := ProgMux(ev, in_prog)

	if ev.GetName() != checkEv.GetName() {
		return checkEv.Visit(in_prog)
	}

	switch in_prog.ProgramID {
	case B_IN_GROUP, B_IN_UNGROUP, B_IN_OUTPUT, B_IN_INPUT:
		arg_num := 1
		evaluators := make([]Mel3Visitor, arg_num)
		for i, prog := range in_prog.NextPrograms {
			evaluators[i] = ProgMux(ev, prog)
			evaluators[i].Visit(prog)
		}
		res := evaluators[0].GetResult()

		result := new(Mel3Program)
		result.LibraryID = res.LibraryID
		result.ProgramID = res.ProgramID
		result.ProgramValue = res.ProgramValue
		result.NextPrograms = res.NextPrograms
		ev.Result = result
	}
	return nil
}

func (ev *BuiltInEvaluator) Inspect() string {
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
func BuiltInCreator() Mel3Visitor {
	return new(BuiltInEvaluator)
}
