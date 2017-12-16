package handle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/lufred/red_envelope/Service/api_service/core/context"
)

//Handler base handler
type Handler struct {
	// context data
	Ctx  *context.Context
	Data map[interface{}]interface{}
}

//HandlerInterface handler interface
type HandlerInterface interface {
	Init(ctx *context.Context)
}

//Init 实现HandlerInterface接口
func (h *Handler) Init(ctx *context.Context) {
	h.Ctx = ctx
	h.Data = make(map[interface{}]interface{})
}

// ServeJSON sends a json response with encoding charset.
func (h *Handler) ServeJSON(encoding ...bool) {
	var (
		hasIndent   = true
		hasEncoding = false
	)
	// if BConfig.RunMode == PROD {
	// 	hasIndent = false
	// }
	if len(encoding) > 0 && encoding[0] {
		hasEncoding = true
	}
	h.Ctx.Response.JSON(h.Data["json"], hasIndent, hasEncoding)
}

// Abort stops controller handler and show the error data if code is defined in ErrorMap or code string.
func (h *Handler) Abort(code string) {
	status, err := strconv.Atoi(code)
	if err != nil {
		status = 200
	}
	h.CustomAbort(status, code)
}

//ErrorAbort 500
func (h *Handler) ErrorAbort(errMsg string) {
	h.Ctx.Response.WriteHeader(http.StatusInternalServerError)
	h.Ctx.Response.Write([]byte(errMsg))
}

//BadRequestAbort 400
func (h *Handler) BadRequestAbort(msg string) {
	//ignore error
	h.Ctx.Response.Header("Content-Type", "application/json")
	h.Ctx.Response.WriteHeader(http.StatusBadRequest)
	em := fmt.Sprintf(`{"message":"%s"}`, msg)
	h.Ctx.Response.Write([]byte(em))
}

//NoContentAbort 204
func (h *Handler) NoContentAbort() {
	h.Ctx.Response.WriteHeader(http.StatusNoContent)
}

//UnauthorizedAbort 401
func (h *Handler) UnauthorizedAbort(msg string) {
	h.Ctx.Response.WriteHeader(http.StatusUnauthorized)
	h.Ctx.Response.Write([]byte(msg))
}

//NotFoundAbort 404
func (h *Handler) NotFoundAbort(msg string) {
	h.Ctx.Response.WriteHeader(http.StatusNotFound)
	h.Ctx.Response.Write([]byte(msg))
}

// CustomAbort stops controller handler and show the error data, it's similar Aborts, but support status code and body.
func (h *Handler) CustomAbort(status int, body string) {
	h.Ctx.Response.WriteHeader(status)
	h.Ctx.Response.Write([]byte(body))

}

//CustomJsonAbort stops controller handler and show the error data,it's similar Aborts,but support status code and json format body.
func (h *Handler) CustomJsonAbort(status int, data interface{}) {
	h.Ctx.Response.Header("Content-Type", "application/json")
	h.Ctx.Response.WriteHeader(400)
	h.Ctx.Response.JSON(data, false, false)

}
