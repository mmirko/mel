package rectangular

import (
	"math/rand"

	"github.com/mmirko/mel/pkg/mel"
)

func (eObj *RectangularMe3li) Crossover(sec *RectangularMe3li, ep *mel.EvolutionParameters) {
	var choose1, choose2 int

	if len(eObj.list) > 1 && len(sec.list) > 1 {
		choose1 = rand.Intn(len(eObj.list))
		choose2 = rand.Intn(len(sec.list))
		eObj.list = append(eObj.list[:choose1], sec.list[choose2+1:]...)
	}
}

func Crossover(p1 mel.Me3li, p2 mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog1 := p1.(*RectangularMe3li)
	prog2 := p2.(*RectangularMe3li)
	newProg := (prog1.MelCopy()).(*RectangularMe3li)
	newProg.Crossover(prog2, ep)

	return newProg
}
