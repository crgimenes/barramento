package udp

import (
	"barramento/command"
	"barramento/config"
	"log"
	"net"
)

type UDP struct {
	network     string
	multicastIP string
	cfg         *config.Config
	cmd         command.CMD
}

func New(network string, cfg *config.Config, cmd command.CMD) *UDP {
	m := cfg.IPv4
	if network == "udp6" {
		m = cfg.IPv6
	}
	return &UDP{
		network:     network,
		multicastIP: m,
		cfg:         cfg,
		cmd:         cmd,
	}
}

func (u *UDP) multicast() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr(u.network, u.multicastIP)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialUDP(u.network, nil, addr)
	return conn, err
}

func (u *UDP) Send(payload []byte) error {
	//conn, err := multicast("192.168.0.255:2222")
	//conn, err := multicast("224.0.0.1:2222")
	conn, err := u.multicast()
	if err != nil {
		log.Println("error resolve multicast", err)
		return err
	}
	defer conn.Close()
	_, err = conn.Write(payload)
	return err
}

func (u *UDP) Server() error {
	pc, err := net.ListenPacket(u.network, u.cfg.ServerAddress)
	if err != nil {
		log.Println("error listen packet", err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			log.Println("error reading packet", err)
			continue
		}
		go u.runCMD(pc, addr, buf[:n])
	}
}

func (u *UDP) runCMD(pc net.PacketConn, addr net.Addr, buf []byte) {
	err := u.cmd.Run(buf)
	if err != nil {
		log.Println(err)
	}
}
