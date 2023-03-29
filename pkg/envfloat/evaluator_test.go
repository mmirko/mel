package envfloat

import (

	//m3uint "github.com/mmirko/mel/pkg/m3uint"

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

	tests := []string{`
readinput(m3uintconst(0)
)`, `
writeoutput(m3uintconst(1),m3numberconst(4.34)
)`}

	env := new(EnvFloat)
	env.Init([]float32{1.5}, 3)
	a.Mel3Object.Environment = env

	fmt.Println(a.Mel3Object.Environment)

	for _, iString := range tests {
		if err := a.MelStringImport(iString); err != nil {
			t.Errorf("Error importing: %s", err)
		}
		a.Compute()
		fmt.Println(a.Inspect())
		fmt.Println(a.Mel3Object.Environment)
	}

}
