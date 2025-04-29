package api

import (
	config "app/api/config"
	"app/api/routes"
	"app/api/routes/profile"
	api "github.com/goaperture/goaperture/lib/aperture"
	"net/http"
)

func (server *Server) Run(port int, token *string) error {
	api.NewRoute(&server.aperture, api.Route[routes.IndexInput, config.Payload]{
		Handler: routes.Index,
		Path:    "/index",
		Test:    routes.IndexTest,
	})
	api.NewRoute(&server.aperture, api.Route[routes.PrivateInput, config.Payload]{
		Handler: routes.Private,
		Path:    "/private",
		Test:    routes.PrivateTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.RefreshInput, config.Payload]{
		Handler: profile.Refresh,
		Path:    "/profile/refresh",
		Test:    profile.RefreshTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SinginInput, config.Payload]{
		Handler: profile.Singin,
		Path:    "/profile/singin",
		Test:    profile.SinginTest,
	})
	api.NewRoute(&server.aperture, api.Route[profile.SingoutInput, config.Payload]{
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
