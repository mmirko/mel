package rectangular

import (
	"math/rand"

	mel "github.com/mmirko/mel"
)

func (eObj *RectangularMe3li) Mutate(ep *mel.EvolutionParameters) {
	choose := rand.Intn(len(eObj.list))
	eObj.list[choose] = rectGenerate(ep)
}

func Mutate(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*RectangularMe3li)
	newProg := (prog.MelCopy()).(*RectangularMe3li)
	newProg.Mutate(ep)
	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}
