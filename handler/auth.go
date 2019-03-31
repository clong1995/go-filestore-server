package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin版本
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		// 验证登陆token是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			c.Abort()
			c.Redirect(http.StatusFound, "/static/view/signin.html")
			return
		}
		c.Next()
	}
}

// http请求拦截器
/*func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		username := r.Form.Get("username")
		token := r.Form.Get("token")

		// 验证登陆token是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			// 失败跳转到登陆界面
			http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
			return
		}
		h(w, r)
	})
}
*/
