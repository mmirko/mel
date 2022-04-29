package mel

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

// Commands
const (
	MEL_COM_NEXT_EVENT       = iota // Go and to your things whateven they are
	MEL_COM_EXECUTE_OPERATOR        // Execute e genetic operator
	MEL_COM_COMPUTE_FITNESS         // Compute e fitness
)

// Responses
const (
	MEL_RESP_EXECUTE_OPERATOR_DONE = iota // Execution of a genetic operator done
	MEL_RESP_COMPUTE_FITNESS_DONE         // Compute of a fitness function done
	MEL_RESP_EVEN_NEW                     // Trigger a new individual creation
)

type individual struct {
	code           *Me3li
	fitness_values []float32
	prev           *individual
	next           *individual
}

type Population struct {
	Population_name    string
	Population_head    *individual
	Newborn_head       *individual
	Genetic_generators []interface{}
	Genetic_unary      []interface{}
	Genetic_binary     []interface{}
	Weight_generators  []float32
	Weight_unary       []float32
	Weight_binary      []float32
	Weight_death       float32
	Threads            int
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

type Plan_simple struct {
	Plan
	Generation_number int
	Population_size   int
	DeathsPerc        float32
	UnaryPerc         float32
	BinaryPerc        float32
}

type event_issued struct {
	command       int8
	code          *Me3li
	opt_code      *Me3li
	operator_args int
	operator_id   int
}

type fitness_issued struct {
	command int8
	popids  []int
	indivs  []*individual
}

type command_response struct {
	response      int8
	id            int
	popid         int
	code          *Me3li
	opt_code      *Me3li
	value         float32
	operator_args int
	operator_id   int
	popids        []int
	indivs        []*individual
}

type g0 func(*EvolutionParameters) Me3li
type g1 func(Me3li, *EvolutionParameters) Me3li
type g2 func(Me3li, Me3li, *EvolutionParameters) Me3li

// genop_carrier struct, it holds the available genetic operators. It is used to pass function pointers to events goroutines
type genop_carrier struct {
	g0slice []g0
	g1slice []g1
	g2slice []g2
	w0slice []float32
	w1slice []float32
	w2slice []float32
	wdeath  float32
}

type population_rundata struct {
	popn              int
	newpopn           int
	total_fitness     float32
	mean_fitness      float32
	total_fitness_sq  float32
	variance          float32
	threshold_fitness float32
	sync.Mutex
}

type event_channels []chan event_issued
type fitness_channels []chan fitness_issued
type command_responses []command_response

type states []bool

func life(id int, popid int, pop_head **individual, running_data *population_rundata, gen *genop_carrier, ep *EvolutionParameters, comm <-chan event_issued, resp chan<- command_response) {
	//fmt.Println("Starting life thread ", id, " on population ", popid)
	running := true
	pop_target := float32(1000.0)

	popn := &running_data.popn

	for iindiv := *pop_head; running; {
		if iindiv == nil {
			// Only the creation can happen while iindiv is nil and then jump to the pop_head
			for i, v := range gen.w0slice {
				if rand.Float32() < v*pop_target/float32(*popn) {
					resp <- command_response{MEL_RESP_EVEN_NEW, id, popid, nil, nil, 0.0, 0, i, nil, nil}
					<-comm
					break
				}
			}

			iindiv = *pop_head
		} else {
			if iindiv != *pop_head {
				// Individual death
				if iindiv.fitness_values[0] < running_data.threshold_fitness || rand.Float32() < gen.wdeath*float32(*popn)/(pop_target*iindiv.fitness_values[0]) {
					running_data.Lock()
					(*popn)--
					if iindiv.next != nil {
						iindiv.next.prev = iindiv.prev
					}
					iindiv.prev.next = iindiv.next
					//fmt.Println("Death ", iindiv.fitness_values[0])
					//for iii := *pop_head; iii != nil; iii = iii.next {
					//	fmt.Printf("%p (p %p - n %p) - ", iii, iii.prev, iii.next)
					//}
					//fmt.Println("")

					running_data.total_fitness -= iindiv.fitness_values[0]
					running_data.total_fitness_sq -= iindiv.fitness_values[0] * iindiv.fitness_values[0]

					running_data.mean_fitness = running_data.total_fitness / (float32(running_data.popn))
					running_data.variance = (running_data.total_fitness_sq / (float32(running_data.popn))) - (running_data.mean_fitness * running_data.mean_fitness)

					if running_data.mean_fitness <= 1 && running_data.threshold_fitness < running_data.mean_fitness {
						running_data.threshold_fitness = running_data.mean_fitness
					}

					//fmt.Print("Del", *running_data, iindiv.fitness_values[0], running_data.popn," ")
					//somma:=float32(0.0)
					//for jindiv:=*pop_head; jindiv!=nil ; jindiv=jindiv.next {
					//	somma+=jindiv.fitness_values[0]
					//}
					//fmt.Println(somma)
					running_data.Unlock()
					iindiv = iindiv.next
					continue
				}

				// Individual Creation
				for i, v := range gen.w0slice {
					if rand.Float32() < v*pop_target/float32(*popn) {
						resp <- command_response{MEL_RESP_EVEN_NEW, id, popid, nil, nil, 0.0, 0, i, nil, nil}
						<-comm
						break
					}
				}

				// Unary genetic operators
				for i, v := range gen.w1slice {
					if rand.Float32() < v*iindiv.fitness_values[0]*pop_target/float32(*popn) {
						resp <- command_response{MEL_RESP_EVEN_NEW, id, popid, iindiv.code, nil, 0.0, 1, i, nil, nil}
						<-comm
						break
					}
				}

				// Binary genetic operators
				for i, v := range gen.w2slice {
					if rand.Float32() < v*iindiv.fitness_values[0]*pop_target/float32(*popn) {
						resp <- command_response{MEL_RESP_EVEN_NEW, id, popid, iindiv.code, nil, 0.0, 2, i, nil, nil}
						<-comm
						break
					}
				}
			}

			iindiv = iindiv.next

		}
	}
}

func event(id int, popid int, gen *genop_carrier, ep *EvolutionParameters, comm <-chan event_issued, resp chan<- command_response) {
	//fmt.Println("Starting event thread ", id, " on population ", popid)
	for {
		received_command := <-comm
		//fmt.Println("thread ", id, " on population ", popid, " received :", received_command)
		switch received_command.command {
		case MEL_COM_EXECUTE_OPERATOR:
			switch received_command.operator_args {
			case 0:
				genetic_op := gen.g0slice[received_command.operator_id]
				cc := genetic_op(ep)
				//fmt.Println(cc)
				resp <- command_response{MEL_RESP_EXECUTE_OPERATOR_DONE, id, popid, &cc, nil, 0.0, 0, received_command.operator_id, nil, nil}
			case 1:
				genetic_op := gen.g1slice[received_command.operator_id]
				cc := genetic_op(*received_command.code, ep)
				//fmt.Println(cc)
				resp <- command_response{MEL_RESP_EXECUTE_OPERATOR_DONE, id, popid, &cc, nil, 0.0, 1, received_command.operator_id, nil, nil}
			case 2:
				genetic_op := gen.g2slice[received_command.operator_id]
				cc := genetic_op(*received_command.code, *received_command.opt_code, ep)
				//fmt.Println(cc)
				resp <- command_response{MEL_RESP_EXECUTE_OPERATOR_DONE, id, popid, &cc, nil, 0.0, 2, received_command.operator_id, nil, nil}
			}
		}
	}
}

func fittcomp(id int, fitid int, fitness func([]Me3li) (float32, bool), ep *EvolutionParameters, comm <-chan fitness_issued, resp chan<- command_response) {
	//fmt.Println("Starting fitness thread ", id, " on fitness ", fitid)
	for {
		received_command := <-comm
		//fmt.Println("fitthread ", id, " on fitness ", fitid, " received :", received_command)
		melislice := make([]Me3li, len(received_command.indivs))
		for i, v := range received_command.indivs {
			melislice[i] = *v.code
		}
		fitness_value, _ := fitness(melislice)
		resp <- command_response{MEL_RESP_COMPUTE_FITNESS_DONE, id, fitid, nil, nil, fitness_value, 0, 0, received_command.popids, received_command.indivs}
	}
}

// Use Dynamic Evolving populations
func (plan *Plan) Execute_dep(ep *EvolutionParameters) {

	// Evoulution variables
	stopiter := plan.ExitAt
	population_num := len(plan.Populations)
	fitness_num := len(plan.Fitnesses)

	gen := make([]*genop_carrier, population_num)

	var logchans []chan log_entry
	var logverb []int

	log_targets, logging := ep.GetMatchingList("log_target:")

	// Eventually prepare the logging channels and spawn the loggers gothreads
	if logging {
		logchans = make([]chan log_entry, len(log_targets))
		logverb = make([]int, len(log_targets))
		for log_values, _ := range log_targets {
			if log_target_id, ok := GetNthParamsInt(log_values, 0); ok {
				if log_target_verbosity, ok := GetNthParamsInt(log_values, 1); ok {
					logchans[log_target_id] = make(chan log_entry)
					logverb[log_target_id] = log_target_verbosity
					go logger(log_target_id, ep, logchans[log_target_id])
				}
			}
		}
	}

	// Preparing the genetic operator carrier
	if logging {
		logit(log_entry{"Running pre execution tasks", 0}, logverb, logchans)
	}

	// Populate all the genetic carriers
	for pop_i := 0; pop_i < population_num; pop_i++ {
		gen[pop_i] = new(genop_carrier)

		gen[pop_i].g0slice = make([]g0, len(plan.Populations[pop_i].Genetic_generators))
		gen[pop_i].g1slice = make([]g1, len(plan.Populations[pop_i].Genetic_unary))
		gen[pop_i].g2slice = make([]g2, len(plan.Populations[pop_i].Genetic_binary))

		gen[pop_i].w0slice = make([]float32, len(plan.Populations[pop_i].Weight_generators))
		gen[pop_i].w1slice = make([]float32, len(plan.Populations[pop_i].Weight_unary))
		gen[pop_i].w2slice = make([]float32, len(plan.Populations[pop_i].Weight_binary))

		gen[pop_i].wdeath = plan.Populations[pop_i].Weight_death

		for genop_i := 0; genop_i < len(plan.Populations[pop_i].Genetic_generators); genop_i++ {
			gen[pop_i].g0slice[genop_i] = plan.Populations[pop_i].Genetic_generators[genop_i].(func(*EvolutionParameters) Me3li)
			gen[pop_i].w0slice[genop_i] = plan.Populations[pop_i].Weight_generators[genop_i]
		}

		for genop_i := 0; genop_i < len(plan.Populations[pop_i].Genetic_unary); genop_i++ {
			gen[pop_i].g1slice[genop_i] = plan.Populations[pop_i].Genetic_unary[genop_i].(func(Me3li, *EvolutionParameters) Me3li)
			gen[pop_i].w1slice[genop_i] = plan.Populations[pop_i].Weight_unary[genop_i]
		}

		for genop_i := 0; genop_i < len(plan.Populations[pop_i].Genetic_binary); genop_i++ {
			gen[pop_i].g2slice[genop_i] = plan.Populations[pop_i].Genetic_binary[genop_i].(func(Me3li, Me3li, *EvolutionParameters) Me3li)
			gen[pop_i].w2slice[genop_i] = plan.Populations[pop_i].Weight_binary[genop_i]
		}
	}

	// Preparing channels
	life_channels := make([]chan event_issued, population_num)

	waiting_events := make([]command_responses, population_num)
	waiting_events_i := make([]int, population_num)

	event_thread_free := make([]int, population_num)
	event_thread_comchan := make([]event_channels, population_num)
	event_thread_status := make([]states, population_num)

	for i := 0; i < population_num; i++ {
		life_channels[i] = make(chan event_issued)

		event_thread_free[i] = plan.Populations[i].Threads

		event_thread_comchan[i] = make([]chan event_issued, plan.Populations[i].Threads)

		event_thread_status[i] = make([]bool, plan.Populations[i].Threads)
		for j := 0; j < plan.Populations[i].Threads; j++ {
			event_thread_comchan[i][j] = make(chan event_issued)
			event_thread_status[i][j] = true

		}

		waiting_events[i] = make([]command_response, plan.Populations[i].Threads)
		waiting_events_i[i] = 0
	}

	fitness_thread_free := make([]int, fitness_num)
	fitness_thread_comchan := make([]fitness_channels, fitness_num)
	fitness_thread_status := make([]states, fitness_num)
	for i := 0; i < fitness_num; i++ {
		fitness_thread_free[i] = plan.Fitnesses[i].Threads

		fitness_thread_comchan[i] = make([]chan fitness_issued, plan.Fitnesses[i].Threads)

		fitness_thread_status[i] = make([]bool, plan.Fitnesses[i].Threads)
		for j := 0; j < plan.Fitnesses[i].Threads; j++ {
			fitness_thread_comchan[i][j] = make(chan fitness_issued)
			fitness_thread_status[i][j] = true
		}
	}

	responses_channel := make(chan command_response)

	newpops := make([]int, population_num)

	running_data := make([]population_rundata, population_num)

	// Spawn the goroutines
	for i := 0; i < population_num; i++ {
		// Spawn the life goroutine for evey population
		go life(i, i, &plan.Populations[i].Population_head, &running_data[i], gen[i], ep, life_channels[i], responses_channel)

		// Spawn n-goroutines for every population
		for j := 0; j < plan.Populations[i].Threads; j++ {
			go event(j, i, gen[i], ep, event_thread_comchan[i][j], responses_channel)
		}

	}

	// Spawn the fittnes calcutors goroutines
	for i := 0; i < fitness_num; i++ {
		for j := 0; j < plan.Fitnesses[i].Threads; j++ {
			go fittcomp(j, i, plan.Fitnesses[i].FitnessFunction, ep, fitness_thread_comchan[i][j], responses_channel)
		}
	}

	for tot_ind := 0; tot_ind < stopiter; {
		select {
		case stat := <-responses_channel:
			//fmt.Println(stat)
			switch stat.response {
			// A new event has happened, it is inserted into the waiting queue for the event population.
			case MEL_RESP_EVEN_NEW:
				waiting_events[stat.popid][waiting_events_i[stat.popid]] = stat
				waiting_events_i[stat.popid]++

				if logging {
					logit(log_entry{"New event triggered", 0}, logverb, logchans)
				}

			// An operator has terminated execution and a new individual has been created.
			case MEL_RESP_EXECUTE_OPERATOR_DONE:
				new_head := new(individual)
				new_head.fitness_values = make([]float32, population_num)
				new_head.code = stat.code
				new_head.prev = nil
				new_head.next = plan.Populations[stat.popid].Newborn_head
				if plan.Populations[stat.popid].Newborn_head != nil {
					plan.Populations[stat.popid].Newborn_head.prev = new_head
				}
				plan.Populations[stat.popid].Newborn_head = new_head
				newpops[stat.popid]++

				event_thread_status[stat.popid][stat.id] = true
				event_thread_free[stat.popid]++

				if logging {
					logit(log_entry{"Genetic operator executed", 0}, logverb, logchans)
				}

			// A Fittness has been computed
			case MEL_RESP_COMPUTE_FITNESS_DONE:
				for res_i, indiv := range stat.indivs {
					// stat.popid here actually is a fitid
					indiv.fitness_values[stat.popid] = stat.value

					complete_check := true
					for _, valfitt := range indiv.fitness_values {
						if valfitt < 0 {
							complete_check = false
						}
					}

					if complete_check {
						new_head := indiv
						new_head.prev = nil
						new_head.next = plan.Populations[stat.popids[res_i]].Population_head
						if plan.Populations[stat.popids[res_i]].Population_head != nil {
							plan.Populations[stat.popids[res_i]].Population_head.prev = new_head
						}
						plan.Populations[stat.popids[res_i]].Population_head = new_head
						tot_ind++

						running_data[stat.popids[res_i]].Lock()
						running_data[stat.popids[res_i]].popn++

						// Every 1000 recount fitnesses to avoid the accomulation of error
						if tot_ind%1000 == 0 {
							running_data[stat.popids[res_i]].total_fitness = 0
							running_data[stat.popids[res_i]].total_fitness_sq = 0
							for jindiv := plan.Populations[stat.popids[res_i]].Population_head; jindiv != nil; jindiv = jindiv.next {
								running_data[stat.popids[res_i]].total_fitness += jindiv.fitness_values[stat.popid]
								running_data[stat.popids[res_i]].total_fitness_sq += jindiv.fitness_values[stat.popid] * jindiv.fitness_values[stat.popid]
							}
						}
						running_data[stat.popids[res_i]].total_fitness += stat.value
						running_data[stat.popids[res_i]].total_fitness_sq += stat.value * stat.value

						running_data[stat.popids[res_i]].mean_fitness = running_data[stat.popids[res_i]].total_fitness / (float32(running_data[stat.popids[res_i]].popn))
						running_data[stat.popids[res_i]].variance = (running_data[stat.popids[res_i]].total_fitness_sq / (float32(running_data[stat.popids[res_i]].popn))) - (running_data[stat.popids[res_i]].mean_fitness * running_data[stat.popids[res_i]].mean_fitness)

						if running_data[stat.popids[res_i]].mean_fitness <= 1 && running_data[stat.popids[res_i]].mean_fitness > running_data[stat.popids[res_i]].threshold_fitness {
							running_data[stat.popids[res_i]].threshold_fitness = running_data[stat.popids[res_i]].mean_fitness
						}

						//fmt.Println("Add", running_data[stat.popids[res_i]], stat.value)

						running_data[stat.popids[res_i]].Unlock()

					}
				}

				fitness_thread_status[stat.popid][stat.id] = true
				fitness_thread_free[stat.popid]++

				if logging {
					logit(log_entry{"Fitness (id " + logi(stat.popid) + ") computed with value " + logf(stat.value), 0}, logverb, logchans)
				}
			}
		}

		//fmt.Println("Total individual: ", tot_ind, " - Population: ", pops[0], " - Newborn: ", newpops[0], " - Free compute threads: ", event_thread_free[0], " - Free fittness threads: ", fitness_thread_free[0], " - Remainig events: ", waiting_events_i[0])

		// Spawn genetic operators
		for i := 0; i < population_num; i++ {
			if event_thread_free[i] > 0 && newpops[i] < 10 {
				for j, tstat := range event_thread_status[i] {
					if tstat {
						if waiting_events_i[i] > 0 {
							event_i := waiting_events_i[i] - 1
							event_thread_comchan[i][j] <- event_issued{MEL_COM_EXECUTE_OPERATOR, waiting_events[i][event_i].code, waiting_events[i][event_i].opt_code, waiting_events[i][event_i].operator_args, waiting_events[i][event_i].operator_id}
							waiting_events_i[i]--
							event_thread_status[i][j] = false
							event_thread_free[i]--
							life_channels[i] <- event_issued{MEL_COM_NEXT_EVENT, nil, nil, 0, 0}
						}
					}
				}
			}
		}

		// Spawn fittness calculation (THIS IS VERY TEMPORARY CODE)
		for i := 0; i < fitness_num; i++ {
			if fitness_thread_free[i] > 0 {
				for j, tstat := range fitness_thread_status[i] {
					if tstat {
						for popid := 0; popid < population_num; popid++ {
							if plan.Populations[popid].Newborn_head != nil {
								indivs := make([]*individual, 1)
								popids := make([]int, 1)
								indivs[0] = plan.Populations[popid].Newborn_head
								popids[0] = popid
								plan.Populations[popid].Newborn_head = plan.Populations[popid].Newborn_head.next
								fitness_thread_comchan[i][j] <- fitness_issued{MEL_COM_COMPUTE_FITNESS, popids, indivs}
								fitness_thread_status[i][j] = false
								fitness_thread_free[i]--
								newpops[popid]--
								break
							}
						}
					}
				}
			}
		}
	}
}

func ordered_place(queue_head **individual, newindiv *individual) {
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

// Use Simple Genetic evolution
func (plan *Plan_simple) Execute_simple(ep *EvolutionParameters) {

	var head *individual

	// Evoulution variables
	gennum := plan.Generation_number
	popsize := plan.Population_size
	population_num := len(plan.Populations)
	fitness_num := len(plan.Fitnesses)

	if population_num != 1 {
		log.Fatal("Simple Plan has to have 1 population")
	}

	if fitness_num != 1 {
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
	deathsPerc := plan.DeathsPerc
	unaryPerc := plan.UnaryPerc
	binaryPerc := plan.BinaryPerc

	// Main loop: cicle throught generations
	for generation := 0; generation < gennum; generation++ {

		generated := 0
		unary_applied := 0
		binary_applied := 0

		// Remove unfitted individuals
		removed := int(float32(current_population) * deathsPerc)
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
			howmany_unary = int(float32(current_population) * unaryPerc)

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
			howmany_binary = int(float32(current_population) * binaryPerc)

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
			unary_applied++
		}

		for i := 0; i < howmany_binary && current_population < popsize; i++ {
			ordered_place(&head, binary_container[i])
			current_population++
			binary_applied++
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
		fmt.Println("Generation: ", generation, " - Population size: ", current_population, " - Removed: ", removed, "- Generated: ", generated, " - Unary: ", unary_applied, " - Binary: ", binary_applied, " - Highest: ", highestFitness)

	}

	plan.Populations[0].Population_head = head

}

func (plan *Plan) Get_best() (*Me3li, float32) {
	result := plan.Populations[0].Population_head.code
	resultfit := plan.Populations[0].Population_head.fitness_values[0]
	return result, resultfit
}
