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

	"github.com/mmirko/mel"
	"github.com/mmirko/mel/conconi"
)

var datafile = flag.String("datafile", "datafile.csv", "CSV Datafile")
var gnuplotfile = flag.String("gnuplotfile", "", "gnuplot output file")

var x []float32
var y []float32

var minx float32
var maxx float32

var ep *mel.EvolutionParameters

func import_tb(filename string) bool {

	minx = 100.0
	maxx = 0.0

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
		if x[j] < minx {
			minx = x[j]
		}
		if x[j] > maxx {
			maxx = x[j]
		}
		j++
	}

	return false
}

func Function_fit_fitness(tocheck []mel.Me3li) (float32, bool) {

	program := tocheck[0].(*conconi.Conconi)

	if result, ok := conconi.ConconiFitness(program, x, y); ok {
		return result, true
	}

	return 0.0, false
}

func main() {
	rand.Seed(int64(time.Now().Unix()))

	flag.Parse()

	err := import_tb(*datafile)
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

	wgenerators := make([]float32, 2)
	wgenerators[0] = 0.1
	wgenerators[1] = 0.1

	wunary := make([]float32, 2)
	wunary[0] = 1
	wunary[1] = 2

	wbinary := make([]float32, 2)
	wbinary[0] = 1
	wbinary[1] = 2

	mypop := make([]mel.Population, 1)
	mypop[0].Population_head = nil
	mypop[0].Newborn_head = nil
	mypop[0].Genetic_generators = generators
	mypop[0].Genetic_unary = unary
	mypop[0].Genetic_binary = binary
	mypop[0].Weight_generators = wgenerators
	mypop[0].Weight_unary = wunary
	mypop[0].Weight_binary = wbinary
	mypop[0].Weight_death = 0.01
	mypop[0].Threads = 2

	myfit := make([]mel.Fitness, 1)
	myfit[0].FitnessFunction = Function_fit_fitness
	myfit[0].Threads = 5

	//ep.Pars["log_target:0:0"] = "stdout"
	//ep.Pars["log_target:1:0"] = "/tmp/prova"

	//ep.Pars["symbolic_math:const:alt:range_int:-10_10"] = "1"

	myplan := new(mel.Plan)
	myplan.Populations = mypop
	myplan.Fitnesses = myfit
	myplan.Exitat = 10000

	mysimpleplan := new(mel.Plan_simple)
	mysimpleplan.Generation_number = 1000
	mysimpleplan.Population_size = 1000
	mysimpleplan.Plan = *myplan
	mysimpleplan.Deaths_perc = 0.5
	mysimpleplan.Unary_perc = 0.5
	mysimpleplan.Binary_perc = 0.25

	mysimpleplan.Execute_simple(ep)

	best, value := mysimpleplan.Get_best()

	//	fmt.Println(*best,value)
	c := (*best).(*conconi.Conconi)
	get_x, get_fca := c.Get_AT()
	m1, k1, m2, k2 := c.Get_params()

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
		gnuplotstring = gnuplotstring + fmt.Sprintf("set xrange [%f-0.2:%f+0.2]\n", minx, maxx)
		gnuplotstring = gnuplotstring + fmt.Sprintf("g(x,min,max)=( (x>=min && x<=max)? 1.0 : (1/0) )\n")
		gnuplotstring = gnuplotstring + fmt.Sprintf("plot (%f*x+%f)*g(x,%f-0.2,%f+0.2) title \"aerobic\", (%f*x+ %f)*g(x,%f-0.2,%f+0.2) title \"anaerobic\" , \"%s\" title \"data\"\n", m1, k1, minx, get_x, m2, k2, get_x, maxx, *datafile)

		if _, err := fo.Write([]byte(gnuplotstring)); err != nil {
			panic(err)
		}
	}
}
