package partitions

import (
	"math/rand"

	"github.com/mmirko/mel/pkg/mel"
)

func (eObj *PartMe3li) Generate(ep *mel.EvolutionParameters) {
	eObj.P = make([]set, 0)
	for i := 0; i < eObj.N; i++ {
		toSet := rand.Intn(len(eObj.P) + 1)
		switch toSet {
		case len(eObj.P):
			eObj.P = append(eObj.P, set{i})
		default:
			eObj.P[toSet] = append(eObj.P[toSet], i)
		}
	}
}

func Generate(ep *mel.EvolutionParameters) mel.Me3li {
	n, _ := ep.GetInt("setlength")
	var result mel.Me3li
	eObj := new(PartMe3li)
	eObj.N = n
	eObj.MelInit(nil, ep)
	eObj.Generate(ep)
	result = eObj
	//fmt.Println("Generated",result)
	return result
}
