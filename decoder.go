package simplejsondecoder

import (
	"encoding/json"
	"unsafe"

	"github.com/bitly/go-simplejson"
)

// Decoder is a simple wrapper of encoding/json's Decoder returning simplejson.Json.
type Decoder struct {
	json.Decoder
}

func (dec *Decoder) Decode() (*simplejson.Json, error) {
	j := new(simplejson.Json)
	// We cannot populate unmarshaled data into simplejson.Json's private field.
	// So, use a black magic!
	x := (*struct{ data interface{} })(unsafe.Pointer(j))
	err := dec.Decoder.Decode(&x.data)
	return j, err
}
