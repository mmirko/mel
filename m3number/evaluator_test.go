package m3number

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel"
	mel3program "github.com/mmirko/mel/mel3program"
)

func TestM3numberEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3number evaluator ----")

	a := new(M3numberMe3li)
	var ep *mel.EvolutionParameters
	a.MelInit(ep)

	istrings := []string{
		`
m3numberconst(54)

`,
		`
add(
        m3numberconst(3.6),
        m3numberconst(1E+1)
)

`,
		`
sub(
        m3numberconst(3),
        m3numberconst(11.2)
)

`,
		`
div(
        m3numberconst(5),
        m3numberconst(2)
)

`,
		`
mult(
        m3numberconst(3),
        m3numberconst(4.05)
)

`,
		`
mult(
	add(
        	m3numberconst(3),
		m3numberconst(5)
	),
        m3numberconst(2)
)

`}

	for _, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		a.MelStringImport(istring)

		fmt.Println("\tEvaluating: " + istring)

		ev := new(Evaluator)
		ev.Impl = a.Implementation
		ev.Mux = M3numbermux
		ev.Result = new(mel3program.Mel3_program)

		mel3program.Walk(ev, a.StartProgram)

		fmt.Println("\t" + ev.Inspect())

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
