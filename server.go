package gws

import (
	"net"
	"net/http"
)

type Server interface {
	http.Handler
	Start(addr string) error
}

type HTTPServer struct {
}

type HTTPSServer struct {
	HTTPServer
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 在这里，可以让用户注册 after start 回调
	// 比如在这里执行一些你业务所需的前置条件，比如连接数据库，或者加载一些配置文件等

	return http.Serve(l, nil)
}
