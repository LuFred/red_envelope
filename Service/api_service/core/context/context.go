package context

import (
	"net/http"
)

//Context b
type Context struct {
	Request  *Request
	Response *Response
	User     *User
}

//User token user info
type User struct {
	UserID int32  `json:"user_id"`
	Name   string `json:"name"`
}

//GetUser 获取user对象
func (c *Context) GetUser() *User {
	if c.User == nil {
		return nil
	}
	return c.User
}

//GetUserID 获取userid
func (c *Context) GetUserID() int32 {
	if c.User == nil {
		return 0
	}
	return c.User.UserID
}

//NewContext 创建context
func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	c := new(Context)
	c.Request = &Request{
		Request: r,
	}
	c.Response = &Response{
		ResponseWriter: w,
	}
	return c
}

//InitUser 初始化user对象
func (c *Context) InitUser(u *User) {
	if u != nil {
		c.User = u
	}
}