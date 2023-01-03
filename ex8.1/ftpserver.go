package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept error: %v", err)
			continue
		}

		session := ftpSession{
			conn: conn,
		}
		go session.handle()
	}
}

type ftpSession struct {
	conn net.Conn
}

func (f ftpSession) handle() {
	defer f.conn.Close()

	f.conn.Write([]byte("220 ahoy, ftp server ready\n"))

	input := bufio.NewScanner(f.conn)
	for input.Scan() {
		f.handleCommand(input.Text())
	}
}

func (f ftpSession) handleCommand(cmd string) {
	f.conn.Write([]byte(fmt.Sprintf("got command: %s\n", cmd)))
}
