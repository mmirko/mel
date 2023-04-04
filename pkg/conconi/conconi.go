package conconi

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/mmirko/mel/pkg/mel"
)

type Conconi struct {
	*mel.MelConfig
	m1         float32
	k1         float32
	m2         float32
	k2         float32
	crosspoint float32
}

func (c *Conconi) MelInit(config *mel.MelConfig, ep *mel.EvolutionParameters) {
	c.MelConfig = config
}

func (c *Conconi) MelCopy() mel.Me3li {
	newConc := new(Conconi)
	newConc.MelConfig = c.MelConfig
	newConc.m1 = c.m1
	newConc.m2 = c.m2
	newConc.k1 = c.k1
	newConc.k2 = c.k2
	newConc.crosspoint = c.crosspoint
	return newConc
}

func (c *Conconi) String() string {
	m1 := strconv.FormatFloat(float64(c.m1), 'f', 2, 64)
	m2 := strconv.FormatFloat(float64(c.m2), 'f', 2, 64)
	k1 := strconv.FormatFloat(float64(c.k1), 'f', 2, 64)
	k2 := strconv.FormatFloat(float64(c.k2), 'f', 2, 64)
	cc := strconv.FormatFloat(float64(c.crosspoint), 'f', 2, 64)
	return m1 + " " + k1 + " " + m2 + " " + k2 + " " + cc
}

func (c *Conconi) GetAT() (float32, float32) {
	x := (c.k2 - c.k1) / (c.m1 - c.m2)
	fca := c.m1*x + c.k1
	return x, fca
}
func (c *Conconi) GetParams() (float32, float32, float32, float32) {
	return c.m1, c.k1, c.m2, c.k2
}

func (c *Conconi) Generate(ep *mel.EvolutionParameters) {
	c.m1 = rand.Float32() * 20.0
	c.m2 = rand.Float32() * 20.0
	c.k1 = rand.Float32()*160.0 + 60.0 // FC
	c.k2 = rand.Float32()*160.0 + 60.0
	c.crosspoint = rand.Float32()*7.0 + 11.0 // m
}

func (c *Conconi) Mutate(ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		c.m1 = rand.Float32() * 20.0
	case 1:
		c.m2 = rand.Float32() * 20.0
	case 2:
		c.k1 = rand.Float32()*160.0 + 60.0
	case 3:
		c.k2 = rand.Float32()*160.0 + 60.0
	case 4:
		c.crosspoint = rand.Float32()*7.0 + 11.0
	}
}

func (c *Conconi) MutateSlow(ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		c.m1 = c.m1 + rand.Float32() - 0.5
	case 1:
		c.m2 = c.m2 + rand.Float32() - 0.5
	case 2:
		c.k1 = c.k1 + rand.Float32() - 0.5
	case 3:
		c.k2 = c.k2 + rand.Float32() - 0.5
	case 4:
		c.crosspoint = c.crosspoint + rand.Float32() - 0.5
	}
}

func (c *Conconi) CrossoverFake(sec *Conconi, ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		c.m1 = c.m1 + rand.Float32() - 0.5
	case 1:
		c.m2 = c.m2 + rand.Float32() - 0.5
	case 2:
		c.k1 = c.k1 + rand.Float32() - 0.5
	case 3:
		c.k2 = c.k2 + rand.Float32() - 0.5
	case 4:
		c.crosspoint = c.crosspoint + rand.Float32() - 0.5
	}
}

func ConconiFitness(in_prog *Conconi, x []float32, y []float32) (float32, bool) {

	var sumM float64
	for i := 0; i < 16; i++ {
		if x[i] < in_prog.crosspoint {
			sumM = sumM + math.Abs(float64(in_prog.m1*x[i]+in_prog.k1-y[i]))
		} else {
			sumM = sumM + math.Abs(float64(in_prog.m2*x[i]+in_prog.k2-y[i]))
		}
	}

	return float32(math.Exp(-1 * sumM / 100)), true
}

func ConconiGenerate(ep *mel.EvolutionParameters) mel.Me3li {
	var result mel.Me3li
	eObj := new(Conconi)
	eObj.MelInit(nil, ep)
	eObj.Generate(ep)
	result = eObj
	//fmt.Println("Generated",result)
	return result
}

func ConconiMutate(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*Conconi)
	newProg := (prog.MelCopy()).(*Conconi)
	newProg.Mutate(ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}

func ConconiMutateSlow(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*Conconi)
	newProg := (prog.MelCopy()).(*Conconi)
	newProg.MutateSlow(ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}

func ConconiCrossoverFake(p1 mel.Me3li, p2 mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog1 := p1.(*Conconi)
	prog2 := p2.(*Conconi)
	newProg := (prog1.MelCopy()).(*Conconi)
	newProg.CrossoverFake(prog2, ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}
