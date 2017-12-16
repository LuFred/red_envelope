package handler

import (
	"strconv"

	"github.com/astaxie/beego/context"
)

type Handler struct {
	// context data
	Ctx  *context.Context
	Data map[interface{}]interface{}
}

// HandlerInterface is an interface to uniform all controller handler.
type HandlerInterface interface {
	Init(ctx *context.Context)
}

// Init generates default values of controller operations.
func (h *Handler) Init(ctx *context.Context) {
	h.Ctx = ctx
	h.Data = ctx.Input.Data()
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
	h.Ctx.Output.JSON(h.Data["json"], hasIndent, hasEncoding)
}

// Abort stops controller handler and show the error data if code is defined in ErrorMap or code string.
func (c *Handler) Abort(code string) {
	status, err := strconv.Atoi(code)
	if err != nil {
		status = 200
	}
	c.CustomAbort(status, code)
}

// CustomAbort stops controller handler and show the error data, it's similar Aborts, but support status code and body.
func (c *Handler) CustomAbort(status int, body string) {
	// first panic from ErrorMaps, it is user defined error functions.
	// if _, ok := ErrorMaps[body]; ok {
	// 	c.Ctx.Output.Status = status
	// 	panic(body)
	// }
	// last panic user string
	c.Ctx.ResponseWriter.WriteHeader(status)
	c.Ctx.ResponseWriter.Write([]byte(body))
	//panic(ErrAbort)
}