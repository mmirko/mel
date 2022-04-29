package rectangular

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"

	mel "github.com/mmirko/mel"
)

func TestRectangular(t *testing.T) {

	fmt.Println("---- Test: Rectangular Image Write----")

	// Random seed based on seconds since epoch
	rand.Seed(int64(time.Now().Second()))

	ep := new(mel.EvolutionParameters)

	ep.SetValue("width", "800")
	ep.SetValue("height", "600")

	for i := 0; i < 1; i++ {
		cTest := Generate(ep)

		out, err := os.Create("./output.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		oBj := cTest.(*RectangularMe3li)
		if img, err := oBj.ToImage(ep); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			err = png.Encode(out, img)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	fmt.Println("---- End test: Rectangular Image Write ----")

}
