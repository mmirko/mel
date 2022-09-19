package statements

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestSymbolicMathImporter(t *testing.T) {

	fmt.Println("---- Test: Statements importer ----")

	a := new(StatementsMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	iStrings := []string{`
multistmt(
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop(),	
	nop()	
)
`}

	for i := 0; i < len(iStrings); i++ {
		fmt.Println("Importing: " + iStrings[i])
		a.MelStringImport(iStrings[i])
		a.MelDump()
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
