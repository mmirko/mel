package rectangular

import (
	"errors"
	"image"
	"image/draw"

	mel "github.com/mmirko/mel"
)

// Image distance function
func imageDistance(img1p, img2p *image.Image) float64 {
	img1 := *img1p
	img2 := *img2p
	var sum float64
	for x := 0; x < img1.Bounds().Dx(); x++ {
		for y := 0; y < img1.Bounds().Dy(); y++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			sump := float64(0)
			sump += (float64(r1) - float64(r2)) * (float64(r1) - float64(r2)) / 4294967296.0
			sump += (float64(g1) - float64(g2)) * (float64(g1) - float64(g2)) / 4294967296.0
			sump += (float64(b1) - float64(b2)) * (float64(b1) - float64(b2)) / 4294967296.0
			sump /= float64(3.0)
			sum += sump
		}
	}
	return sum / (float64(img1.Bounds().Dx()) * float64(img1.Bounds().Dy()))
}

// Convert the mel.Me3li rectangle list to an image.Image
func (eObj *RectangularMe3li) ToImage(ep *mel.EvolutionParameters) (image.Image, error) {

	var height int
	var width int

	if widthR, ok := ep.GetInt("width"); ok {
		width = widthR

	} else {
		return nil, errors.New("width not set")
	}

	if heightR, ok := ep.GetInt("height"); ok {
		height = heightR

	} else {
		return nil, errors.New("height not set")
	}

	im := image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for _, r := range eObj.list {
		draw.Draw(im, image.Rect(int(r.x0), int(r.y0), int(r.x1), int(r.y1)), &image.Uniform{r.pColor}, image.ZP, draw.Src)
	}
	return im, nil
}
