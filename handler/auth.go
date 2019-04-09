package handler

import (
	"github.com/gin-gonic/gin"
	"go-filestore-server/common"
	"go-filestore-server/util"
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
			resp := util.NewRespMsg(int(common.StatusTokenInvalid), "token无效", nil)
			c.JSON(http.StatusOK, resp)
			// c.Redirect(http.StatusFound, "/static/view/signin.html")
			return
		}
		c.Next()
	}
}
