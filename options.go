package sessions

import "net/http"

type AweOptions struct {
	Path   string
	Domain string
	// MaxAge=0 表示未指定 Max-Age 属性，并且 cookie 将在浏览器会话结束后被删除。
	// MaxAge<0 表示立即删除 cookie。
	// MaxAge>0 表示存在 Max-Age 属性，并以秒为单位给出。
	MaxAge   int
	Secure   bool
	HttpOnly bool
	//Partitioned bool
	SameSite http.SameSite
}
