package rectangular

import mel "github.com/mmirko/mel"

func (eobj *RectangularMe3li) Crossover(sec *RectangularMe3li, ep *mel.EvolutionParameters) {
}

func Crossover(p1 mel.Me3li, p2 mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog1 := p1.(*RectangularMe3li)
	prog2 := p2.(*RectangularMe3li)
	newProg := (prog1.MelCopy()).(*RectangularMe3li)
	newProg.Crossover(prog2, ep)
	return newProg
}
