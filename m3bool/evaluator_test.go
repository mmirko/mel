package m3bool

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3uintEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uint evaluator ----")

	a := new(M3boolMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{
		`
and(m3boolconst(true),m3boolconst(true))

`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		ev := new(Evaluator)
		ev.Impl = a.Implementation
		ev.Mux = M3boolmux
		ev.Result = new(mel3program.Mel3Program)

		mel3program.Walk(ev, a.StartProgram)

		fmt.Println("\t" + ev.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
