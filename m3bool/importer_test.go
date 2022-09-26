package m3bool

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3boolImporter(t *testing.T) {

	fmt.Println("---- Test: M3bool importer ----")

	a := new(M3boolMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(nil, ep)

	istrings := []string{`
	m3boolconst(true)
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
