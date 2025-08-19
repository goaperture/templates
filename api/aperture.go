package api

import (
	config "app/api/config"
	"app/api/routes/v1"
	"app/api/routes/v1/profile"
	api "github.com/goaperture/goaperture/lib/aperture"
	"net/http"
)

func (server *Server) Run(port int, token *string) error {
	api.NewRoute(&server.aperture, api.Route[v1.IndexInput, config.Payload]{
		Handler: v1.Index,
		Path:    "/v1/index",
		Test:    v1.IndexTest,
	})
	api.NewRoute(&server.aperture, api.Route[v1.PrivateInput, config.Payload]{
		Handler: v1.Private,
		Path:    "/v1/private",
		Test:    v1.PrivateTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.RefreshInput, config.Payload]{
		Handler: profile.Refresh,
		Path:    "/v1/profile/refresh",
		Test:    profile.RefreshTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SinginInput, config.Payload]{
		Handler: profile.Singin,
		Path:    "/v1/profile/singin",
		Test:    profile.SinginTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SingoutInput, config.Payload]{
		Handler: profile.Singout,
		Path:    "/v1/profile/singout",
		Test:    profile.SingoutTest,
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
