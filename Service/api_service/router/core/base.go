package core

import (
	"fmt"

	"net/http"
	"reflect"

	"golang.org/x/net/context"

	corecontext "github.com/lufred/red_envelope/Service/api_service/core/context"
	corehandle "github.com/lufred/red_envelope/Service/api_service/core/handle"
	"github.com/lufred/red_envelope/Service/api_service/middleware"
)

//HandlerTransport imnplements Transport by forwarding context to a handler
type HandlerTransport struct {
	Ctx context.Context
}

// HandlerOption sets a parameter for the InjectBaseHandler
type HandlerOption func(h *HandlerTransport)

//InjectBaseHandler 请求函数注入
func InjectBaseHandler(dto interface{}, h corehandle.HandlerInterface, mappingMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handleValue := reflect.ValueOf(h)
		if handleValue.Type().Kind() != reflect.Ptr {
			panic("must pass a pointer, not a value, to h destination")
		}
		newHandle := reflect.New(handleValue.Type().Elem())

		//init dto
		if dto != nil {
			err := InjectionParam(dto, r)
			if err != nil {
				w.Header().Add("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusBadRequest)
				em := fmt.Sprintf(`{"message":"%s"}`, err.Error())
				w.Write([]byte(em))
				return
			}
		}
		controllerCtx := corecontext.NewContext(r, w)
		ctxTokenInfo := r.Context().Value("tokeninfo")
		if ctxTokenInfo != nil {
			InitUser(controllerCtx, ctxTokenInfo)
		}
		if val := newHandle.MethodByName("Init"); val.IsValid() {
			val.Call([]reflect.Value{reflect.ValueOf(controllerCtx)})
		}
		//call method
		t := reflect.Indirect(newHandle).Type()
		//methods := make(map[string]string)
		if val := newHandle.MethodByName(mappingMethod); val.IsValid() {
			if dto == nil {
				val.Call(nil)
			} else {

				val.Call([]reflect.Value{reflect.ValueOf(dto)})

			}
		} else {
			panic("'" + mappingMethod + "' method doesn't exist in the controller " + t.Name())
		}
	}
}

//InitUser 初始化token携带用户信息
func InitUser(httpContext *corecontext.Context, ctxTokenInfo interface{}) {
	coreUser := &corecontext.User{}
	_oauthUser := ctxTokenInfo.(middleware.CtxValues).Get("OauthUser")
	oauthUserEntity := _oauthUser.(middleware.OauthUserEntity)
	coreUser.Name = oauthUserEntity.UserName
	coreUser.UserID = oauthUserEntity.UserID
	httpContext.InitUser(coreUser)
}