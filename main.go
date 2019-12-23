package main

import (
	"encoding/hex"
	"flag"
	"log"
	"net"
	"strings"
)

func main() {
	args := struct {
		Mac     string
		Address string
		Port    int
	}{}
	flag.StringVar(&args.Mac, "m", "", "mac address")
	flag.StringVar(&args.Address, "a", "", "ip")
	flag.IntVar(&args.Port, "p", 9, "port")
	flag.Parse()
	if args.Mac == "" || args.Address == "" {
		flag.Usage()
		return
	}
	payload := make([]byte, 6, 1024)
	for i := 0; i < 6; i++ {
		payload[i] = 255
	}
	m := strings.ReplaceAll(args.Mac, ":", "")
	mb, err := hex.DecodeString(m)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 16; i++ {
		payload = append(payload, mb...)
	}
	l, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(args.Address),
		Port: args.Port,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	n, err := l.Write(payload)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Magic package has sent. size: ", n)
}
