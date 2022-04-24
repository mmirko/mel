package mel

import (
	"math/rand"
	"time"
)

// The main interface, it states: It is a mel object
type Me3li interface {
	MelInit(*EvolutionParameters)
	MelCopy() Me3li
}

type Me3li_string_import interface {
	MelStringImport(string) error
}

type Me3li_dump interface {
	MelDump()
}

func Init() {
	rand.Seed(int64(time.Now().Unix()))
}
