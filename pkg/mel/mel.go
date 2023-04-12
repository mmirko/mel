package mel

import (
	"math/rand"
	"time"
)

type DumpConfig struct {
	Numeric bool
	Types   bool
}

type MelConfig struct {
	Debug bool
}

// The main interface, it states: It is a mel object
type Me3li interface {
	MelInit(*MelConfig, *EvolutionParameters)
	MelCopy() Me3li
}

type Me3liStringImport interface {
	MelStringImport(string) error
}

type Me3liDump interface {
	MelDump(*DumpConfig)
}

func Init() {
	rand.Seed(int64(time.Now().Unix()))
}
