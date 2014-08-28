package zip

type Point struct {
	X float64 
	Y float64
}

//func (d *Data) GobEncode() ([]byte, error) {
    //w := new(bytes.Buffer)
    //encoder := gob.NewEncoder(w)
    //err := encoder.Encode(d.id)
    //if err!=nil {
        //return nil, err
    //}
    //err = encoder.Encode(d.name)
    //if err!=nil {
        //return nil, err
    //}
    //return w.Bytes(), nil
//}

//func (d *Data) GobDecode(buf []byte) error {
    //r := bytes.NewBuffer(buf)
    //decoder := gob.NewDecoder(r)
    //err := decoder.Decode(&d.id)
    //if err!=nil {
        //return err
    //}
    //return decoder.Decode(&d.name)
//}
