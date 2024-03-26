package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BrunoTeixeira1996/gowned/internal/bh"
	"github.com/BrunoTeixeira1996/gowned/internal/kerbrute"
)

/*
   kerbrute ... | gowned -source="kerbrute" output
*/

func run() error {
	var sourceFlag = flag.String("source", "", "use this to provide the source")
	var outputFlag = flag.Bool("output", false, "use this if you want to see kerbrute output")
	var neo4jUser = flag.String("neo4juser", "neo4j", "use this to provide the user of neo4j")
	var neo4jPass = flag.String("neo4jpass", "neo4j", "use this to provide the password of neo4j")
	var neo4jPort = flag.String("neo4jport", "7687", "use this to provide the port of neo4j")
	var neo4jUri = flag.String("neo4juri", "127.0.0.1", "use this to provide the uri of neo4j")

	flag.Parse()

	neo4j := bh.Neo4j{
		User: *neo4jUser,
		Pass: *neo4jPass,
		Port: *neo4jPort,
		Uri:  *neo4jUri,
	}

	if err := neo4j.Connect(); err != nil {
		return fmt.Errorf("Error while connecting to neo4j: %s", err)
	}

	switch *sourceFlag {
	case "kerbrute":
		_ = kerbrute.Execute(*outputFlag, neo4j)
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
