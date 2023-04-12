package m3dates

import (
	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3datesImporter(t *testing.T) {

	fmt.Println("---- Test: M3dates importer ----")

	a := new(M3datesMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(nil, ep)

	istrings := []string{`dateconst(2014)`, `timestampconst(34242342)`, `
add(
	diffdays(
		dateconst(today),
		dateconst(2014-03-15)
	),
	m3uintconst(1)
)
`}
	for i := 0; i < len(istrings); i++ {
		fmt.Println("Importing: " + istrings[i])
		err := a.MelStringImport(istrings[i])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			a.MelDump(nil)
		}
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
