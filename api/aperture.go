package api

import (
	config "app/api/config"
	"app/api/routes"
	"app/api/routes/profile"
	"net/http"

	api "github.com/goaperture/goaperture/lib/aperture"
)

func (server *Server) Run(port int, token *string) error {
	api.NewRoute(&server.aperture, api.Route[routes.IndexInput, config.Client]{
		Handler: routes.Index,
		Path:    "/index",
		Test:    routes.IndexTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.RefreshInput, config.Client]{
		Handler: profile.Refresh,
		Path:    "/profile/refresh",
		Test:    profile.RefreshTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SinginInput, config.Client]{
		Handler: profile.Singin,
		Path:    "/profile/singin",
		Test:    profile.SinginTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SingoutInput, config.Client]{
		Handler: profile.Singout,
		Path:    "/profile/singout",
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
