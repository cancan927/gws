package gws

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var s Server
	// 用户依然可以直接使用 http.ListenAndServe 和 http.ListenAndServeTLS 来启动服务器
	http.ListenAndServe("8080", s)
	http.ListenAndServeTLS("8081", "cert.pem", "key.pem", s)

	// 也可以使用自定义的 Start 方法来启动
	s.Start("8080")
}
