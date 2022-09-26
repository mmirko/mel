package m3uint

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uint evaluator ----")

	a := new(M3uintMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.MelInit(c, ep)

	istrings := []string{
		`
m3uintconst(54)

`,
		`
add(
        m3uintconst(3),
        m3uintconst(1)
)

`,
		`
sub(
        m3uintconst(3),
        m3uintconst(1)
)

`,
		`
div(
        m3uintconst(3),
        m3uintconst(1)
)

`,
		`
mult(
        m3uintconst(3),
        m3uintconst(1)
)

`,
		`
mult(
	add(
        	m3uintconst(3),
		m3uintconst(5)
	),
        m3uintconst(2)
)

`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		a.Compute()

		fmt.Println("\t" + a.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
