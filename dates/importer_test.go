package dates

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestDatesImporter(t *testing.T) {

	fmt.Println("---- Test: Dates importer ----")

	a := new(Dates_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

	istrings := []string{`dateconst(2014)`, `timestampconst(34242342)`, `
add(
	diffdays(
		dateconst(today),
		dateconst(2014-03-15)
	),
	m3uintconst(1)
)
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
