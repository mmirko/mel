package bool3

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestBool3Importer(t *testing.T) {

	fmt.Println("---- Test: Bool3 importer ----")

	a := new(Bool3_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

	istrings := []string{`
	bool3const(true)
`}

	for i := 0; i < len(istrings); i++ {
		fmt.Println("Importing: " + istrings[i])
		err := a.Mel_string_import(istrings[i])
		if err != nil {
			fmt.Println(err.Error())
		} else {
			a.Mel_dump()
		}
		fmt.Println("---")
	}

	fmt.Println("---- End test ----")

}
