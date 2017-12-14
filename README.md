# Pingo
Simple Go library and executable wrapping ICMP

## Install

```
go get github.com/onetwopunch/pingo
```

## Use it like so:

```
import "pingo"
packet := &pingo.Packet{ Host: *host, Timeout: *timeout }
if ok, err := pingo.Ping(packet); ok {
  fmt.Println("PONG")
} else {
  fmt.Println(err)
  os.Exit(1)
}
```
