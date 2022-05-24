package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"

	"github.com/mmirko/mel"
	"github.com/mmirko/mel/rectangular"
)

var imageFile = flag.String("imagefile", "", "Target image file")
var imageTarget *image.Image

var ep *mel.EvolutionParameters

func FitnessImageDistance(toCheck []mel.Me3li) (float32, bool) {

	me3li := toCheck[0].(*rectangular.RectangularMe3li)

	if result, ok := rectangular.FitnessImageDistance(me3li, imageTarget, ep); ok {
		return result, true
	}

	return 0.0, false
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil

}

func main() {
	rand.Seed(int64(time.Now().Unix()))

	flag.Parse()

	if *imageFile == "" {
		fmt.Println("Please specify an image file")
		return
	}

	imageI, err := getImageFromFilePath(*imageFile)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return
	} else {
		imageTarget = &imageI
	}

	ep = new(mel.EvolutionParameters)

	// Get the width and height of the image
	width := (*imageTarget).Bounds().Max.X
	height := (*imageTarget).Bounds().Max.Y

	ep.SetValue("width", fmt.Sprintf("%d", width))
	ep.SetValue("height", fmt.Sprintf("%d", height))

	generators := make([]interface{}, 1)
	var gen func(*mel.EvolutionParameters) mel.Me3li
	gen = rectangular.Generate
	generators[0] = gen

	unary := make([]interface{}, 1)
	var un1 func(mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	un1 = rectangular.Mutate
	unary[0] = un1

	binary := make([]interface{}, 1)
	var cr1 func(mel.Me3li, mel.Me3li, *mel.EvolutionParameters) mel.Me3li
	cr1 = rectangular.Crossover
	binary[0] = cr1

	wGenerators := make([]float32, 1)
	wGenerators[0] = 1

	wUnary := make([]float32, 1)
	wUnary[0] = 1

	wBinary := make([]float32, 1)
	wBinary[0] = 1

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
	myPop[0].Threads = 1

	myFit := make([]mel.Fitness, 1)
	myFit[0].FitnessFunction = FitnessImageDistance
	myFit[0].Threads = 1

	//ep.Pars["log_target:0:0"] = "stdout"
	//ep.Pars["log_target:1:0"] = "/tmp/prova"

	//ep.Pars["symbolic_math:const:alt:range_int:-10_10"] = "1"

	myPlan := new(mel.Plan)
	myPlan.Populations = myPop
	myPlan.Fitnesses = myFit
	myPlan.ExitAt = 10000

	myInfo := new(mel.RunInfo)

	mySecPlan := new(mel.PlanBasic)
	mySecPlan.GenerationNumber = 100
	mySecPlan.PopulationSize = 10
	mySecPlan.Plan = *myPlan
	mySecPlan.DeathsRate = 0.5
	mySecPlan.UnaryRate = 0.5
	mySecPlan.BinaryRate = 0.25
	mySecPlan.RunInfo = *myInfo

	mySecPlan.Execute(ep)

	best, value := mySecPlan.GetBest()

	fmt.Println(*best, value)

	// Write the image to a file
	file, err := os.Create("out.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	outImage, err := (*best).(*rectangular.RectangularMe3li).ToImage(ep)
	err = png.Encode(file, outImage)

	if err != nil {
		fmt.Println("Error writing image:", err)
		return
	}
}
