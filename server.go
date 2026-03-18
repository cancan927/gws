package gws

import "net/http"

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
	//TODO implement me
	panic("implement me")
}
