package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BrunoTeixeira1996/gowned/internal/bh"
	"github.com/BrunoTeixeira1996/gowned/internal/kerbrute"
)

/*
	kerbrute passwordspray users_test.txt  --dc 192.168.30.51 --domain MARVEL.local 'P@ssword321' | go run cmd/gowned/main.go -source 'kerbrute' -output
*/

func run() error {
	var sourceFlag = flag.String("source", "", "use this to provide the source")
	var outputFlag = flag.Bool("output", false, "use this if you want to see kerbrute output")
	var neo4jUser = flag.String("neo4juser", "neo4j", "use this to provide the user of neo4j")
	var neo4jPass = flag.String("neo4jpass", "neo4j", "use this to provide the password of neo4j")
	var neo4jPort = flag.String("neo4jport", "7687", "use this to provide the port of neo4j")
	var neo4jUri = flag.String("neo4juri", "127.0.0.1", "use this to provide the uri of neo4j")

	flag.Parse()

	// Starts neo4j connection with a drive
	neo4jManagerTemp, err := bh.NewNeo4jManager(*neo4jUri, *neo4jUser, *neo4jPass, *neo4jPort)
    if err != nil {
        fmt.Errorf("Error initializing Neo4j connection: %v", err)
    }


    neo4jManager := bh.Neo4jManager{
    	User: *neo4jUser,
		Pass: *neo4jPass,
		Port: *neo4jPort,
		Uri:  *neo4jUri,
    	Driver: neo4jManagerTemp.Driver,
    }

   defer neo4jManager.Driver.Close()

	switch *sourceFlag {
	case "kerbrute":
		_ = kerbrute.Execute(*outputFlag, neo4jManager)
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
