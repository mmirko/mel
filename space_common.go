package mel

type pointVSpace interface {
	getIndex() string
	getNeighborhood() []pointVSpace
	//getCamped() []Me3li
	//getDistance(pointVSpace) int8

}

type vSpace interface {
	init(string) error
	dump() string
	getPoint(string) pointVSpace
}
