package dao

import (
	"blog/app/models"
	"blog/pkg/database"
	"gorm.io/gorm"
)

type ITestDao interface {
	GetTestById(Id int64) (*models.Test, error)
}

func (d *Dao) GetTestById(Id int64) (*models.Test, error) {
	var test *models.Test

	err := database.DB.MysqlConn.Model(&models.Test{}).Where("id=?", Id).First(&test).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return test, nil
}
