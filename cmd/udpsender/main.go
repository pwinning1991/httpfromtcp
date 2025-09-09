package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+c to exit\n", addr)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Message sent: %s\n", text)

	}
}
