package m3redmon

import (
	"fmt"
	"testing"

	mel "github.com/mmirko/mel/pkg/mel"
)

func TestM3redmonEvaluator(t *testing.T) {

	fmt.Println("---- Test: M3redmon evaluator ----")

	a := new(M3redmonMe3li)
	var ep *mel.EvolutionParameters
	config := new(mel.MelConfig)
	a.MelInit(config, ep)

	istrings := []string{`
eq(
	add(
		diffdays(
			dateconst(2020-05-06),
			dateconst(2020-03-02)
		),
		m3uintconst(1)),
	m3uintconst(66)
)
`, `gt(diffdays(dateconst(today),dateconst(2019-05-20 16:44:59)),m3uintconst(9))`}
	expectedResults := []string{"m3boolconst(true)", "m3boolconst(true)"}

	for i, istring := range istrings {

		fmt.Println(">>>")

		fmt.Println("\tImporting: " + istring)
		if err := a.MelStringImport(istring); err != nil {
			t.Fatalf("MelStringImport(%q): %v", istring, err)
		}

		fmt.Println("\tEvaluating: " + istring)

		if err := a.Compute(); err != nil {
			t.Fatalf("Compute(%q): %v", istring, err)
		}
		if a.Result == nil {
			t.Fatalf("Compute(%q) produced no result", istring)
		}
		result := a.Inspect()
		fmt.Println("\t" + result)
		if result != expectedResults[i] {
			t.Errorf("expected %s, got %s", expectedResults[i], result)
		}

		fmt.Println("<<<")

	}
	fmt.Println("---- End test ----")

}
