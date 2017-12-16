package handler

import (
	"fmt"
	"strconv"

	middlewre "github.com/lufred/red_envelope/Service/api_service/middleware"
	"golang.org/x/net/context"
)

//GetUserID 从context中获取userid
func GetUserID(ctx context.Context) (int64, error) {
	ctxTokenInfo := ctx.Value("tokeninfo")
	if ctxTokenInfo == nil {
		return 0, fmt.Errorf("Unauthorized")
	}
	_userID := ctxTokenInfo.(middlewre.CtxValues).Get("UserID")
	if _userID == nil {
		return 0, fmt.Errorf("Unauthorized")
	}
	return _userID.(int64), nil

}

// HandleError records an error and the operation.
type HandleError struct {
	ServerCilent string
	Op           string
	Err          error
}

func (e *HandleError) Error() string { return e.ServerCilent + "  " + e.Op + " : " + e.Err.Error() }

//BadRequestError http status bad request error
type BadRequestError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BadRequestError) Error() string { return strconv.Itoa(e.Code) + " : " + e.Message }
