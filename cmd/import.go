package cmd

import (
	"reflect"
	"fmt"
	"log"
	"github.com/codegangsta/cli"
	"github.com/jonas-p/go-shp"
)

var CmdImport = cli.Command{
	Name: "import",
	Usage: "import -file /absolute/path/to/shapefile.shp",
	Description: `Import shapefile into GoZip database which is backed by leveldb`,
	Action: runImport,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "file",
			Value: "",
			Usage: "shp file path to import",
		},
	},
}

func runImport(c *cli.Context) {
	var file = c.String("file")
	log.Printf("Importing local repositories...%s", file)
	open(file)
	log.Println("Finish importing!")
}

func open(file string) {
	shape, err := shp.Open(file)
	if err != nil { log.Fatal(err) } 
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
			n, p := shape.Shape()

			// print feature
			fmt.Println(reflect.TypeOf(p).Elem(), p.BBox())

			// print attributes
			for k, f := range fields {
					val := shape.ReadAttribute(n, k)
					fmt.Printf("\t%v: %v\n", f, val)
			}
			fmt.Println()
	}
}
