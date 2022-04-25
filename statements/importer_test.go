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

	istrings := []string{`
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

	for i := 0; i < len(istrings); i++ {
		fmt.Println("Importing: " + istrings[i])
		a.MelStringImport(istrings[i])
		a.MelDump()
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
