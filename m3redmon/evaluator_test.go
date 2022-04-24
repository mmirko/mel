package m3redmon

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3redmonEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3redmon evaluator ----")

	a := new(M3redmonMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{`
eq(
	add(
		diffdays(
			dateconst(2020-05-06),
			dateconst(2020-03-02)
		),
		m3uintconst(1)),
	m3uintconst(66)
)
`, `gt(diffdays(dateconst(today),dateconst(2019-05-20 16:44:59)),m3uintconst(9))`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		v := new(Evaluator)
		v.Impl = a.Implementation
		v.Mux = M3redmonmux

		mel3program.Walk(v, a.StartProgram)

		fmt.Println("\t" + v.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
