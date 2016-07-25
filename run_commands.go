package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

func runListIssues(c *cli.Context) error {

	if c.NArg() < 1 {
		return cli.NewExitError("No arguments given", 10)
	}

	var terms []string

	// get the name of the repository from the argument
	repo := c.Args().Get(0)

	// get the value of the flag state
	state := c.String("state")

	// append the flags and arguments in terms
	terms = append(terms, "repo:"+repo)

	if state != "" {
		terms = append(terms, "state:"+state)
	}

	result, err := listIssues(terms)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	return nil
}
