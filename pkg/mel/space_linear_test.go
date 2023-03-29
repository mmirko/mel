package mel

import (
	"fmt"
	"testing"
)

func TestLinearVSpace(t *testing.T) {

	fmt.Println("---- Test: Linear Space ----")

	var s vSpace
	s = new(linearSpace)

	l := 5
	if err := s.init(fmt.Sprint(l)); err != nil {
		t.Error(err)
	}

	fmt.Print(s.dump())

	for i := 0; i < l; i++ {
		p := s.getPoint(fmt.Sprint(i))
		fmt.Println(p.getNeighborhood())
	}

	fmt.Println("---- End test ----")

}
