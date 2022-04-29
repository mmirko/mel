package rectangular

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"

	mel "github.com/mmirko/mel"
)

type Rect struct {
	pColor color.Color
	x0     uint16
	y0     uint16
	x1     uint16
	y1     uint16
}

type RectangularMe3li struct {
	list []Rect
}

func (m3 *RectangularMe3li) MelInit(ep *mel.EvolutionParameters) {
}

func (m3 *RectangularMe3li) MelCopy() mel.Me3li {
	result := new(RectangularMe3li)
	result.list = make([]Rect, len(m3.list))

	for i, rect := range m3.list {
		result.list[i].pColor = rect.pColor
		result.list[i].x0 = rect.x0
		result.list[i].x1 = rect.x1
		result.list[i].y0 = rect.y0
		result.list[i].y1 = rect.y1
	}

	return result
}

func rectGenerate(ep *mel.EvolutionParameters) Rect {

	width, _ := ep.GetInt("width")
	height, _ := ep.GetInt("height")

	var result Rect
	result.x0 = uint16(rand.Intn(width))
	result.x1 = uint16(rand.Intn(width))
	result.y0 = uint16(rand.Intn(height))
	result.y1 = uint16(rand.Intn(height))
	result.pColor = color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
	return result
}

func (eObj *RectangularMe3li) Generate(ep *mel.EvolutionParameters) {
	n := 20
	eObj.list = make([]Rect, n)
	for i := 0; i < n; i++ {
		eObj.list[i] = rectGenerate(ep)
	}
}

func (eObj *RectangularMe3li) Mutate(ep *mel.EvolutionParameters) {

	choose := rand.Intn(len(eObj.list))
	eObj.list[choose] = rectGenerate(ep)
}

func (eobj *RectangularMe3li) Crossover(sec *RectangularMe3li, ep *mel.EvolutionParameters) {
}

func Generate(ep *mel.EvolutionParameters) mel.Me3li {
	var result mel.Me3li
	eObj := new(RectangularMe3li)
	eObj.MelInit(ep)
	eObj.Generate(ep)
	result = eObj
	//fmt.Println("Generated",result)
	return result
}

func Mutate(p mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog := p.(*RectangularMe3li)
	newProg := (prog.MelCopy()).(*RectangularMe3li)
	newProg.Mutate(ep)
	//fmt.Println("Mutated ",prog, " in ",&result)
	return newProg
}

func Crossover(p1 mel.Me3li, p2 mel.Me3li, ep *mel.EvolutionParameters) mel.Me3li {
	prog1 := p1.(*RectangularMe3li)
	prog2 := p2.(*RectangularMe3li)
	newProg := (prog1.MelCopy()).(*RectangularMe3li)
	newProg.Crossover(prog2, ep)
	return newProg
}

func Fitness(r *RectangularMe3li, target *image.Image, ep *mel.EvolutionParameters) (float32, bool) {

	if genImage, err := r.ToImage(ep); err == nil {
		distance := imageDistance(&genImage, target)
		return float32(math.Exp(-1 * distance)), true
	}
	return 0, false

}

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
