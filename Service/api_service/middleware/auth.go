package middleware

import (
	"encoding/json"

	"net/http"
	"strings"

	"strconv"

	"github.com/coocood/freecache"
	"github.com/lufred/red_envelope/Service/api_service/config"
	"golang.org/x/net/context"
)

var cache *freecache.Cache

const (
	expireTime int = 30 * 60
	cacheSize  int = 100 * 1024 * 1024
)

func init() {
	cache = freecache.NewCache(cacheSize)
}

type oauthAuth struct {
	opts AuthOptions
	h    http.Handler
}
type AuthOptions struct {
	Token               string
	AuthFunc            func(string, *http.Request) (*http.Request, bool)
	UnauthorizedHandler http.Handler
}
type CtxValues struct {
	m map[string]interface{}
}

func (v CtxValues) Get(key string) interface{} {
	return v.m[key]
}

//OauthUserEntity oauth用户对象
type OauthUserEntity struct {
	UserID   int32  `json:"userId"`
	UserName string `json:"userName"`
}

//ServeHTTP
func (o oauthAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if we have a user-provided error handler, else set a default
	if o.opts.UnauthorizedHandler == nil {
		o.opts.UnauthorizedHandler = http.HandlerFunc(defaultUnauthorizedHandler)
	}
	if r.Method != http.MethodOptions {
		// Check that the provided details match
		if r, resu := o.authenticate(r); resu == false {
			o.requestAuth(w, r)
			return
		}
	}
	// Call the next handler on success.
	o.h.ServeHTTP(w, r)
}

//OauthMiddleware oauth认证中间件
func OauthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := oauthAuth{}
		if auth.opts.UnauthorizedHandler == nil {
			auth.opts.UnauthorizedHandler = http.HandlerFunc(defaultUnauthorizedHandler)
		}
		if r, resu := auth.authenticate(r); resu == false {
			auth.requestAuth(w, r)
		} else {
			next.ServeHTTP(w, r)
		}

	}
	return http.HandlerFunc(fn)

}
func (o *oauthAuth) authenticate(r *http.Request) (*http.Request, bool) {
	const oauthScheme string = "bearer"
	auth := r.Header.Get("Authorization")
	//default token 默认通过
	if auth == config.ProConfig.DefaultToken {
		return r, true
	}
	if auth == "" {
		auth = "bearer " + r.URL.Query().Get("access_token")
	}
	authstrings := strings.Split(auth, " ")
	// Confirm the request is sending Basic Authentication credentials.
	if len(authstrings) != 2 || strings.ToLower(authstrings[0]) != oauthScheme {
		return r, false
	}
	givenToken := auth[len(oauthScheme)+1:]
	if o.opts.AuthFunc == nil {
		o.opts.AuthFunc = o.simpleBasicFunc
	}
	return o.opts.AuthFunc(givenToken, r)
}

func (o *oauthAuth) simpleBasicFunc(token string, r *http.Request) (req *http.Request, result bool) {
	defer func() {
		if err := recover(); err != nil {
			result = false
		}
	}()
	var oauthUser = &OauthUserEntity{}
	cacheUserMeta := getCacheOauth(token)
	if cacheUserMeta != nil {
		json.Unmarshal(cacheUserMeta, oauthUser)
	} else {
		//
		//todo 通过token调用指定服务查询用户信息
		//这里作为测试直接取token中_末尾作为用户
		ts := strings.Split(token, "_")
		if len(ts) == 2 {
			oauthUser = &OauthUserEntity{}
			_id, _ := strconv.Atoi(ts[1])
			oauthUser.UserID = int32(_id)
			userMetaString, _ := json.Marshal(*oauthUser)
			cache.Set([]byte(token), userMetaString, expireTime)
		}
	}
	if oauthUser.UserID > 0 {
		v := CtxValues{
			map[string]interface{}{
				"OauthUser": *oauthUser,
			},
		}
		req = r.WithContext(context.WithValue(r.Context(), "tokeninfo", v))
		result = true
	}
	return
}
func getCacheOauth(token string) []byte {
	got, err := cache.Get([]byte(token))
	if err != nil {
		return nil
	}
	return got
}

// defaultUnauthorizedHandler provides a default HTTP 401 Unauthorized response.
func defaultUnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

// Require authentication, and serve our error handler otherwise.
func (o *oauthAuth) requestAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", `Bearer realm="example"`)
	o.opts.UnauthorizedHandler.ServeHTTP(w, r)
}
