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
		shape.WriteAttribute(n, 0, strconv.Itoa(n + 1))
	}
}

func TestParse(t *testing.T) {
	generateTestFile("/tmp/points.shp")
	var p [3]string
	count := 0
	openWithHandler("/tmp/points.shp", func (zip string, geocoder []byte) bool {
		//db.Put([]byte(zip), geocoder, nil)
		p[count] = string(geocoder)
		count = count + 1
		return true
	})
	if p[0] != "10.000000:10.000000" {
		t.Error("First element expect 10.00:10.00. Got", p[0])
	}
	if p[1] != "10.000000:15.000000" {
		t.Error("First element expect 10.00:10.00. Got", p[1])
	}
	if p[2] != "15.000000:15.000000" {
		t.Error("First element expect 10.00:10.00. Got", p[2])
	}
}
