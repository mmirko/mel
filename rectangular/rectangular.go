package rectangular

import (
	"image/color"

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

func (m3 *RectangularMe3li) MelInit(c *mel.MelConfig, ep *mel.EvolutionParameters) {
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
