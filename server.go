package gws

import "net/http"

type Server interface {
	http.Handler
}
