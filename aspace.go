package mel

import ()

type Point_vspace interface {
	Get_neighbourhood() []Point_vspace
	Get_camped() []Me3li
	Get_distance(Point_vspace) int8
}
