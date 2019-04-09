package handler

import (
	"context"
	"fmt"
	"go-filestore-server/common"
	"go-filestore-server/config"
	"go-filestore-server/model"
	"go-filestore-server/service/account/proto"
	"go-filestore-server/util"
	"time"
)

// User: 用于实现UserServiceHandler接口的对象
type User struct{}

func GenToken(username string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + config.DefaultConfig.PwdSalt))
	return tokenPrefix + ts[:8]
}

func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, resp *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	if len(username) < 3 || len(passwd) < 5 {
		resp.Code = common.StatusParamInvalid
		resp.Message = "注册参数无效"
		return nil
	}

	encPasswd := util.Sha1([]byte(passwd + config.DefaultConfig.PwdSalt))
	suc := model.UserSignup(username, encPasswd)
	if suc {
		resp.Code = common.StatusOK
		resp.Message = "注册成功"
	} else {
		resp.Code = common.StatusRegisterFailed
		resp.Message = "注册失败"
	}
	return nil
}

func (u *User) Signin(ctx context.Context, req *proto.ReqSignin, resp *proto.RespSignin) error {
	username := req.Username
	password := req.Password

	encPasswd := util.Sha1([]byte(password + config.DefaultConfig.PwdSalt))

	pwdChecked := model.UserSignin(username, encPasswd)
	if !pwdChecked {
		resp.Code = common.StatusLoginFailed
		return nil
	}

	token := GenToken(username)
	upRes := model.UpdateToken(username, token)
	if !upRes {
		resp.Code = common.StatusServerError
		return nil
	}

	resp.Code = common.StatusOK
	resp.Token = token
	return nil
}

func (u *User) UserInfo(ctx context.Context, req *proto.ReqUserInfo, resp *proto.RespUserInfo) error {
	user, err := model.GetUserInfo(req.Username)
	if err != nil {
		resp.Code = common.StatusServerError
		resp.Message = "服务错误"
		return nil
	}

	if user.UserName == "" {
		resp.Code = common.StatusUserNotExists
		resp.Message = "用户不存在"
		return nil
	}

	resp.Code = common.StatusOK
	resp.Username = user.UserName
	resp.SignupAt = user.SignupAt
	resp.LastActiveAt = user.LastActiveAt
	resp.Status = int32(user.Status)
	resp.Email = user.Email
	resp.Phone = user.Phone
	return nil
}
