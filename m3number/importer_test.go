package m3number

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3numberImporter(t *testing.T) {

	fmt.Println("---- Test: M3number importer ----")

	a := new(M3numberMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{
		`
m3numberconst(54)

`,
		`
add(
	m3numberconst(3),
	m3numberconst(1)
)

`}

	for i := 0; i < len(istrings); i++ {
		fmt.Println(">>>")
		fmt.Println(istrings[i])
		err := a.MelStringImport(istrings[i])
		fmt.Println("---")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			a.MelDump()
		}
		fmt.Println("<<<")
	}

	fmt.Println("---- End test ----")

}
