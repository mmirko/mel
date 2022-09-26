package rectangular

import (
	"image/color"
	"math/rand"

	mel "github.com/mmirko/mel"
)

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

func Generate(ep *mel.EvolutionParameters) mel.Me3li {
	var result mel.Me3li
	eObj := new(RectangularMe3li)
	eObj.MelInit(nil, ep)
	eObj.Generate(ep)
	result = eObj
	//fmt.Println("Generated",result)
	return result
}
