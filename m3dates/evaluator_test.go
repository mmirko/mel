package m3dates

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3datesEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3dates evaluator ----")

	a := new(M3datesMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{`
add(
	diffdays(
		dateconst(2020-05-06),
		dateconst(2020-03-02)
	),
	m3uintconst(1)
)
`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		//		var ev mel3program.Visitor

		//		switch a.StartProgram.LibraryID {
		//		case MYLIBID:
		v := new(Evaluator)
		v.Impl = a.Implementation
		v.Mux = M3datesmux
		//	v.Result = new(mel3program.Mel3_program)
		ev := v
		//		case m3uint.MYLIBID:
		//			v := new(m3uint.Evaluator)
		//			v.Impl = a.Implementation
		//			v.Result = new(mel3program.Mel3_program)
		//			ev = v
		//		}

		mel3program.Walk(ev, a.StartProgram)

		fmt.Println("\t" + ev.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
