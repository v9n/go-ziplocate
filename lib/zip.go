package zip

import (
	"bytes"
	"encoding/gob"
)
type Point struct {
	X float64 
	Y float64
}

func (d *Point) GobEncode() ([]byte, error) {
	var data bytes.Buffer 
	enc := gob.NewEncoder(&data)
	enc.Encode(&d)
	return data.Bytes(), nil
}

func (d *Point) GobDecode(buf []byte) error {
    r := bytes.NewBuffer(buf)
    decoder := gob.NewDecoder(r)
    err := decoder.Decode(&d)
    if err!=nil {
        return err
    }
    return nil
}
