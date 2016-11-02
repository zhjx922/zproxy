//代理服务器Server
package server

import (
	"net/http"
	"fmt"
	"net"
	"time"
	"strings"
)

//ProxyServer 提供代理服务
type ProxyServer struct {

}

func (s *ProxyServer) copyConn(iConn net.Conn, oConn net.Conn) {
	buffer := [4096]byte{}
	defer iConn.Close()
	defer oConn.Close()
	for {
		if n, err := iConn.Read(buffer[:]); err == nil {
			oConn.Write(buffer[:n])
		} else {
			return
		}
	}
}

func (s *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if iConn, _, err := w.(http.Hijacker).Hijack(); err == nil {
		fmt.Printf("url:%s\n", r.URL.Host)

		host := r.URL.Host
		if strings.IndexRune(host, ':') == -1 {
			host += ":80"
		}

		if oConn, err := net.DialTimeout("tcp", host, time.Second * 5); err == nil {
			if r.Method == "CONNECT" {
				//客户端你好,我已成功连接对方https服务器,请继续发送数据
				iConn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
			} else {
				r.Write(oConn)
			}

			go s.copyConn(iConn, oConn)
			go s.copyConn(oConn, iConn)
		} else {
			iConn.Close()
		}
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func NewProxyServer() *ProxyServer {
	return &ProxyServer{}
}