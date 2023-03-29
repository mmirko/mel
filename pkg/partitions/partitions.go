package partitions

import (
	"github.com/mmirko/mel/pkg/mel"
)

type set []int

type PartMe3li struct {
	N int
	P []set
}

func (m3 *PartMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
}

func (m3 *PartMe3li) MelCopy() mel.Me3li {
	result := new(PartMe3li)
	result.N = m3.N
	result.P = make([]set, len(m3.P))

	for _, s := range m3.P {
		sc := make(set, len(s))
		copy(s, sc)
		result.P = append(result.P, sc)
	}

	return result
}
