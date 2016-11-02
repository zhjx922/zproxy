package main

import (
	"github.com/zhjx922/zproxy/server"
	"net/http"
)

func main() {

	webServer := server.NewWebServer()
	go http.ListenAndServe(":7777", webServer)

	proxyServer := server.NewProxyServer()
	go http.ListenAndServe(":8888", proxyServer)

	/*
	isStop := make(chan bool, 1)
	<- isStop
	*/
	select {}
}
