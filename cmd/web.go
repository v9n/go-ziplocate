package cmd

import (
	"fmt"
	"log"
	"bytes"
	"encoding/binary"

	"github.com/codegangsta/cli"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/kureikain/go-ziplocate/lib"
)


var CmdWeb = cli.Command{
	Name: "web",
	Usage: "web -p port",
	Description: `Run API server on the specify port. default to 12385`,
	Action: runWeb,
	Flags: []cli.Flag{},
}

func runWeb(*cli.Context) {
	var geocoder zip.Point
	log.Printf("Web runs on Port %d", 12385)

	db, err := leveldb.OpenFile("zipdata", nil)
	if err != nil {
		log.Fatal("Db not found")
	}
	defer db.Close()
	zip := "95111"
	geo, err := db.Get([]byte(zip), nil)
	if err != nil {
		log.Fatal("Canot found the zip code")
	}
	fmt.Printf("Geo Byte Array: %v", geo)

	buffer := bytes.NewReader(geo)
	err = binary.Read(buffer, binary.LittleEndian, &geocoder)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("Geo Coder: %v", geocoder)

	log.Println("Finish webing!")
}

func fetchAllZip(db leveldb.DB) {
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


