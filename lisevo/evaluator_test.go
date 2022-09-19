package lisevo

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestLisevoEvaluator(t *testing.T) {

	fmt.Println("---- Test: Lisevo evaluator ----")

	a := new(LisevoMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{`
eq(
	add(
		m3uintconst(1),
		m3uintconst(66)
	),
	m3uintconst(1)
)
`, `multistmt(
	gt(m3uintconst(3),m3uintconst(9)),
	lt(m3uintconst(3),m3uintconst(9))
)`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		v := new(Evaluator)
		v.Impl = a.Implementation
		v.Mux = lisevoMux

		mel3program.Walk(v, a.StartProgram)

		fmt.Println("\t" + v.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
