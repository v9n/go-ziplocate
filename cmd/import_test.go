package cmd

import (
	"log"
	//"github.com/codegangsta/cli"
	"testing"
	"github.com/jonas-p/go-shp"
	"strconv"
)

func TestDef(t *testing.T) {
	if CmdImport.Name != "import" {
		t.Fatal("Wrong Command Name")
	}
}

func generateTestFile(file string) {
	// points to write
	points := []shp.Point{
		shp.Point{10.0, 10.0},
		shp.Point{10.0, 15.0},
		shp.Point{15.0, 15.0},
		shp.Point{15.0, 10.0},
	}

	// fields to write
	fields := []shp.Field{
		// String attribute field with length 25
		shp.StringField("ZIP", 25),
	}

	// create and open a shapefile for writing points
	shape, err := shp.Create(file, shp.POINT)
	if err != nil { log.Fatal(err) }
	defer shape.Close()

	// setup fields for attributes
	shape.SetFields(fields)

	// write points and attributes
	for n, point := range points {
		shape.Write(&point)

		// write attribute for object n for field 0 (ZIP)
		shape.WriteAttribute(n, 0, "Point " + strconv.Itoa(n + 1))
	}
}

func TestParse(t *testing.T) {
	generateTestFile("/tmp/points.shp")

}
