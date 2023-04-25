package lisevo

import (

	//m3uint "github.com/mmirko/mel/pkg/m3uint"

	"fmt"
	"testing"

	"github.com/mmirko/mel/pkg/m3bool"
	"github.com/mmirko/mel/pkg/m3statements"
	"github.com/mmirko/mel/pkg/m3uint"
	"github.com/mmirko/mel/pkg/mel"
	"github.com/mmirko/mel/pkg/mel3program"
)

func TestLisevoGenerator(t *testing.T) {

	a := new(LisevoMe3li)
	var ep *mel.EvolutionParameters
	c := new(mel.MelConfig)
	c.Debug = true
	a.Init(c, ep, []string{"m3uint", "m3statements", "m3uintcmp", "m3bool"})

	gm := mel3program.CreateGenerationMatrix(a.Implementation)
	gm.AddTerminalGenerator(mel3program.ProgType{LibraryID: m3uint.MYLIBID, ProgramID: m3uint.M3UINTCONST, Arity: 0}, m3uint.M3UintConstGenerator)
	gm.AddTerminalGenerator(mel3program.ProgType{LibraryID: m3bool.MYLIBID, ProgramID: m3bool.CONST, Arity: 0}, m3bool.M3BoolConstGenerator)
	gm.AddTerminalGenerator(mel3program.ProgType{LibraryID: m3bool.MYLIBID, ProgramID: m3bool.VAR, Arity: 0}, m3bool.M3BoolVarGenerator)
	gm.AddTerminalGenerator(mel3program.ProgType{LibraryID: m3statements.MYLIBID, ProgramID: m3statements.NOP, Arity: 0}, m3statements.M3StmtNopGenerator)
	fmt.Println(gm.Init())
}
