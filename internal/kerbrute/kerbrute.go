package kerbrute

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BrunoTeixeira1996/gowned/internal/bh"
)



// Function that reads the output
// created the Accounts slice of structs
// and returns that slice
func readOutput(wantOutput bool) []bh.ValidAccount {
	var (
		accounts  []bh.ValidAccount
		rawOutput string
	)

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		if wantOutput {
			fmt.Println(sc.Text())
		}
		rawOutput += sc.Text() + "\n"
		if strings.Contains(sc.Text(), "VALID LOGIN") {
			var acc bh.ValidAccount
			tmp := strings.Split(strings.Split(sc.Text(), " ")[7], ":")
			acc.Email = tmp[0]
			acc.Password = tmp[1]

			accounts = append(accounts, acc)
		}
	}

	return accounts
}

func Execute(wantOutput bool, n bh.Neo4jManager) error {
	accounts := readOutput(wantOutput)

	if accounts != nil {
		for _, user := range accounts {
			err := bh.MarkAsOwned(n, user)
			if err != nil {
				log.Println(err)
			} else {
			log.Printf("Added %s as owned\n", user.Email)
			}
		}
	}

	return nil
}
