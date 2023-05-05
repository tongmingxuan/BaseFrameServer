// Package Model /*
package Model

import (
	"gorm.io/gorm"
)

type AdminModel struct {
	BaseModel
	Id            int    `gorm:"primaryKey" json:"id"`
	AdminName     string `json:"admin_name"`
	AdminPassword string `json:"admin_password"`
}

func (m AdminModel) TableName() string {
	return "admin"
}

func (m AdminModel) GetConnection() string {
	return GetDefaultConnectionName()
}

func (m AdminModel) GetDb() *gorm.DB {
	return m.Model(m)
}
