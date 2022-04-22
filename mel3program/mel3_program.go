package mel3program

import ()

// ********** Data structures

const (
	BUILTINS = uint16(0) + iota
	LIB_STATEMENTS
	LIB_LISEVO
	LIB_SYMBOLIC_MATH3
	LIB_GROWSTOCK
	LIB_BOOL3
	LIB_M3UINT
	LIB_DATES
	LIB_M3REDMON
	LIB_M3UINTCMP
	LIB_M3BAENGINE
	LIB_MELBINANCE
	LIB_M3NUMBER
)

type ArgType struct {
	LibraryID  uint16
	TypeID     uint16   // Return values are multitype
	TensorRapr []uint64 // 0 means that the dimension may be anything, !=0 the dimension is fixed (for example a 3 element vector is []uint64{3}, a generic vector is []uint64{0}
}

type ArgumentsTypes []ArgType

// This is the language implementation.
type Mel3_implementation struct {
	ProgramNames    map[uint16]string
	TypeNames       map[uint16]string
	ProgramTypes    map[uint16]ArgumentsTypes
	NonVariadicArgs map[uint16]ArgumentsTypes // Non variadic arguments
	IsVariadic      map[uint16]bool
	VariadicType    map[uint16]ArgType // The type of the variadic argument (eventyally)
	Implname        string
	Signatures      map[uint16]string
}

type Mel3_program struct {
	LibraryID    uint16
	ProgramID    uint16
	NextPrograms []*Mel3_program
	ProgramValue string
}

type Mel3_object struct {
	StartProgram   *Mel3_program
	Implementation map[uint16]*Mel3_implementation
	Environment    interface{}
}

func (a ArgType) String(impl *Mel3_implementation) string {
	return impl.TypeNames[a.TypeID]
}

func (as ArgumentsTypes) String(impl *Mel3_implementation) string {
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
		// TODO Not necessarly correct, check it when the time comes
		for i, j := range t1.TensorRapr {
			if j != t2.TensorRapr[i] {
				return false
			}
		}
	}
	return true
}

type VisitFunction func(Visitor, *Mel3_program) Visitor

type Visitor interface {
	GetName() string
	Visit(*Mel3_program) Visitor
	Get_Implementations() map[uint16]*Mel3_implementation
	GetMux() Mux
	SetMux(Mux)
	GetError() error
	GetResult() *Mel3_program
	Inspect() string
}

type Mux func(Visitor, *Mel3_program) Visitor

// Walk
func Walk(v Visitor, in_prog *Mel3_program) {
	implementations := v.Get_Implementations()
	programid := in_prog.ProgramID
	libraryid := in_prog.LibraryID

	implementation := implementations[libraryid]
	mymux := v.GetMux()

	if v = v.Visit(in_prog); v == nil {
		return
	}

	isfunctional := true

	if len(implementation.NonVariadicArgs[programid]) == 0 && !implementation.IsVariadic[programid] {
		isfunctional = false
	}

	if isfunctional {
		for _, nextprog := range in_prog.NextPrograms {
			evaluator := mymux(v, nextprog)
			evaluator.Visit(nextprog)
		}
	}
}
