package main

import (
	"flag"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/mmirko/mel"
	"github.com/mmirko/mel/lisevo"
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

		l.Init(c, ep, libs)

		if err := l.LoadProgramFromFile(lisevoFile); err != nil {
			log.Fatal("Error while parsing file:", err)
			return
		}
		l.MelDump()
	}

}
