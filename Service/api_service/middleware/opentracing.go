package middleware

import (
	"log"
	"net/http"
	"reflect"

	extopentracing "github.com/lufred/red_envelope/Service/api_service/core/extension/opentracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type OpentracingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newOpentracingResponseWriter(w http.ResponseWriter) *OpentracingResponseWriter {
	return &OpentracingResponseWriter{w, 0}
}
func (orw *OpentracingResponseWriter) WriteHeader(code int) {
	orw.statusCode = code
	orw.ResponseWriter.WriteHeader(code)
}

//OpentracingMiddleware opentracing middleware
func OpentracingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(reflect.TypeOf(w))

		span := extopentracing.Tracer.StartSpan(r.RequestURI)
		span.SetTag(string(ext.HTTPUrl), r.RequestURI)
		span.SetTag(string(ext.HTTPMethod), r.Method)
		defer span.Finish()
		req := r.WithContext(opentracing.ContextWithSpan(r.Context(), span))

		orw := newOpentracingResponseWriter(w)

		next.ServeHTTP(w, req)
		span.SetTag(string(ext.HTTPStatusCode), orw.statusCode)

	}
	return http.HandlerFunc(fn)
}