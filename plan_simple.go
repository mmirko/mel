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

func orderedPlace(queue_head **individual, newindiv *individual) {
	head := *queue_head
	tail := head
	if head == nil {
		*queue_head = newindiv
	} else {
		for curr := head; curr != nil; curr = curr.next {
			tail = curr
			if newindiv.fitness_values[0] > curr.fitness_values[0] {
				newindiv.prev = curr.prev
				newindiv.next = curr
				if curr.prev != nil {
					curr.prev.next = newindiv
				} else {
					*queue_head = newindiv
				}
				curr.prev = newindiv
				return
			}
		}
		newindiv.prev = tail
		tail.next = newindiv

	}
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

	currentPopulation := 0
	head = plan.Populations[0].PopulationHead

	// Creating the generators array and normalizing the generators weights
	generatorsNum := len(plan.Populations[0].GeneticGenerators)
	if generatorsNum == 0 {
		log.Fatal("Simple Plan has to have at least a generator")
	}

	generators := make([]g0, generatorsNum)

	generatorWeightSum := float32(0.0)

	for i := 0; i < generatorsNum; i++ {
		generators[i] = plan.Populations[0].GeneticGenerators[i].(func(*EvolutionParameters) Me3li)
		generatorWeightSum += plan.Populations[0].WeightGenerators[i]
	}

	generatorsWeights := make([]float32, generatorsNum)

	for i := 0; i < generatorsNum; i++ {
		generatorsWeights[i] = plan.Populations[0].WeightGenerators[i] / generatorWeightSum
	}

	// Normalizing the unary operators weights

	unaryNum := len(plan.Populations[0].GeneticUnary)

	unary := make([]g1, unaryNum)

	unaryWeightSum := float32(0.0)

	for i := 0; i < unaryNum; i++ {
		unary[i] = plan.Populations[0].GeneticUnary[i].(func(Me3li, *EvolutionParameters) Me3li)
		unaryWeightSum += plan.Populations[0].WeightUnary[i]
	}

	unaryWeights := make([]float32, unaryNum)

	for i := 0; i < unaryNum; i++ {
		unaryWeights[i] = plan.Populations[0].WeightUnary[i] / unaryWeightSum
	}

	// Normalizing the binary operators weights

	binaryNum := len(plan.Populations[0].GeneticBinary)

	binary := make([]g2, binaryNum)

	binaryWeightSum := float32(0.0)

	for i := 0; i < binaryNum; i++ {
		binary[i] = plan.Populations[0].GeneticBinary[i].(func(Me3li, Me3li, *EvolutionParameters) Me3li)
		binaryWeightSum += plan.Populations[0].WeightBinary[i]
	}

	binaryWeights := make([]float32, binaryNum)

	for i := 0; i < binaryNum; i++ {
		binaryWeights[i] = plan.Populations[0].WeightBinary[i] / binaryWeightSum
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
		removed := int(float32(currentPopulation) * deathRate)
		cut := currentPopulation - removed
		indivPoint := make([]*individual, cut)

		for i, curr := 0, head; curr != nil; curr = curr.next {
			indivPoint[i] = curr
			if i == cut-1 {
				currentPopulation = i + 1
				curr.next = nil
				break
			}
			i++
		}

		// Apply genetic operators (unary and binary)

		var howmanyUnary int
		var unaryContainer []*individual
		if unaryNum != 0 {

			// Compute how many operators has to be applied
			howmanyUnary = int(float32(currentPopulation) * unaryRate)

			// Prepare the new elements containers
			unaryContainer = make([]*individual, howmanyUnary)

			for i := 0; i < howmanyUnary; i++ {
				undergoing := rand.Intn(currentPopulation)
				uIndiv := indivPoint[undergoing]
				switch unaryNum {
				case 1:
					newIndiv := new(individual)
					newcode := unary[0](*(uIndiv.code), ep)
					codeslice := make([]Me3li, 1)
					codeslice[0] = newcode
					newfitness, _ := fitness(codeslice)
					newIndiv.code = &newcode
					fitslice := make([]float32, 1)
					fitslice[0] = newfitness
					newIndiv.fitness_values = fitslice
					unaryContainer[i] = newIndiv
				default:
					chosen := rand.Float32()
					partial := float32(0.0)
					for j := 0; j < unaryNum; j++ {
						partial = partial + unaryWeights[j]
						if chosen < partial || j == unaryNum-1 {
							newIndiv := new(individual)
							newcode := unary[j](*(uIndiv.code), ep)
							codeslice := make([]Me3li, 1)
							codeslice[0] = newcode
							newfitness, _ := fitness(codeslice)
							newIndiv.code = &newcode
							fitslice := make([]float32, 1)
							fitslice[0] = newfitness
							newIndiv.fitness_values = fitslice
							unaryContainer[i] = newIndiv
							break
						}
					}
				}
			}
		}

		var howmanyBinary int
		var binaryContainer []*individual
		if binaryNum != 0 {

			// Compute how many operators has to be applied
			howmanyBinary = int(float32(currentPopulation) * binaryRate)

			// Prepare the new elements containers
			binaryContainer = make([]*individual, howmanyBinary)

			for i := 0; i < howmanyBinary; i++ {
				undergoing1 := rand.Intn(currentPopulation)
				undergoing2 := rand.Intn(currentPopulation)
				uIndiv1 := indivPoint[undergoing1]
				uIndiv2 := indivPoint[undergoing2]
				switch binaryNum {
				case 1:
					newIndiv := new(individual)
					newcode := binary[0](*(uIndiv1.code), *(uIndiv2.code), ep)
					codeslice := make([]Me3li, 1)
					codeslice[0] = newcode
					newfitness, _ := fitness(codeslice)
					newIndiv.code = &newcode
					fitslice := make([]float32, 1)
					fitslice[0] = newfitness
					newIndiv.fitness_values = fitslice
					binaryContainer[i] = newIndiv
				default:
					chosen := rand.Float32()
					partial := float32(0.0)
					for j := 0; j < binaryNum; j++ {
						partial = partial + binaryWeights[j]
						if chosen < partial || j == binaryNum-1 {
							newIndiv := new(individual)
							newcode := binary[j](*(uIndiv1.code), *(uIndiv2.code), ep)
							codeslice := make([]Me3li, 1)
							codeslice[0] = newcode
							newfitness, _ := fitness(codeslice)
							newIndiv.code = &newcode
							fitslice := make([]float32, 1)
							fitslice[0] = newfitness
							newIndiv.fitness_values = fitslice
							binaryContainer[i] = newIndiv
							break
						}
					}
				}
			}
		}

		for i := 0; i < howmanyUnary && currentPopulation < popsize; i++ {
			orderedPlace(&head, unaryContainer[i])
			currentPopulation++
			unaryApplied++
		}

		for i := 0; i < howmanyBinary && currentPopulation < popsize; i++ {
			orderedPlace(&head, binaryContainer[i])
			currentPopulation++
			binaryApplied++
		}

		// Grow the population with new individuals (Apply generators)
		for i := currentPopulation; i < popsize; i++ {
			switch generatorsNum {
			case 1:
				newIndiv := new(individual)
				newcode := generators[0](ep)
				codeslice := make([]Me3li, 1)
				codeslice[0] = newcode
				newfitness, _ := fitness(codeslice)
				//fmt.Println(fitness(codeslice))
				newIndiv.code = &newcode
				//fmt.Println(newcode)
				fitslice := make([]float32, 1)
				fitslice[0] = newfitness
				newIndiv.fitness_values = fitslice
				orderedPlace(&head, newIndiv)
				currentPopulation++
				generated++
			default:
				chosen := rand.Float32()
				partial := float32(0.0)
				for j := 0; j < generatorsNum; j++ {
					partial = partial + generatorsWeights[j]
					if chosen < partial || j == generatorsNum-1 {
						newIndiv := new(individual)
						newcode := generators[j](ep)
						codeslice := make([]Me3li, 1)
						codeslice[0] = newcode
						newfitness, _ := fitness(codeslice)
						newIndiv.code = &newcode
						fitslice := make([]float32, 1)
						fitslice[0] = newfitness
						newIndiv.fitness_values = fitslice
						orderedPlace(&head, newIndiv)
						currentPopulation++
						generated++
						break
					}
				}
			}
		}
		highestFitness := head.fitness_values[0]
		fmt.Println("Generation: ", generation, " - Population size: ", currentPopulation, " - Removed: ", removed, "- Generated: ", generated, " - Unary: ", unaryApplied, " - Binary: ", binaryApplied, " - Highest: ", highestFitness)

	}

	plan.Populations[0].PopulationHead = head

}
