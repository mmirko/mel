package m3uint

import (
	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3uintImporter(t *testing.T) {

	fmt.Println("---- Test: M3uint importer ----")

	a := new(M3uintMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(nil, ep)

	istrings := []string{
		`
m3uintconst(54)

`,
		`
add(
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
