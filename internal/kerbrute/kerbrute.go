package kerbrute

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BrunoTeixeira1996/gowned/internal/bh"
)

type Account struct {
	Email    string
	Password string
}

// Function that reads the output
// created the Accounts slice of structs
// and returns that slice
func readOutput(wantOutput bool) []Account {
	var (
		accounts  []Account
		rawOutput string
	)

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		if wantOutput {
			fmt.Println(sc.Text())
		}
		rawOutput += sc.Text() + "\n"
		if strings.Contains(sc.Text(), "VALID LOGIN") {
			var acc Account
			tmp := strings.Split(strings.Split(sc.Text(), " ")[7], ":")
			acc.Email = tmp[0]
			acc.Password = tmp[1]

			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func Execute(wantOutput bool, neo4j bh.Neo4j) error {
	accounts := readOutput(wantOutput)

	if accounts != nil {
		for _, v := range accounts {
			log.Printf("Adding %s as owned\n", v.Email)
		}
	}

	return nil
}
