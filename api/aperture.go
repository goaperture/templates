package api

import (
	"app/api/config"
	"app/api/routes"
	"net/http"

	api "github.com/goaperture/goaperture/lib/aperture"
)

func (server *Server) Run(port int, token *string) error {
	api.NewRoute(&server.aperture, api.Route[routes.IndexInput, config.Client]{
		Path:    "/index",
		Handler: routes.Index,
		Test:    routes.IndexTest,
	})
	return server.aperture.Run(port, token)
}

type Server struct {
	aperture api.Aperture
}

func NewServer() *Server {
	return &Server{aperture: *api.NewServer()}
}
func (server *Server) Middleware(middleware func(next http.Handler) http.Handler) {
	server.aperture.Middleware = middleware
}
func Serve(port int, token *string) {
	if err := NewServer().Run(port, token); err != nil {
		panic(err)
	}
}
