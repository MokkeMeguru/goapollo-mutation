package pkg

import (
	"fmt"
	"math/rand"

	"golang.org/x/xerrors"
)

type RandomGenerator interface {
	RandInt(closedLeft, openedRight int) (int, error)
}

type BasicRandomGenerator struct {
	Seed int64
}

func (m *BasicRandomGenerator) init() {
	rand.Seed(m.Seed)
}

func NewBasicRandomGenerator(seed int64) *BasicRandomGenerator {
	rndg := &BasicRandomGenerator{
		Seed: seed,
	}
	rndg.init()
	return rndg
}

func (m *BasicRandomGenerator) RandInt(closedLeft, openedRight int) (int, error) {
	if openedRight-closedLeft <= 0 {
		return 0, xerrors.New(fmt.Sprintf("invalid argument: closedLeft %d should be less than openedRight %d", closedLeft, openedRight))
	}
	return rand.Intn(openedRight-closedLeft) + closedLeft, nil
}
