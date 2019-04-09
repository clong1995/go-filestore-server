package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"go-filestore-server/common"
	"go-filestore-server/service/account/proto"
	"go-filestore-server/util"
	"log"
	"net/http"
)

var (
	userCli proto.UserService
)

func init() {
	service := micro.NewService()

	service.Init()

	userCli = proto.NewUserService("go.micro.service.user", service.Client())
}

func SignupHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

func DoSignupHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userCli.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}

func SigninHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

func DoSigninHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	resp, err := userCli.Signin(context.TODO(), &proto.ReqSignin{
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	cliResp := util.RespMsg{
		Code: int(common.StatusOK),
		Msg:  "登陆成功",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    resp.Token,
		},
	}

	c.Data(http.StatusOK, "application/json", cliResp.JSONBytes())
}

func UserInfoHandler(c *gin.Context) {
	username := c.Request.FormValue("username")
	resp, err := userCli.UserInfo(context.TODO(), &proto.ReqUserInfo{
		Username: username,
	})
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	cliResp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: gin.H{
			"Username":   username,
			"SignupAt":   resp.SignupAt,
			"LastActive": resp.LastActiveAt,
		},
	}
	c.Data(http.StatusOK, "application/json", cliResp.JSONBytes())
}
