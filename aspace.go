package mel

type PointVspace interface {
	GetNeighborhood() []PointVspace
	GetCamped() []Me3li
	GetDistance(PointVspace) int8
}
