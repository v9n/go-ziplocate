package cmd

import (
	"strings"
	"strconv"
	"fmt"
	"log"
	//"bytes"
	//"encoding/binary"

	"github.com/codegangsta/cli"
	"github.com/syndtr/goleveldb/leveldb"
	zipi "github.com/kureikain/go-ziplocate/lib"
)

var CmdWeb = cli.Command{
	Name: "web",
	Usage: "web -p port",
	Description: `Run API server on the specify port. default to 12385`,
	Action: runWeb,
	Flags: []cli.Flag{},
}

func runWeb(c *cli.Context) {
	port := c.Int("p")
	log.Printf("Web runs on Port %d", port)

	db, err := leveldb.OpenFile("zipdata", nil)
	if err != nil {
		log.Fatal("Db not found")
	}
	defer db.Close()
	zip := "95111"
	fetchZip(db, zip)
	log.Println("Finish webing!")
	fetchAllZip(db)
}

func fetchZip(db *leveldb.DB, zip string) *zipi.Point {
	//geocoder := new(zip.Point)
	geo, err := db.Get([]byte(zip), nil)
	if err != nil {
		log.Fatal("Canot found the zip code")
	}
	fmt.Printf("Geo Byte Array: %v", geo)

	fmt.Printf("Geo %v+", string(geo))
	//buffer := bytes.NewReader(geo)
	//err = binary.Read(buffer, binary.LittleEndian, &geocoder)
	//if err != nil {
		//fmt.Println("binary.Read failed:", err)
	//}
	part := strings.Split(string(geo), ":")
	x,_ := strconv.ParseFloat(part[0], 64)
	y,_ := strconv.ParseFloat(part[1], 64)
	geocoder := new(zipi.Point)
	geocoder.X = x
	geocoder.Y = y

	fmt.Printf("Geo Coder: x= %v, y=%v", geocoder.X, geocoder.Y)
	return geocoder
}

func fetchAllZip(db *leveldb.DB) {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
			// Remember that the contents of the returned slice should not be modified, and
			// only valid until the next call to Next.
			key := iter.Key()
			value := iter.Value()
			fmt.Printf("\t%v: %v\n", string(key), string(value))
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Fatal("Cannot fetch all data. ")
	}
}
