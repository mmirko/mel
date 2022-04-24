package m3uintcmp

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintcmpImporter(t *testing.T) {

	fmt.Println("---- Test: M3uintcmp importer ----")

	a := new(M3uintcmp_me3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{
		`
m3uintconst(54)

`,
		`
eq(
	m3uintconst(3),
	m3uintconst(1)
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
