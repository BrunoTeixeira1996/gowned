package bh

import (
	"log"
	"time"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

)

type Neo4jManager struct {
	User string
	Pass string
	Port string
	Uri  string
	Domain string
	Driver neo4j.Driver
}

type ValidAccount struct {
	Email    string
	Password string
}

// Connect to neo4j
func NewNeo4jManager(uri, username, password, port string) (*Neo4jManager, error) {
    driver, err := neo4j.NewDriver("bolt://"+uri+":"+port, neo4j.BasicAuth(username, password, ""))
    if err != nil {
        return nil, err
    }
    return &Neo4jManager{Driver: driver}, nil
}

// Closes neo4j driver
func (n *Neo4jManager) Close() {
    if n.Driver != nil {
        if err := n.Driver.Close(); err != nil {
            log.Printf("Error closing Neo4j connection: %v\n", err)
        }
    }
}

// Mark user as owned in neo4j
func MarkAsOwned(n Neo4jManager, user ValidAccount) error {
	session := n.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close()
	
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		q := "MATCH (H:User {name:$user}) SET H.owned=true RETURN H.name"
		params := map[string]interface{}{"user": strings.ToUpper(user.Email)}
		result, err := tx.Run(q, params)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			list = append(list, result.Record().Values[0].(string))
		}

		if err := result.Err(); err != nil {
			return nil, err
		}

		return list, nil 
	}, neo4j.WithTxTimeout(3*time.Second))

	return err
}
