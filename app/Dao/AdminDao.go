// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
)

func GetAdminDao() AdminDao {
	AdminDao := AdminDao{}
	AdminDao.BindDao = AdminDao
	return AdminDao
}

type AdminDao struct {
	BaseDao
}

func (dao AdminDao) GetModel() Model.InterfaceModel {
	return Model.AdminModel{}
}
