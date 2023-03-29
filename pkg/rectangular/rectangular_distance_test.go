package rectangular

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/mmirko/mel/pkg/mel"
)

func TestRectangularDistance(t *testing.T) {

	fmt.Println("---- Test: Rectangular Mutate ----")

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
		fmt.Println("Mutated:")
		fmt.Println("[", mutation, "]")

		im1, _ := cTest.(*RectangularMe3li).ToImage(ep)
		im2, _ := mutation.(*RectangularMe3li).ToImage(ep)

		fmt.Println("Distance:")
		fmt.Println(imageDistance(&im1, &im2))

	}

	fmt.Println("---- End test: Rectangular Distance ----")

}
