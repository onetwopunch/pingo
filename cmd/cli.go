package main

import (
  "flag"
  "time"
  "fmt"
  "os"
  "github.com/onetwopunch/pingo"
)

func main() {
  timeout := flag.Duration("t", 1 * time.Second, "Request timeout. Defaults to 1s.")
  host := flag.String("H", "", "Required: Host or IPv4 address to ping")
  flag.Parse()

  packet := &pingo.Packet{ Host: *host, Timeout: *timeout }
  if ok, err := pingo.Ping(packet); ok {
    fmt.Println("PONG")
  } else {
    fmt.Println(err)
    os.Exit(1)
  }
}
