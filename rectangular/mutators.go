package rectangular

import (
	"math/rand"

	mel "github.com/mmirko/mel"
)

func (eObj *RectangularMe3li) MutateRectSubstitute(ep *mel.EvolutionParameters) {
	choose := rand.Intn(len(eObj.list))
	eObj.list[choose] = rectGenerate(ep)
}

func MutateRectSubstitute(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*RectangularMe3li)
	newProg := (prog.MelCopy()).(*RectangularMe3li)
	newProg.MutateRectSubstitute(ep)
	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}

func (eObj *RectangularMe3li) MutateRectElide(ep *mel.EvolutionParameters) {
	if len(eObj.list) > 1 {
		choose := rand.Intn(len(eObj.list))
		eObj.list = append(eObj.list[:choose], eObj.list[choose+1:]...)
	}
}

func MutateRectElide(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*RectangularMe3li)
	newProg := (prog.MelCopy()).(*RectangularMe3li)
	newProg.MutateRectElide(ep)
	return newProg
}
