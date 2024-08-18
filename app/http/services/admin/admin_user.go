package admin

import (
	"blog/app/http/dao"
	"blog/app/http/requests"
	"blog/app/utils"
	"blog/pkg/cache"
	"blog/pkg/config"
	"blog/pkg/enum"
	"blog/pkg/jwt"
	"blog/pkg/logger"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// LoginService 登录
func (s *BaseService) LoginService(cxt *gin.Context, req *requests.BackLoginReq) *response.JsonResponse {
	// 查询用户
	user, err := dao.IDao.GetUserByAccount(req.Account)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("LoginService---GetUserByAccount---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误，请稍后再试！")
	}
	if user == nil {
		return response.Fail(enum.HttpFail, "账户不存在！")
	}

	// 判断密码是否正确
	password := utils.Md5Password(req.Password, user.Salt)
	if password != user.Password {
		return response.Fail(enum.HttpFail, "账户或者密码错误！")
	}

	// 验证成功，生成token
	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("LoginService---CreateToken---err:%v", err))
		return response.Fail(enum.HttpFail, "")
	}

	// 更新用户登录信息
	timeNow := time.Now().Format(time.DateTime)
	ip := utils.GetIPAddress(cxt.Request)
	err = dao.IDao.UpdateUserByParams(user.ID, map[string]interface{}{
		"last_ip":   ip,
		"last_time": timeNow,
	})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("LoginService---UpdateUserByParams---err:%v", err))
		return response.Fail(enum.HttpFail, "登录失败，请稍后再试！")
	}

	// 单点登录，将登录凭证保存到redis中
	cache.CH.RedisConn.Set(fmt.Sprintf("%s:%d", enum.JwtKey, user.ID), token, time.Duration(config.GetConf().Jwt.TTL)*time.Second)

	// 返回数据
	return response.Success("登录成功！", requests.BackLoginResp{
		AccessToken: token,
		ExpiresTime: int64(config.GetConf().Jwt.TTL),
		AdminId:     user.ID,
		Account:     user.Account,
		Nickname:    user.Nickname,
		CoverImg:    user.CoverImg,
		LastIp:      ip,
		LastTime:    timeNow,
	})
}

// LogoutService 退出
func (s *BaseService) LogoutService(cxt *gin.Context) *response.JsonResponse {
	id, exists := cxt.Get("userId")
	if !exists {
		return response.Fail(enum.TokenFailure, "你还未登录呢！")
	}
	userId := id.(int64)
	// 查询出用户信息
	if userId == int64(0) {
		return response.Fail(enum.TokenFailure, "登录状态已失效，请重新登录！")
	}

	// 清除redis中的凭证
	cache.CH.RedisConn.Del(fmt.Sprintf("%s:%d", enum.JwtKey, userId))

	return response.Success("")
}

// UserInfoService 获取用户信息
func (s *BaseService) UserInfoService(cxt *gin.Context) *response.JsonResponse {
	user, err := dao.IDao.GetUserByCtx(cxt)
	if err != nil {
		return response.Fail(enum.TokenFailure, err.Error())
	}

	return response.Success("获取成功！", requests.UserInfoResp{
		AdminId:  user.ID,
		Account:  user.Account,
		Nickname: user.Nickname,
		CoverImg: user.CoverImg,
		LastIp:   user.LastIp,
		LastTime: user.LastTime.Format(time.DateTime),
	})
}

// SetUserInfoService 修改用户信息
func (s *BaseService) SetUserInfoService(cxt *gin.Context, req *requests.SetUserInfoReq) *response.JsonResponse {
	user, err := dao.IDao.GetUserByCtx(cxt)
	if err != nil {
		return response.Fail(enum.TokenFailure, err.Error())
	}

	updateData := make(map[string]interface{})

	if req.Nickname != "" {
		updateData["nickname"] = strings.TrimSpace(req.Nickname)
	}

	if req.CoverImg != "" {
		updateData["cover_img"] = strings.TrimSpace(req.CoverImg)
	}

	if len(updateData) == 0 {
		return response.Fail(enum.HttpFail, "保存失败！")
	}

	err = dao.IDao.UpdateUserByParams(user.ID, updateData)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("SetUserInfoService---UpdateUserByParams---err:%v", err))
		return response.Fail(enum.TokenFailure, "保存失败！")
	}

	return response.Success("保存成功！")
}

// SetUserPasswordService 修改用户密码
func (s *BaseService) SetUserPasswordService(cxt *gin.Context, req *requests.SetUserPasswordReq) *response.JsonResponse {
	user, err := dao.IDao.GetUserByCtx(cxt)
	if err != nil {
		return response.Fail(enum.TokenFailure, err.Error())
	}

	// 重新生成密码
	salt := utils.GetRandomString(6, 0)
	newPassword := utils.Md5Password(req.NewPassword, salt)

	updateData := map[string]interface{}{
		"password": newPassword,
		"salt":     salt,
	}

	err = dao.IDao.UpdateUserByParams(user.ID, updateData)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("SetUserPasswordService---UpdateUserByParams---err:%v", err))
		return response.Fail(enum.TokenFailure, "修改失败！")
	}

	// 修改成功，将登录凭证注销，然后重新登录
	cache.CH.RedisConn.Del(fmt.Sprintf("%s:%d", enum.JwtKey, user.ID))

	return response.Success("修改成功！")
}
