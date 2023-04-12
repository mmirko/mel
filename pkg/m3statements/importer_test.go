package m3statements

import (
	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestSymbolicMathImporter(t *testing.T) {

	fmt.Println("---- Test: Statements importer ----")

	a := new(M3statementsMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(nil, ep)

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
		a.MelDump(nil)
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
