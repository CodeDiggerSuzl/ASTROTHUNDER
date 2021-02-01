package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// 1. 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲时间
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  // 100-continue 状态码超时时间
	}
	// 2. 创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, // 请求超时时间
		Transport: transport,
	}
	// 3. 请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// read content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
