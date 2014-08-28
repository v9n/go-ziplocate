package cmd

import (
	"log"
	"github.com/codegangsta/cli"
	"github.com/jonas-p/go-shp"
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
