package main

import (
	"github.com/zhjx922/zproxy/server"
	"net/http"
)

func main() {
	//Web Server
	go http.ListenAndServe(":7777", server.NewWebServer())

	//Proxy Server
	go http.ListenAndServe(":8888", server.NewProxyServer())

	select {}
}
