package envfloat

import (
	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/mel"
)

func TestEnvFloatEvaluator(t *testing.T) {

	a := new(EnvFloatMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.Init(c, ep, []string{"envfloat"})

	tests := []string{`readinput(m3uintconst(0))`,
		`writeoutput(m3uintconst(1),m3numberconst(4.34))`,
		`writeoutput(m3uintconst(1),add(m3numberconst(4.34),m3numberconst(1.0))))`,
		`pushkeep(m3numberconst(1.0))`,
		`pushkeep(m3numberconst(1.0))`,
		`writekeep(m3uintconst(1),add(m3numberconst(3.0),m3numberconst(1.0)))`,
		`popkeep()`,
	}

	env := new(EnvFloat)
	env.Init([]float32{1.5}, 3)

	var envI interface{}
	envI = env

	a.Mel3Object.Environment = &envI

	for _, iString := range tests {
		fmt.Println("-----", iString, "-----")
		fmt.Println(*a.Mel3Object.Environment)
		if err := a.MelStringImport(iString); err != nil {
			t.Errorf("Error importing: %s", err)
		}
		a.Compute()
		fmt.Println(a.Inspect())
		fmt.Println(*a.Mel3Object.Environment)
	}

}
