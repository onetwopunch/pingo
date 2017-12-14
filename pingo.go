package pingo

import (
  "golang.org/x/net/icmp"
  "golang.org/x/net/ipv4"
  "net"
  "os"
  "fmt"
  "time"
)

type Packet struct {
  Host string
  Timeout time.Duration
}

func Ping(data *Packet) (bool, error) {
  conn, err := listen(data.Timeout)
  if err != nil {
    return false, err
  }
  defer conn.Close()

  msg, err := icmpData()
  if err != nil {
    return false, err
  }

  ip, err := ipAddress(data.Host)
  if err != nil {
    return false, err
  }

  if _, err = conn.WriteTo(msg, ip); err != nil {
    return false, err
  }

  responseBuffer := make([]byte, 1500)
  n, _, err := conn.ReadFrom(responseBuffer)
  if err != nil {
    return false, err
  }

  // NOTE: Change this to 58 make the proto IPv6-ICMP
  responseMessage, err := icmp.ParseMessage(1, responseBuffer[:n])
  return responseMessage.Type == ipv4.ICMPTypeEchoReply, err
}

func listen(timeout time.Duration) (*icmp.PacketConn, error) {
  //TODO: We can make these configurable in the Packet struct
  conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
  conn.SetDeadline(time.Now().Add(timeout))
  if err != nil {
    return nil, fmt.Errorf("Could not start listener: %s", err)
  }
  return conn, nil
}

func ipAddress(host string) (*net.UDPAddr, error) {
  addrs, err := net.LookupHost(host)
  if err != nil {
    return nil, err
  }

  if len(addrs) == 0 {
    err = fmt.Errorf("Unable to resolve host")
    return nil, err
  }

  return &net.UDPAddr{IP: net.ParseIP(addrs[0]), Zone: "en0"}, nil
}

func icmpData() ([]byte, error) {
  msg := icmp.Message{
    Type: ipv4.ICMPTypeEcho,
    Code: 0,
    Body: &icmp.Echo{
        ID: os.Getpid() & 0xffff,
        Seq: 1,
        Data: []byte("HELLO"),
    },
  }
  data, err := msg.Marshal(nil)
  return data, err
}
