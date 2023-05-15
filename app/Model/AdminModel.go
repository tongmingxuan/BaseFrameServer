// Package Model /*
package Model

import (
	"BaseFrameServer/app/Constant"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
)

type AdminModel struct {
	tmxServer.BaseModel
	Id            int    `gorm:"primaryKey" json:"id"`
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
}

func (m AdminModel) TableName() string {
	return "admin"
}

func (m AdminModel) GetPollName() string {
	return Constant.DefaultDatabaseConnection
}

func (m AdminModel) GetConnection() *gorm.DB {
	return m.Model(m)
}
