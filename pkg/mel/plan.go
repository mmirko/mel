package mel

type individual struct {
	code           *Me3li
	fitness_values []float32
	prev           *individual
	next           *individual
}

type Population struct {
	Population_name   string
	PopulationHead    *individual
	NewbornHead       *individual
	GeneticGenerators []interface{}
	GeneticUnary      []interface{}
	GeneticBinary     []interface{}
	WeightGenerators  []float32
	WeightUnary       []float32
	WeightBinary      []float32
	WeightDeath       float32
	Threads           int
}

type Fitness struct {
	Fitness_name    string
	FitnessFunction func([]Me3li) (float32, bool)
	Threads         int
}

type Plan struct {
	Populations []Population
	Fitnesses   []Fitness
	ExitAt      int
}

type g0 func(*EvolutionParameters) Me3li
type g1 func(Me3li, *EvolutionParameters) Me3li
type g2 func(Me3li, Me3li, *EvolutionParameters) Me3li

func (plan *Plan) GetBest() (*Me3li, float32) {
	result := plan.Populations[0].PopulationHead.code
	resultFit := plan.Populations[0].PopulationHead.fitness_values[0]
	return result, resultFit
}
