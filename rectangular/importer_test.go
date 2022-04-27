package rectangular

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestRectangularImporter(t *testing.T) {

	fmt.Println("---- Test: Rectangular importer ----")

	a := new(RectangularMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{`
multistmt(
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	rectconst(prova-34-34-34-23),	
	nop(),	
	nop(),	
	nop()	
)
`}

	for i := 0; i < len(istrings); i++ {
		fmt.Println("Importing: " + istrings[i])
		fmt.Println(a.MelStringImport(istrings[i]))
		a.MelDump()
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
