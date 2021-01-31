package main

import (
	"fmt"
	packaging "go-gateway/basic-web/packaging"
	"net"
)

func main() {
	// listen port
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := packaging.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect faild, err: %v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from client, data: %v\n", str)
	}

}
