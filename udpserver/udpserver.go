package udpserver

import (
	"fmt"
	"log"
	"net"
)

func Server() error {
	pc, err := net.ListenPacket("udp", ":2222")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(pc, addr, buf[:n])
	}

}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Printf("%s\n", buf)
	pc.WriteTo(buf, addr)
}
