# go-simplejson-decoder

A brain-dead simple json.Decoder wrapper for [go-simplejson](https://github.com/bitly/go-simplejson).


# Usage

```
decoder := New(reader)
for {
    j, err = decoder.Decode()
    if err != nil {
        break
    }
    // Here, j is simplejson object.
    // You can access properties of JSON data simply!
    i := j.MustMap()["number"].MustInt()
    .
    .
    .
}
```


# LICENCE

MIT License


# Author

Daisuke (yet another) Maki
