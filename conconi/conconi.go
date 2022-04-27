package conconi

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/mmirko/mel"
)

type Conconi struct {
	m1         float32
	k1         float32
	m2         float32
	k2         float32
	crosspoint float32
}

func (gen *Conconi) MelInit(ep *mel.EvolutionParameters) {
}

func (gen *Conconi) MelCopy() mel.Me3li {
	newconc := new(Conconi)
	newconc.m1 = gen.m1
	newconc.m2 = gen.m2
	newconc.k1 = gen.k1
	newconc.k2 = gen.k2
	newconc.crosspoint = gen.crosspoint
	return newconc
}

func (c *Conconi) String() string {
	m1 := strconv.FormatFloat(float64(c.m1), 'f', 2, 64)
	m2 := strconv.FormatFloat(float64(c.m2), 'f', 2, 64)
	k1 := strconv.FormatFloat(float64(c.k1), 'f', 2, 64)
	k2 := strconv.FormatFloat(float64(c.k2), 'f', 2, 64)
	cc := strconv.FormatFloat(float64(c.crosspoint), 'f', 2, 64)
	return m1 + " " + k1 + " " + m2 + " " + k2 + " " + cc
}

func (c *Conconi) Get_AT() (float32, float32) {
	x := (c.k2 - c.k1) / (c.m1 - c.m2)
	fca := c.m1*x + c.k1
	return x, fca
}
func (c *Conconi) Get_params() (float32, float32, float32, float32) {
	return c.m1, c.k1, c.m2, c.k2
}

func (eobj *Conconi) Generate(ep *mel.EvolutionParameters) {
	eobj.m1 = rand.Float32() * 20.0
	eobj.m2 = rand.Float32() * 20.0
	eobj.k1 = rand.Float32()*160.0 + 60.0 // FC
	eobj.k2 = rand.Float32()*160.0 + 60.0
	eobj.crosspoint = rand.Float32()*7.0 + 11.0 // m
}

func (eobj *Conconi) Mutate(ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		eobj.m1 = rand.Float32() * 20.0
	case 1:
		eobj.m2 = rand.Float32() * 20.0
	case 2:
		eobj.k1 = rand.Float32()*160.0 + 60.0
	case 3:
		eobj.k2 = rand.Float32()*160.0 + 60.0
	case 4:
		eobj.crosspoint = rand.Float32()*7.0 + 11.0
	}
}

func (eobj *Conconi) MutateSlow(ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		eobj.m1 = eobj.m1 + rand.Float32() - 0.5
	case 1:
		eobj.m2 = eobj.m2 + rand.Float32() - 0.5
	case 2:
		eobj.k1 = eobj.k1 + rand.Float32() - 0.5
	case 3:
		eobj.k2 = eobj.k2 + rand.Float32() - 0.5
	case 4:
		eobj.crosspoint = eobj.crosspoint + rand.Float32() - 0.5
	}
}

func (eobj *Conconi) CrossoverFake(sec *Conconi, ep *mel.EvolutionParameters) {

	choose := rand.Intn(5)
	switch choose {
	case 0:
		eobj.m1 = eobj.m1 + rand.Float32() - 0.5
	case 1:
		eobj.m2 = eobj.m2 + rand.Float32() - 0.5
	case 2:
		eobj.k1 = eobj.k1 + rand.Float32() - 0.5
	case 3:
		eobj.k2 = eobj.k2 + rand.Float32() - 0.5
	case 4:
		eobj.crosspoint = eobj.crosspoint + rand.Float32() - 0.5
	}
}

func ConconiFitness(in_prog *Conconi, x []float32, y []float32) (float32, bool) {

	var summ float64
	for i := 0; i < 16; i++ {
		if x[i] < in_prog.crosspoint {
			summ = summ + math.Abs(float64(in_prog.m1*x[i]+in_prog.k1-y[i]))
		} else {
			summ = summ + math.Abs(float64(in_prog.m2*x[i]+in_prog.k2-y[i]))
		}
	}

	return float32(math.Exp(-1 * summ / 100)), true
}

func ConconiGenerate(ep *mel.EvolutionParameters) mel.Me3li {
	var result mel.Me3li
	var eobj *Conconi
	eobj = new(Conconi)
	eobj.MelInit(ep)
	eobj.Generate(ep)
	result = eobj
	//fmt.Println("Generated",result)
	return result
}

func ConconiMutate(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*Conconi)
	newprog := (prog.MelCopy()).(*Conconi)
	newprog.Mutate(ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newprog
}

func ConconiMutateSlow(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*Conconi)
	newprog := (prog.MelCopy()).(*Conconi)
	newprog.MutateSlow(ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newprog
}

func ConconiSrossover_fake(p1 mel.Me3li, p2 mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog1 := p1.(*Conconi)
	prog2 := p2.(*Conconi)
	newprog := (prog1.MelCopy()).(*Conconi)
	newprog.CrossoverFake(prog2, ep)
	//	//fmt.Println("Mutated ",prog, " in ",&result)
	return newprog
}
