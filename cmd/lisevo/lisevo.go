package main

import (
	"flag"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/mmirko/mel/pkg/lisevo"
	"github.com/mmirko/mel/pkg/mel"
)

var debug = flag.Bool("debug", false, "debug mode")
var libraries = flag.String("libraries", "m3uint", "libraries to use")

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {

	libs := strings.Split(*libraries, ",")

	for _, lisevoFile := range flag.Args() {

		l := new(lisevo.LisevoMe3li)
		var ep *mel.EvolutionParameters
		c := new(mel.MelConfig)
		c.Debug = true

		if err := l.Init(c, ep, libs); err != nil {
			log.Fatal(err)
		}

		if err := l.LoadProgramFromFile(lisevoFile); err != nil {
			log.Fatal("Error while parsing file:", err)
		}
		l.MelDump()
	}

}
