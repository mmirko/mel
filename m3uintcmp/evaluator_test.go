package m3uintcmp

import (
	"fmt"
	//m3uint "github.com/mmirko/mel/m3uint"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintcmpEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uintcmp evaluator ----")

	a := new(M3uintcmpMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.MelInit(c, ep)

	istrings := []string{`
eq(
	m3uintconst(2),
	m3uintconst(4)
)
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
