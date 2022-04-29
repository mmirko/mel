package rectangular

import "image"

// Image distance function
func imageDistance(img1p, img2p *image.Image) float64 {
	img1 := *img1p
	img2 := *img2p
	var sum float64
	for x := 0; x < img1.Bounds().Dx(); x++ {
		for y := 0; y < img1.Bounds().Dy(); y++ {
			r1, g1, b1, _ := img1.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			sum += float64(r1-r2) * float64(r1-r2)
			sum += float64(g1-g2) * float64(g1-g2)
			sum += float64(b1-b2) * float64(b1-b2)
		}
	}
	return sum / float64(img1.Bounds().Dx()*img1.Bounds().Dy()*3*256*256)
}
