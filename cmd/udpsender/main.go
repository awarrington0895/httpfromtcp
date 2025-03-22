package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:42069")

	if err != nil {
		log.Panicf("unable to resolve udp address: %v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		log.Panicf("unable to dial udp address: %v\n", err)
	}

	defer conn.Close()

	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := stdin.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		_, err = conn.Write([]byte(line))

		if err != nil {
			log.Printf("%v", err)
		}
	}

}
