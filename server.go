package gws

import (
	"net"
	"net/http"
)

type Server interface {
	http.Handler
	Start(addr string) error

	// AddRoute 路由注册
	// method 是 HTTP 方法
	// path 是路由
	// handFunc 是你的业务逻辑
	AddRoute(method string, path string, handler HandleFunc)
}

type HandleFunc func(ctx *Context)

// 确保 HTTPServer 实现了 Server 接口
var _ Server = &HTTPServer{}

type HTTPServer struct {
}

func (h *HTTPServer) AddRoute(method string, path string, handler HandleFunc) {
	//TODO implement me
	panic("implement me")
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

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
