package route

import (
	"github.com/gin-gonic/gin"
	"go-filestore-server/service/apigw/handler"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static/", "./static")

	router.GET("/user/signup", handler.SignupHandler)
	router.POST("/user/signup", handler.DoSignupHandler)

	router.GET("/user/signin", handler.SigninHandler)
	router.POST("/user/signin", handler.DoSigninHandler)

	router.POST("/user/info", handler.UserInfoHandler)

	return router
}
