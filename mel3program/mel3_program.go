package mel3program

const (
	BUILTINS = uint16(0) + iota
	LIB_STATEMENTS
	LIB_M3SYMBOLIC
	LIB_M3BOOL
	LIB_M3UINT
	LIB_M3DATES
	LIB_M3REDMON
	LIB_M3UINTCMP
	LIB_M3NUMBER
	LIB_RECTANGULAR
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

type Mel3Object struct {
	StartProgram   *Mel3Program
	Implementation map[uint16]*Mel3Implementation
	Environment    interface{}
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

type VisitFunction func(Visitor, *Mel3Program) Visitor

type Visitor interface {
	GetName() string
	Visit(*Mel3Program) Visitor
	Get_Implementations() map[uint16]*Mel3Implementation
	GetMux() Mux
	SetMux(Mux)
	GetError() error
	GetResult() *Mel3Program
	Inspect() string
}

type Mux func(Visitor, *Mel3Program) Visitor

// Walk
func Walk(v Visitor, in_prog *Mel3Program) {
	implementations := v.Get_Implementations()
	programID := in_prog.ProgramID
	libraryID := in_prog.LibraryID

	implementation := implementations[libraryID]
	myMux := v.GetMux()

	if v = v.Visit(in_prog); v == nil {
		return
	}

	isFunctional := true

	if len(implementation.NonVariadicArgs[programID]) == 0 && !implementation.IsVariadic[programID] {
		isFunctional = false
	}

	if isFunctional {
		for _, nextProg := range in_prog.NextPrograms {
			evaluator := myMux(v, nextProg)
			evaluator.Visit(nextProg)
		}
	}
}
