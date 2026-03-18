package gws

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var s Server
	http.ListenAndServe("8080", s)
	http.ListenAndServeTLS("8081", "cert.pem", "key.pem", s)
}
