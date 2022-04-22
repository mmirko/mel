package dates

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestDatesEvaluator(t *testing.T) {

	fmt.Println("---- Test: Dates evaluator ----")

	a := new(Dates_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

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
		a.Mel_string_import(istring)

		fmt.Println("\tEvaluating: " + istring)

		//		var ev mel3program.Visitor

		//		switch a.StartProgram.LibraryID {
		//		case MYLIBID:
		v := new(Evaluator)
		v.Impl = a.Implementation
		v.Mux = Datesmux
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
