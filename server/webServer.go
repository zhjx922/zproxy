//web服务器Server
package server

import "net/http"

//WebServer 提供web页面服务
type WebServer struct {

}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func NewWebServer() *WebServer {
	return &WebServer{}
}