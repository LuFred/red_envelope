package context

import "net/http"

//Request Request扩展
type Request struct {
	*http.Request
}