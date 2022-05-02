package mel

import (
	"fmt"
	"log"
	"math/rand"
)

type PlanSimple struct {
	Plan
	GenerationNumber int
	PopulationSize   int
	DeathsRate       float32
	UnaryRate        float32
	BinaryRate       float32
}

// Execute the simple evolution plan
func (plan *PlanSimple) Execute(ep *EvolutionParameters) {

	var head *individual

	// Evolution variables
	gennum := plan.GenerationNumber
	popsize := plan.PopulationSize
	populationNum := len(plan.Populations)
	fitnessNum := len(plan.Fitnesses)

	if populationNum != 1 {
		log.Fatal("Simple Plan has to have 1 population")
	}

	if fitnessNum != 1 {
		log.Fatal("Simple Plan has to have 1 fitness function")
	}

	current_population := 0
	head = plan.Populations[0].Population_head

	// Creating the generators array and normalizing the generators weights
	generators_num := len(plan.Populations[0].Genetic_generators)
	if generators_num == 0 {
		log.Fatal("Simple Plan has to have at least a generator")
	}

	generators := make([]g0, generators_num)

	generator_weight_sum := float32(0.0)

	for i := 0; i < generators_num; i++ {
		generators[i] = plan.Populations[0].Genetic_generators[i].(func(*EvolutionParameters) Me3li)
		generator_weight_sum += plan.Populations[0].Weight_generators[i]
	}

	generators_weights := make([]float32, generators_num)

	for i := 0; i < generators_num; i++ {
		generators_weights[i] = plan.Populations[0].Weight_generators[i] / generator_weight_sum
	}

	// Normalizing the unary operators weights

	unary_num := len(plan.Populations[0].Genetic_unary)

	unary := make([]g1, unary_num)

	unary_weight_sum := float32(0.0)

	for i := 0; i < unary_num; i++ {
		unary[i] = plan.Populations[0].Genetic_unary[i].(func(Me3li, *EvolutionParameters) Me3li)
		unary_weight_sum += plan.Populations[0].Weight_unary[i]
	}

	unary_weights := make([]float32, unary_num)

	for i := 0; i < unary_num; i++ {
		unary_weights[i] = plan.Populations[0].Weight_unary[i] / unary_weight_sum
	}

	// Normalizing the binary operators weights

	binary_num := len(plan.Populations[0].Genetic_binary)

	binary := make([]g2, binary_num)

	binary_weight_sum := float32(0.0)

	for i := 0; i < binary_num; i++ {
		binary[i] = plan.Populations[0].Genetic_binary[i].(func(Me3li, Me3li, *EvolutionParameters) Me3li)
		binary_weight_sum += plan.Populations[0].Weight_binary[i]
	}

	binary_weights := make([]float32, binary_num)

	for i := 0; i < binary_num; i++ {
		binary_weights[i] = plan.Populations[0].Weight_binary[i] / binary_weight_sum
	}

	// Fitness
	fitness := plan.Fitnesses[0].FitnessFunction

	// Reading variables
	deathRate := plan.DeathsRate
	unaryRate := plan.UnaryRate
	binaryRate := plan.BinaryRate

	// Main loop: cycle throughout generations
	for generation := 0; generation < gennum; generation++ {

		generated := 0
		unaryApplied := 0
		binaryApplied := 0

		// Remove unfitted individuals
		removed := int(float32(current_population) * deathRate)
		cut := current_population - removed
		indiv_point := make([]*individual, cut)

		for i, curr := 0, head; curr != nil; curr = curr.next {
			indiv_point[i] = curr
			if i == cut-1 {
				current_population = i + 1
				curr.next = nil
				break
			}
			i++
		}

		// Apply genetic operators (unary and binary)

		var howmany_unary int
		var unary_container []*individual
		if unary_num != 0 {

			// Compute how many operators has to be applied
			howmany_unary = int(float32(current_population) * unaryRate)

			// Prepare the new elements containers
			unary_container = make([]*individual, howmany_unary)

			for i := 0; i < howmany_unary; i++ {
				undergoing := rand.Intn(current_population)
				u_indiv := indiv_point[undergoing]
				switch unary_num {
				case 1:
					newindiv := new(individual)
					newcode := unary[0](*(u_indiv.code), ep)
					codeslice := make([]Me3li, 1)
					codeslice[0] = newcode
					newfitness, _ := fitness(codeslice)
					newindiv.code = &newcode
					fitslice := make([]float32, 1)
					fitslice[0] = newfitness
					newindiv.fitness_values = fitslice
					unary_container[i] = newindiv
				default:
					choosen := rand.Float32()
					partial := float32(0.0)
					for j := 0; j < unary_num; j++ {
						partial = partial + unary_weights[j]
						if choosen < partial || j == unary_num-1 {
							newindiv := new(individual)
							newcode := unary[j](*(u_indiv.code), ep)
							codeslice := make([]Me3li, 1)
							codeslice[0] = newcode
							newfitness, _ := fitness(codeslice)
							newindiv.code = &newcode
							fitslice := make([]float32, 1)
							fitslice[0] = newfitness
							newindiv.fitness_values = fitslice
							unary_container[i] = newindiv
							break
						}
					}
				}
			}
		}

		var howmany_binary int
		var binary_container []*individual
		if binary_num != 0 {

			// Compute how many operators has to be applied
			howmany_binary = int(float32(current_population) * binaryRate)

			// Prepare the new elements containers
			binary_container = make([]*individual, howmany_binary)

			for i := 0; i < howmany_binary; i++ {
				undergoing1 := rand.Intn(current_population)
				undergoing2 := rand.Intn(current_population)
				u_indiv1 := indiv_point[undergoing1]
				u_indiv2 := indiv_point[undergoing2]
				switch binary_num {
				case 1:
					newindiv := new(individual)
					newcode := binary[0](*(u_indiv1.code), *(u_indiv2.code), ep)
					codeslice := make([]Me3li, 1)
					codeslice[0] = newcode
					newfitness, _ := fitness(codeslice)
					newindiv.code = &newcode
					fitslice := make([]float32, 1)
					fitslice[0] = newfitness
					newindiv.fitness_values = fitslice
					binary_container[i] = newindiv
				default:
					choosen := rand.Float32()
					partial := float32(0.0)
					for j := 0; j < binary_num; j++ {
						partial = partial + binary_weights[j]
						if choosen < partial || j == binary_num-1 {
							newindiv := new(individual)
							newcode := binary[j](*(u_indiv1.code), *(u_indiv2.code), ep)
							codeslice := make([]Me3li, 1)
							codeslice[0] = newcode
							newfitness, _ := fitness(codeslice)
							newindiv.code = &newcode
							fitslice := make([]float32, 1)
							fitslice[0] = newfitness
							newindiv.fitness_values = fitslice
							binary_container[i] = newindiv
							break
						}
					}
				}
			}
		}

		for i := 0; i < howmany_unary && current_population < popsize; i++ {
			ordered_place(&head, unary_container[i])
			current_population++
			unaryApplied++
		}

		for i := 0; i < howmany_binary && current_population < popsize; i++ {
			ordered_place(&head, binary_container[i])
			current_population++
			binaryApplied++
		}

		// Grow the population with new individuals (Apply generators)
		for i := current_population; i < popsize; i++ {
			switch generators_num {
			case 1:
				newindiv := new(individual)
				newcode := generators[0](ep)
				codeslice := make([]Me3li, 1)
				codeslice[0] = newcode
				newfitness, _ := fitness(codeslice)
				//fmt.Println(fitness(codeslice))
				newindiv.code = &newcode
				//fmt.Println(newcode)
				fitslice := make([]float32, 1)
				fitslice[0] = newfitness
				newindiv.fitness_values = fitslice
				ordered_place(&head, newindiv)
				current_population++
				generated++
			default:
				choosen := rand.Float32()
				partial := float32(0.0)
				for j := 0; j < generators_num; j++ {
					partial = partial + generators_weights[j]
					if choosen < partial || j == generators_num-1 {
						newindiv := new(individual)
						newcode := generators[j](ep)
						codeslice := make([]Me3li, 1)
						codeslice[0] = newcode
						newfitness, _ := fitness(codeslice)
						newindiv.code = &newcode
						fitslice := make([]float32, 1)
						fitslice[0] = newfitness
						newindiv.fitness_values = fitslice
						ordered_place(&head, newindiv)
						current_population++
						generated++
						break
					}
				}
			}
		}
		highestFitness := head.fitness_values[0]
		fmt.Println("Generation: ", generation, " - Population size: ", current_population, " - Removed: ", removed, "- Generated: ", generated, " - Unary: ", unaryApplied, " - Binary: ", binaryApplied, " - Highest: ", highestFitness)

	}

	plan.Populations[0].Population_head = head

}
