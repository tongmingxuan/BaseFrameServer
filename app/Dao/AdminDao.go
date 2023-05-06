// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
)

type AdminDao struct {
	BaseDao
}

func (dao AdminDao) GetModel() Model.InterfaceModel {
	return Model.AdminModel{}
}

func (dao AdminDao) Login(adminName, adminPassword string) (result Model.AdminModel, err error) {
	result = Model.AdminModel{}
	err = dao.GetModel().GetDb().
		Where("admin_name", adminName).
		Where("admin_password", adminPassword).
		First(&result).Error

	return
}

func (dao AdminDao) FindInfoByWhere(where interface{}) (adminInfo Model.AdminModel, err error) {
	return adminInfo, dao.GetModel().GetDb().Where(where).First(&adminInfo).Error
}
