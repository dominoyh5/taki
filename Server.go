package main

import (
	"fmt"
	"log"
	"net"
)

func StartServer() {
	fmt.Println("StartServer")

	l, err := net.Listen("tcp", ":5252")
	if nil != err {
		log.Println(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Println(err)
			continue
		}
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			log.Println(err)
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
