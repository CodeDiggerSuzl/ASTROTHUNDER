package main

import (
	"fmt"
	packing "go-gateway/basic-web/packaging"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("conn failed ,err: %v\n", err)
		return
	}
	_ = packing.Encode(conn, "hello world 0!!!")
}
