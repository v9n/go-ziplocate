package cmd

import (
	"reflect"
	"fmt"
	"log"
	"github.com/codegangsta/cli"
	"github.com/jonas-p/go-shp"
	//"code.google.com/p/leveldb-go/leveldb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/kureikain/go-ziplocate/lib"
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
	
	db, err := leveldb.OpenFile("zipdata", nil)
	if err != nil {
		log.Fatal("Cannot create database")
	}

	defer db.Close()

	openWithHandler(file, func (zip string, geocoder []byte) bool {
		db.Put([]byte(zip), geocoder, nil)
		return true
	})
	log.Println("Finish importing!")
}

func openWithHandler(file string, handler func(zip string, geocoder []byte) bool) {
	shape, err := shp.Open(file)
	if err != nil { log.Fatal(err) } 
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	var centroid zip.Point
	var boundary shp.Box
	for shape.Next() {
			n, p := shape.Shape()
			boundary = p.BBox()
			fmt.Println(reflect.TypeOf(p).Elem(), boundary)
			centroid.X = (boundary.MinX + boundary.MaxX) / 2
			centroid.Y = (boundary.MinY + boundary.MaxY) / 2
			fmt.Println(centroid)
			//This is a naive way to convert struct to string to byte. Probably http://golang.org/pkg/encoding/gob/ is better
			handler(shape.ReadAttribute(n, 0), []byte(fmt.Sprintf("%f:%f", centroid.X, centroid.Y)))
			for k, f := range fields {
					val := shape.ReadAttribute(n, k)
					fmt.Printf("\t%v: %v\n", f, val)
			}
	}
}
