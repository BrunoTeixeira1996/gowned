package bh

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4j struct {
	User string
	Pass string
	Port string
	Uri  string
}

// Connect to neo4j
func (n *Neo4j) Connect() error {
	driver, err := neo4j.NewDriver("bolt://"+n.Uri+":"+n.Port, neo4j.BasicAuth(n.User, n.Pass, ""))
	if err != nil {
		return err
	}
	defer driver.Close()

	if err := driver.VerifyConnectivity(); err != nil {
		return err
	}
	log.Println("Connected with neo4j database")

	return nil
}

// https://github.com/byt3bl33d3r/CrackMapExec/blob/master/cme/modules/bh_owned.py

// TODO Recieves user and mark as owned
func MarkAsOwned() {}
