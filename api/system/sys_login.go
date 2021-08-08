package system

import (
	"blogo/models/request"
	"blogo/models/response"
	"blogo/services/account_service"
	"blogo/utils"
)

type LoginApi struct {
}

func (b *LoginApi) Login(c *response.GinContextE) {
	var loginInfo request.LoginRequest
	c.C.ShouldBindJSON(&loginInfo)
	if loginInfo.Username == "" || loginInfo.Password == "" {
		c.FailWithMessage(response.API_ERROR, "账号或密码不能为空")
		return
	}
	_, err := account_service.CheckAuth(loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.FailWithMessage(response.API_ERROR, "账号或密码错误")
		return
	}
	accessToken, expireTime, err := utils.GetToken(0, loginInfo.Username) // FIXME: 无法取得UserID
	if err != nil {
		c.FailWithMessage(response.API_ERROR, "获取Token失败")
		return
	}
	c.OkWithDetailed(response.SUCCESS, &response.LoginResponse{
		AccessToken: accessToken,
		ExpireTime:  expireTime,
		User:        loginInfo.Username},
		"登录成功")
}
