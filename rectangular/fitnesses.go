package rectangular

import (
	"image"
	"math"

	mel "github.com/mmirko/mel"
)

func FitnessImageDistance(r *RectangularMe3li, target *image.Image, ep *mel.EvolutionParameters) (float32, bool) {

	if genImage, err := r.ToImage(ep); err == nil {
		distance := imageDistance(&genImage, target)
		return float32(math.Exp(-1 * distance)), true
	}
	return 0, false

}
