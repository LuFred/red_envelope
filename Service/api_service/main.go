// @APIVersion 1.0.0
// @APITitle Swagger Example API
// @APIDescription Swagger Example API
// @BasePath http://localhost:8077/api-json
// @Contact fred@gmail.com
// @TermsOfServiceUrl http://yvasiyarov.com/
// @License BSD
// @LicenseUrl http://yvasiyarov.com/
package main

import (
	"flag"
	"net/http"

	"github.com/lufred/red_envelope/Service/api_service/config"
	"github.com/lufred/red_envelope/Service/api_service/microservice_client"
	"github.com/lufred/red_envelope/Service/api_service/router"
	"github.com/lufred/red_envelope/util/log"
)

func init() {
	log.Enabled = true
	log.Debugged = config.ProConfig.Debug
	flag.Parse()
	config.RegisterConfig()
	microservice_client.RegisterMicroService()

}
func main() {
	log.Infof("listen:%s", config.ProConfig.Listen)
	if err := http.ListenAndServe(config.ProConfig.Listen, router.Router); err != nil {
		log.Errorf("failed to server:%v", err)
	}
}
