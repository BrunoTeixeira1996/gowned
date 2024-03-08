package main

import (
	"flag"
	"log"

	"github.com/BrunoTeixeira1996/gowned/internal/kerbrute"
)

/*
   kerbrute ... | gowned -source="kerbrute" output
*/

func run() error {
	var sourceFlag = flag.String("source", "", "use this to provide the source")
	var outputFlag = flag.Bool("output", false, "use this if you want to see kerbrute output")

	flag.Parse()

	switch *sourceFlag {
	case "kerbrute":
		_ = kerbrute.Execute(*outputFlag)
	default:
		log.Println("not a valid source")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
