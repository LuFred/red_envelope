package middleware

import (
	"net/http"

	"github.com/lufred/red_envelope/util/log"
	// 引入proto包
)

//CORSMiddleware 支持跨域
func UrlProcessMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request url: %v", r.URL.Path)
		// if r.URL.Path[len(r.URL.Path)-1] == '/' {

		// 	r.URL.Path = r.URL.Path[:len(r.URL.Path)-1]
		// 	log.Println(r.URL.Path)
		// 	log.Println(r.Method)
		// 	log.Println("-------")
		// }
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}
