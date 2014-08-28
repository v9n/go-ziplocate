package cmd

import (
	"log"
	"github.com/codegangsta/cli"
)

var CmdWeb = cli.Command{
	Name: "web",
	Usage: "web -p port",
	Description: `Run API server on the specify port. default to 12385`,
	Action: runWeb,
	Flags: []cli.Flag{},
}

func runWeb(*cli.Context) {
	log.Printf("Web runs on Port %d", 12385)

	log.Println("Finish webing!")
}


