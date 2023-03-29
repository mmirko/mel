package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mmirko/mel/pkg/conconi"
	"github.com/mmirko/mel/pkg/mel"
)

var datafile = flag.String("datafile", "datafile.csv", "CSV Datafile")
var gnuplotfile = flag.String("gnuplotfile", "", "gnuplot output file")

var x []float32
var y []float32

var minX float32
var maxX float32

var ep *mel.EvolutionParameters

func importTB(filename string) bool {

	minX = 100.0
	maxX = 0.0

	x = make([]float32, 16)
	y = make([]float32, 16)

	fi, err := os.Open(filename)
	if err != nil {
		return true
	}

	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	j := 0
	for scanner.Scan() {
		varval := strings.Split(scanner.Text(), " ")
		x_val, _ := strconv.ParseFloat(varval[0], 32)
		y_val, _ := strconv.ParseFloat(varval[1], 32)
		x[j] = float32(x_val)
		y[j] = float32(y_val)
		if x[j] < minX {
			minX = x[j]
		}
		if x[j] > maxX {
			maxX = x[j]
		}
		j++
	}

	return false
}

func FitnessFunction(toCheck []mel.Me3li) (float32, bool) {

	program := toCheck[0].(*conconi.Conconi)

	if result, ok := conconi.ConconiFitness(program, x, y); ok {
		return result, true
	}

	return 0.0, false
}

func main() {
	rand.Seed(int64(time.Now().Unix()))

	flag.Parse()

	err := importTB(*datafile)
	if err {
		log.Fatal("Data file error")
	}

	ep = new(mel.EvolutionParameters)
	ep.Pars = make(map[string]string)

	generators := make([]interface{}, 2)
	var gen func(*mel.EvolutionParameters) mel.Me3li
	gen = conconi.ConconiGenerate
	generators[0] = gen
	generators[1] = gen

	unary := make([]interface{}, 2)
	var un1 func(mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	var un2 func(mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	un1 = conconi.ConconiMutate
	un2 = conconi.ConconiMutateSlow
	unary[0] = un1
	unary[1] = un2

	binary := make([]interface{}, 2)
	var cr1 func(mel.Me3li, mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	var cr2 func(mel.Me3li, mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	cr1 = conconi.ConconiCrossoverFake
	cr2 = conconi.ConconiCrossoverFake
	binary[0] = cr1
	binary[1] = cr2

	wGenerators := make([]float32, 2)
	wGenerators[0] = 0.1
	wGenerators[1] = 0.1

	wUnary := make([]float32, 2)
	wUnary[0] = 1
	wUnary[1] = 2

	wBinary := make([]float32, 2)
	wBinary[0] = 1
	wBinary[1] = 2

	myPop := make([]mel.Population, 1)
	myPop[0].PopulationHead = nil
	myPop[0].NewbornHead = nil
	myPop[0].GeneticGenerators = generators
	myPop[0].GeneticUnary = unary
	myPop[0].GeneticBinary = binary
	myPop[0].WeightGenerators = wGenerators
	myPop[0].WeightUnary = wUnary
	myPop[0].WeightBinary = wBinary
	myPop[0].WeightDeath = 0.01
	myPop[0].Threads = 2

	myFit := make([]mel.Fitness, 1)
	myFit[0].FitnessFunction = FitnessFunction
	myFit[0].Threads = 5

	//ep.Pars["log_target:0:0"] = "stdout"
	//ep.Pars["log_target:1:0"] = "/tmp/prova"

	//ep.Pars["symbolic_math:const:alt:range_int:-10_10"] = "1"

	myPlan := new(mel.Plan)
	myPlan.Populations = myPop
	myPlan.Fitnesses = myFit
	myPlan.ExitAt = 10000

	mySimplePlan := new(mel.PlanSimple)
	mySimplePlan.GenerationNumber = 1000
	mySimplePlan.PopulationSize = 1000
	mySimplePlan.Plan = *myPlan
	mySimplePlan.DeathsRate = 0.5
	mySimplePlan.UnaryRate = 0.5
	mySimplePlan.BinaryRate = 0.25

	mySimplePlan.Execute(ep)

	best, value := mySimplePlan.GetBest()

	//	fmt.Println(*best,value)
	c := (*best).(*conconi.Conconi)
	get_x, get_fca := c.GetAT()
	m1, k1, m2, k2 := c.GetParams()

	fmt.Println("Fitness:", value, "- FCA:", get_fca, "Bpm - Speed:", get_x, "Km/h")

	if *gnuplotfile != "" {

		fo, err := os.Create(*gnuplotfile)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		gnuplotstring := ""
		gnuplotstring = gnuplotstring + fmt.Sprintf("#!/usr/bin/env gnuplot\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("set terminal png size 1024,768 enhanced font \"Helvetica,20\"\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("set output 'output.png'\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("set key outside\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("set xrange [%f-0.2:%f+0.2]\n", minX, maxX)
		gnuplotstring = gnuplotstring + fmt.Sprintf("g(x,min,max)=( (x>=min && x<=max)? 1.0 : (1/0) )\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("plot (%f*x+%f)*g(x,%f-0.2,%f+0.2) title \"aerobic\", (%f*x+ %f)*g(x,%f-0.2,%f+0.2) title \"anaerobic\" , \"%s\" title \"data\"\n", m1, k1, minX, get_x, m2, k2, get_x, maxX, *datafile)

		if _, err := fo.Write([]byte(gnuplotstring)); err != nil {
			panic(err)
		}
	}
}
