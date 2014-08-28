package cmd

import (
	"errors"
	"strings"
	"strconv"
	"fmt"
	"log"
	//"bytes"
	//"encoding/binary"

	"github.com/codegangsta/cli"
	"github.com/syndtr/goleveldb/leveldb"
	zipi "github.com/kureikain/go-ziplocate/lib"
	"net/http"	
	"encoding/json"
	"io"
)

var CmdWeb = cli.Command{
	Name: "web",
	Usage: "web -p port",
	Description: `Run API server on the specify port. default to 12385`,
	Action: runWeb,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name: "p",
			Value: 0,
			Usage: "port to bind",
		},
},
}

func runWeb(c *cli.Context) {
	port := c.Int("p")
	log.Printf("Web runs on Port %d", port)

	db, err := leveldb.OpenFile("zipdata", nil)
	if err != nil {
		log.Fatal("Db not found")
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/", func (w http.ResponseWriter, r *http.Request) {
		log.Printf(r.URL.Path)
		urlSegment := strings.Split(r.URL.Path, "/")
		if len(urlSegment)<2 {
			io.WriteString(w, "{\"error\":\"Zip not found\"}")
			//w.Header().Set("Content-Type", "application/json")
			panic("Fail to parse URL")
		}
		log.Printf("Result %v", urlSegment)
		zip := urlSegment[2]
		log.Printf("zip= %v.", zip)
		result, err := fetchZip(db, zip)
		if err != nil {
			io.WriteString(w, "{\"error\":\"Zip not found\"}")
		} else {

			b, err := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			if err != nil {
				io.WriteString(w, "{\"error\":\"Zip not found\"}")
			} else {
				io.WriteString(w, string(b))
			}
		}
	})
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("curl -i domain/api/zip"))
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
  //if err != nil {
		//panic(err)
  //}
	log.Println("Finish webing!")
	//fetchAllZip(db)
}

func fetchZip(db *leveldb.DB, zip string) (*zipi.Point, error) {
	//geocoder := new(zip.Point)
	geo, err := db.Get([]byte(zip), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("zip: %v not found", zip))
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
	return geocoder, nil
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
