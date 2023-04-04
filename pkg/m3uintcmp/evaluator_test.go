package m3uintcmp

import (
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3uintcmpEvaluator(t *testing.T) {

	a := new(M3uintcmpMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = false
	a.MelInit(c, ep)

	tests := []string{"eq(m3uintconst(45),m3uintconst(45))", "m3boolconst(true)"}
	tests = append(tests, "eq(add(m3uintconst(45),m3uintconst(5)),m3uintconst(50))", "m3boolconst(true)")

	for i, iString := range tests {

		if i%2 == 1 {
			continue
		}

		a.MelStringImport(iString)
		a.Compute()
		if a.Inspect() != tests[i+1] {
			t.Errorf("Expected %s, got %s", tests[i+1], a.Inspect())
		}

	}
}
