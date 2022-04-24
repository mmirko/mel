package statements

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestSymbolicMathImporter(t *testing.T) {

	fmt.Println("---- Test: Statements importer ----")

	a := new(Statements_me3li)
	var ep *mel.EvolutionParameters
	a.Mel_init(ep)

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
