package m3uintcmp

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintcmpImporter(t *testing.T) {

	fmt.Println("---- Test: M3uintcmp importer ----")

	a := new(M3uintcmp_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

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
		err := a.Mel_string_import(istrings[i])
		fmt.Println("---")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			a.Mel_dump()
		}
		fmt.Println("<<<")
	}

	fmt.Println("---- End test ----")

}
