package mel

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	LEFT = iota
	RIGHT
)

type pointLinearSpace struct {
	x int
	*linearSpace
}

func (p *pointLinearSpace) getIndex() string {
	return fmt.Sprint(p.x)
}

func (p *pointLinearSpace) getNeighborhood() []pointVSpace {
	res := make([]pointVSpace, 2)
	if p.x == 0 {
		res[LEFT] = nil
		res[RIGHT] = p.linearSpace.neighborhood[p.x+1]
	} else if p.x == len(p.linearSpace.neighborhood)-1 {
		res[LEFT] = p.linearSpace.neighborhood[p.x-1]
		res[RIGHT] = nil
	} else {
		res[LEFT] = p.linearSpace.neighborhood[p.x-1]
		res[RIGHT] = p.linearSpace.neighborhood[p.x+1]
	}
	return res
}

type linearSpace struct {
	neighborhood []*pointLinearSpace
}

func (s *linearSpace) init(length string) error {
	if s == nil {
		return errors.New("uninitialized")
	}
	// Convert length to int
	n, err := strconv.Atoi(length)
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("length must be greater than 0")
	}

	s.neighborhood = make([]*pointLinearSpace, n)

	for i := 0; i < n; i++ {
		s.neighborhood[i] = new(pointLinearSpace)
		s.neighborhood[i].x = i
		s.neighborhood[i].linearSpace = s
	}

	return nil
}

func (s *linearSpace) dump() string {
	var res string
	for _, p := range s.neighborhood {
		res += strconv.Itoa(p.x) + ":\n"
	}
	return res
}

func (s *linearSpace) getPoint(index string) pointVSpace {
	i, err := strconv.Atoi(index)
	if err != nil {
		return nil
	}
	if i < 0 || i >= len(s.neighborhood) {
		return nil
	}
	return s.neighborhood[i]
}

func (s *linearSpace) getPoints() []pointVSpace {
	res := make([]pointVSpace, len(s.neighborhood))
	for i, p := range s.neighborhood {
		res[i] = p
	}
	return res
}
