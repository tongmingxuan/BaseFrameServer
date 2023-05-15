// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func GetAdminDao(d *tmxServer.Dao) *AdminDao {
	return &AdminDao{
		Dao: d,
	}
}

type AdminDao struct {
	*tmxServer.Dao
	tmxServer.BaseDao
}

func (dao AdminDao) GetModel() tmxServer.InterfaceModel {
	return Model.AdminModel{}
}

func (dao AdminDao) Login(adminName, adminPassword string) (result Model.AdminModel, err error) {
	result = Model.AdminModel{}

	//err = dao.GetModel().GetConnection().
	//	Where("admin_name", adminName).
	//	Where("admin_password", adminPassword).
	//	First(&result).Error
	err = dao.GormDb.Model(Model.AdminModel{}).
		Where("admin_name", adminName).
		Where("admin_password", adminPassword).
		First(&result).Error

	return
}

func (dao AdminDao) FindInfoByWhere(where interface{}) (adminInfo Model.AdminModel, err error) {
	return adminInfo, dao.GetModel().GetConnection().Where(where).First(&adminInfo).Error
}
