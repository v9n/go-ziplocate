package cmd

import (
	"log"
	"github.com/codegangsta/cli"
)

var CmdImport = cli.Command{
	Name: "import",
	Usage: "gozip -f shapefile",
	Description: `Import shapefile into GoZip database which is backed by leveldb`,
	Action: runImport,
	Flags: []cli.Flag{},
}

func runImport(*cli.Context) {
	log.Printf("Importing local repositories...%s", "tl_2014_us_zcta510.shp")

	log.Println("Finish importing!")
}

