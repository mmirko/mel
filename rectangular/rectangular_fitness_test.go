package rectangular

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	mel "github.com/mmirko/mel"
)

func TestRectangularFitness(t *testing.T) {

	fmt.Println("---- Test: Rectangular Fitness ----")

	// Random seed based on seconds since epoch
	rand.Seed(int64(time.Now().Second()))

	ep := new(mel.EvolutionParameters)

	ep.SetValue("width", "800")
	ep.SetValue("height", "600")

	for i := 0; i < 1; i++ {
		cTest := Generate(ep)
		mutation := MutateRectSubstitute(cTest, ep)
		fmt.Println("Generated: ")
		fmt.Println("[", cTest, "]")

		im2, _ := mutation.(*RectangularMe3li).ToImage(ep)

		fmt.Println("Fitness:")
		fmt.Println(FitnessImageDistance(cTest.(*RectangularMe3li), &im2, ep))

	}

	fmt.Println("---- End test: Rectangular Fitness ----")

}
