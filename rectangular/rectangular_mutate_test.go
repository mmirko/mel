package rectangular

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	mel "github.com/mmirko/mel"
)

func TestRectangularMutate(t *testing.T) {

	fmt.Println("---- Test: Rectangular Mutate ----")

	// Random seed based on seconds since epoch
	rand.Seed(int64(time.Now().Second()))

	ep := new(mel.EvolutionParameters)

	ep.SetValue("width", "800")
	ep.SetValue("height", "600")

	for i := 0; i < 1; i++ {
		cTest := Generate(ep)
		mutation := Mutate(cTest, ep)
		fmt.Println("Generated: ")
		fmt.Println("[", cTest, "]")
		fmt.Println("Mutated:")
		fmt.Println("[", mutation, "]")
	}

	fmt.Println("---- End test: Rectangular Mutate ----")

}
