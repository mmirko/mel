package m3boolcmp

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3boolcmpImporter(t *testing.T) {

	fmt.Println("---- Test: M3boolcmp importer ----")

	a := new(M3boolcmpMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.MelInit(c, ep)

	istrings := []string{
		`
m3boolconst(true)

`,
		`
eq(
	m3boolconst(true),
	m3boolconst(false)
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
