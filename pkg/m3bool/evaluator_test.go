package m3bool

import (
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3boolEvaluator(t *testing.T) {

	a := new(M3boolMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = false
	a.MelInit(c, ep)

	tests := []string{"m3boolconst(true)", "m3boolconst(true)"}

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
