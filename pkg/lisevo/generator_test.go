package lisevo

import (

	//m3uint "github.com/mmirko/mel/pkg/m3uint"

	"testing"

	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

func TestLisevoGenerator(t *testing.T) {

	a := new(LisevoMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.Init(c, ep, []string{"m3uint", "m3statements", "m3uintcmp", "m3bool"})

	mel3program.CreateGenerationMatrix(a.Implementation)

}
