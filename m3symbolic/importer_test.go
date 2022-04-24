package m3symbolic

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestSymbolicMathImporter(t *testing.T) {

	fmt.Println("---- Test: Symbolic math3 importer ----")

	a := new(Symbolic_math3_me3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{`
sum(
	mul(
		var(x),
		var(y)
	),
	const(5)
)
`}

	for i := 0; i < len(istrings); i++ {
		fmt.Println("Importing: " + istrings[i])
		err := a.MelStringImport(istrings[i])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			a.MelDump()
		}
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
