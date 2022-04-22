package m3uint

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3uintEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3uint evaluator ----")

	a := new(M3uint_me3li)
	var ep *mel.Evolution_parameters
	a.Mel_init(ep)

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
		a.Mel_string_import(istring)

		fmt.Println("\tEvaluating: " + istring)

		ev := new(Evaluator)
		ev.Impl = a.Implementation
		ev.Mux = M3uintmux
		ev.Result = new(mel3program.Mel3_program)

		mel3program.Walk(ev, a.StartProgram)

		fmt.Println("\t" + ev.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
