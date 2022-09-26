package m3bool

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uint evaluator ----")

	a := new(M3boolMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(nil, ep)

	istrings := []string{
		`
and(m3boolconst(true),m3boolconst(true))

`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		a.Walk()

		fmt.Println("\t" + a.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
