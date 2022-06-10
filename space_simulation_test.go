package mel

import (
	"fmt"
	"testing"
)

func TestSpaceSimulation(t *testing.T) {

	fmt.Println("---- Test: Space Simulation ----")

	var s vSpace
	s = new(linearSpace)

	l := 5
	if err := s.init(fmt.Sprint(l)); err != nil {
		t.Error(err)
	}

	SpaceSimulate(s)

	fmt.Println("---- End test ----")

}
