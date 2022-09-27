package m3uint

import (
	"testing"

	mel "github.com/mmirko/mel"
)

func TestM3uintEvaluator(t *testing.T) {

	a := new(M3uintMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = false
	a.MelInit(c, ep)

	tests := []string{"m3uintconst(45)", "m3uintconst(45)"}
	tests = append(tests, "add(m3uintconst(4),m3uintconst(2))", "m3uintconst(6)")
	tests = append(tests, "mult(m3uintconst(3),m3uintconst(5))", "m3uintconst(15)")

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
