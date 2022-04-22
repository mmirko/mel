package m3uintcmp

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3uintcmpEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uintcmp evaluator ----")

	a := new(M3uintcmp_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

	istrings := []string{`
eq(
	m3uintconst(2),
	m3uintconst(4)
)
`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.Mel_string_import(istring)

		fmt.Println("\tEvaluating: " + istring)

		v := new(Evaluator)
		v.Impl = a.Implementation
		v.Mux = M3uintcmpmux

		ev := v

		mel3program.Walk(ev, a.StartProgram)

		fmt.Println("\t" + ev.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
