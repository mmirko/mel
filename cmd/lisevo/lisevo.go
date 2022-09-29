package main

import (
	"flag"
	"log"
	"math/rand"
	"time"
)

var debug = flag.Bool("debug", false, "debug mode")
var program = flag.String("program", "", "program to run")
var libraries = flag.String("libraries", "", "libraries to use")

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
}

func main() {
	if *program == "" {
		log.Fatal("No program specified")
	}
}
