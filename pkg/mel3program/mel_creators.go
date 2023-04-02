package mel3program

import "github.com/mmirko/mel/pkg/mel"

func CreateGenericCreators(c *mel.MelConfig, ep *mel.EvolutionParameters) map[uint16]Mel3VisitorCreator {

}

func BasmCreator() mel3program.Mel3Visitor {
	return new(Evaluator)
}

func DumpCreator() mel3program.Mel3Visitor {
	return new(Evaluator)
}
