package router

import (
	"github.com/go-chi/chi"
	"github.com/lufred/red_envelope/Service/api_service/middleware"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()
	Router.Use(middleware.CORSMiddleware)
	Router.Use(middleware.OauthMiddleware)
	routerInit()

}
func routerInit() {
	redEnvelopeRouter()
}
