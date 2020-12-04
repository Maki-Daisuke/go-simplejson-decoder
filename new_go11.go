package simplejsondecoder

import (
	"encoding/json"
	"io"
)

// New returns a new decoder that reads from r.
func New(r io.Reader) *Decoder {
	dec := &Decoder{*json.NewDecoder(r)}
	dec.UseNumber()
	return dec
}
