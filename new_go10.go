// +build !go1.1

package simplejsondecoder

import (
	"encoding/json"
	"io"
)

func New(r io.Reader) *Decoder {
	dec := &Decoder{*json.NewDecoder(r)}
	return dec
}
