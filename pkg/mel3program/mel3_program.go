package mel3program

import (
	"fmt"

	"github.com/mmirko/mel/pkg/mel"
)

const (
	BUILTINS = uint16(0) + iota
	LIB_M3STATEMENTS
	LIB_M3SYMBOLIC
	LIB_M3BOOL
	LIB_M3BOOLCMP
	LIB_M3UINT
	LIB_M3DATES
	LIB_M3REDMON
	LIB_M3UINTCMP
	LIB_M3NUMBER
	LIB_LISEVO
	LIB_ENVFLOAT
	LIB_M3MELBOND
)

type ArgType struct {
	LibraryID  uint16
	TypeID     uint16   // Return values are multitype
	TensorRapr []uint64 // 0 means that the dimension may be anything, !=0 the dimension is fixed (for example a 3 element vector is []uint64{3}, a generic vector is []uint64{0}
}

type ArgumentsTypes []ArgType

// This is the language implementation.
type Mel3Implementation struct {
	ProgramNames    map[uint16]string
	TypeNames       map[uint16]string
	ProgramTypes    map[uint16]ArgumentsTypes
	NonVariadicArgs map[uint16]ArgumentsTypes // Non variadic arguments
	IsVariadic      map[uint16]bool
	VariadicType    map[uint16]ArgType // The type of the variadic argument (eventually)
	ImplName        string
	Signatures      map[uint16]string
}

type Mel3Program struct {
	LibraryID    uint16
	ProgramID    uint16
	NextPrograms []*Mel3Program
	ProgramValue string
}

type Mel3Visitor interface {
	GetName() string
	GetMel3Object() *Mel3Object
	SetMel3Object(*Mel3Object)
	Visit(*Mel3Program) Mel3Visitor
	GetError() error
	GetResult() *Mel3Program
	Inspect() string
}

type Mel3VisitorCreator func() Mel3Visitor
type Mel3Object struct {
	StartProgram   *Mel3Program
	Config         *mel.MelConfig
	Implementation map[uint16]*Mel3Implementation
	VisitorCreator map[uint16]Mel3VisitorCreator
	DefaultCreator Mel3VisitorCreator
	Result         *Mel3Program
	Environment    *interface{}
}

func (a ArgType) String(impl *Mel3Implementation) string {
	return impl.TypeNames[a.TypeID]
}

func (as ArgumentsTypes) String(impl *Mel3Implementation) string {
	result := ""
	for i, arg := range as {
		if i != 0 {
			result += ","
		}
		result += arg.String(impl)
	}
	return result
}

func SameType(t1 ArgType, t2 ArgType) bool {
	switch {
	case t1.LibraryID != t2.LibraryID:
		return false
	case t1.TypeID != t2.TypeID:
		return false
	case len(t1.TensorRapr) != len(t2.TensorRapr):
		return false
	default:
		// TODO Not necessarily correct, check it when the time comes
		for i, j := range t1.TensorRapr {
			if j != t2.TensorRapr[i] {
				return false
			}
		}
	}
	return true
}

func (mo *Mel3Object) Compute() error {
	prog := mo.StartProgram
	v := mo.VisitorCreator[prog.LibraryID]
	ev := v()
	ev.SetMel3Object(mo)
	Walk(ev, prog)
	if err := ev.GetError(); err != nil {
		return err
	} else {
		mo.Result = ev.GetResult()
		return nil
	}
}

func (mo *Mel3Object) Inspect() string {
	result, _ := ProgDump(mo.Implementation, mo.Result)
	return result
}

type Mux func(Mel3Visitor, *Mel3Program) Mel3Visitor

// Walking the program tree
func Walk(v Mel3Visitor, in_prog *Mel3Program) {
	obj := v.GetMel3Object()

	if obj.Config.Debug {
		fmt.Printf("walk enter\n")
		defer fmt.Printf("walk exit\n")
	}

	implementations := obj.Implementation
	programID := in_prog.ProgramID
	libraryID := in_prog.LibraryID

	implementation := implementations[libraryID]

	if v = v.Visit(in_prog); v == nil {
		return
	}

	isFunctional := true

	if len(implementation.NonVariadicArgs[programID]) == 0 && !implementation.IsVariadic[programID] {
		isFunctional = false
	}

	if isFunctional {
		for _, nextProg := range in_prog.NextPrograms {
			evaluator := ProgMux(v, nextProg)
			evaluator.Visit(nextProg)
		}
	}
}

// Mux different libraries
func ProgMux(v Mel3Visitor, in_prog *Mel3Program) Mel3Visitor {

	libId := in_prog.LibraryID
	obj := v.GetMel3Object()

	if creator, ok := obj.VisitorCreator[libId]; ok {
		c := creator()
		c.SetMel3Object(obj)
		return c
	} else {
		return v
	}
}
