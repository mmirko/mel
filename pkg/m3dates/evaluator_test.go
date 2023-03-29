package m3dates

import (

	//m3uint "github.com/mmirko/mel/pkg/m3uint"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestM3datesEvaluator(t *testing.T) {

	a := new(M3datesMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = false
	a.MelInit(c, ep)

	tests := []string{"dateconst(2020-05-06)", "dateconst(2020-05-06)"}
	tests = append(tests, "add(diffdays(dateconst(2020-05-06),dateconst(2020-03-02)),m3uintconst(1))", "m3uintconst(66)")

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
