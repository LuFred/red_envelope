package router

import (
	"github.com/go-chi/chi"
	"github.com/lufred/red_envelope/Service/api_service/handler/redenvelopehandler"
	redemodel "github.com/lufred/red_envelope/Service/api_service/model/redenvelope"
	routercore "github.com/lufred/red_envelope/Service/api_service/router/core"
)

func redEnvelopeRouter() {
	h := &redenvelopehandler.RedEnvelopeHandler{}
	Router.Route("/rede", func(r chi.Router) {
		r.Post("/", routercore.InjectBaseHandler(&redemodel.CreateRequest{}, h, "Create"))
		r.Post("/take", routercore.InjectBaseHandler(&redemodel.TakeRequest{}, h, "Take"))
		r.Get("/balance", routercore.InjectBaseHandler(&redemodel.GetBalanceRequest{}, h, "GetBalance"))
		r.Get("/record", routercore.InjectBaseHandler(&redemodel.GetRecordRequest{}, h, "GetRecord"))

	})

}
