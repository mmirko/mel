package rectangular

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mmirko/mel/pkg/mel"
)

func TestRectangularCrossover(t *testing.T) {

	fmt.Println("---- Test: Rectangular Crossover ----")

	// Random seed based on seconds since epoch
	rand.Seed(int64(time.Now().Second()))

	ep := new(mel.EvolutionParameters)

	ep.SetValue("width", "800")
	ep.SetValue("height", "600")

	for i := 0; i < 1; i++ {
		cTest1 := Generate(ep)
		cTest2 := Generate(ep)
		crossover := Crossover(cTest1, cTest2, ep)
		fmt.Println("Generated1: ")
		fmt.Println("[", cTest1, "]")
		fmt.Println("Generated2: ")
		fmt.Println("[", cTest2, "]")
		fmt.Println("Crossover:")
		fmt.Println("[", crossover, "]")
	}

	fmt.Println("---- End test: Rectangular Crossover ----")

}
