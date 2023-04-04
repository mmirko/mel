package m3melbond

import (

	//m3uint "github.com/mmirko/mel/pkg/m3uint"

	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestLisevoEvaluator(t *testing.T) {

	a := new(M3MelBondMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	// c.VisitorCreatorSet = mel.VISDUMP
	c.Debug = true
	a.Init(c, ep, []string{"m3uint", "m3statements", "m3uintcmp", "m3bool"})

	tests := []string{`
eq(
	add(
		m3uintconst(1),
		m3uintconst(66)
	),
	m3uintconst(1)
)`, "m3boolconst(false)", `
multistmt(
	nop(),
	nop()
)`, "multistmt(nop(),nop())"}

	for i, iString := range tests {
		if i%2 == 1 {
			continue
		}

		if err := a.MelStringImport(iString); err != nil {
			t.Errorf("Error importing: %s", err)
		}
		a.Compute()
		if a.Inspect() != tests[i+1] {
			t.Errorf("Expected %s, got %s", tests[i+1], a.Inspect())
		}

	}
}
