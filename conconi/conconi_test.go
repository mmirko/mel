package conconi

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mmirko/mel"
)

func TestActions(t *testing.T) {

	rand.Seed(int64(time.Now().Unix()))

	fmt.Println("---- Test: Conconi ----")

	x := []float32{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600}
	y := []float32{70, 80, 90, 100, 110, 120, 130, 140, 150, 151, 152, 153, 154, 155, 156, 157}

	var ep *mel.EvolutionParameters
	for i := 0; i < 1000; i++ {
		cTest := ConconiGenerate(ep)
		mutation := ConconiMutate(cTest, ep)
		value, _ := ConconiFitness(cTest.(*Conconi), x, y)
		fmt.Println("Generated: [", cTest, "] - Fitness: [", value, "] - Mutated: [", mutation, "]")
	}

	fmt.Println("---- End test ----")
}
