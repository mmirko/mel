package m3uintcmp

import (
	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3uintcmpImporter(t *testing.T) {

	fmt.Println("---- Test: M3uintcmp importer ----")

	a := new(M3uintcmpMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.MelInit(c, ep)

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
			a.MelDump(nil)
		}
		fmt.Println("<<<")
	}

	fmt.Println("---- End test ----")

}
