package dao

import (
	"blog/app/models"
	"blog/pkg/database"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAdminUsers interface {
	// 根据账号查询用户
	GetUserByAccount(account string) (*models.AdminUsers, error)
	// 更新用户信息
	UpdateUserByParams(id int64, data map[string]interface{}) error
	// 根据用户id查询
	GetUserById(id int64) (*models.AdminUsers, error)
	// 根据上下文获取用户
	GetUserByCtx(cxt *gin.Context) (*models.AdminUsers, error)
}

func (d *Dao) GetUserByAccount(account string) (*models.AdminUsers, error) {
	var adminUsers models.AdminUsers
	err := database.DB.MysqlConn.Model(&models.AdminUsers{}).Where("binary account = ?", account).First(&adminUsers).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &adminUsers, nil
}

func (d *Dao) UpdateUserByParams(id int64, data map[string]interface{}) error {
	return database.DB.MysqlConn.Model(&models.AdminUsers{}).Where("id = ?", id).Updates(data).Error
}

func (d *Dao) GetUserById(id int64) (*models.AdminUsers, error) {
	var adminUsers models.AdminUsers
	err := database.DB.MysqlConn.Model(&models.AdminUsers{}).Where("id = ?", id).First(&adminUsers).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &adminUsers, nil
}

func (d *Dao) GetUserByCtx(cxt *gin.Context) (*models.AdminUsers, error) {
	id, exists := cxt.Get("userId")
	if !exists {
		return nil, errors.New("你还未登录呢！")
	}
	userId := id.(int64)
	// 查询出用户信息
	if userId == int64(0) {
		return nil, errors.New("登录状态已失效，请重新登录！")
	}

	user, err := d.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("登录状态已失效，请重新登录！")
	}
	return user, nil
}
