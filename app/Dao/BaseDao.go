// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/DaoPlugin"
	"gorm.io/gorm"
)

type InterfaceDao interface {
	GetModel() Model.InterfaceModel
}

type BaseDao struct {
	DaoPlugin.DaoPlugin
	BindDao InterfaceDao
}

func (dao BaseDao) bindDao(bindDao InterfaceDao) {
	dao.BindDao = bindDao
}

func (dao BaseDao) FindInfoByWhere(where []interface{}) (result Model.InterfaceModel) {

	gormDb := dao.BindDao.GetModel().GetDb()

	table := dao.BindDao.GetModel()

	byWhere, err := dao.BuildFindByWhere(gormDb, where)

	if err != nil {
		panic(err.Error())
	}

	db := byWhere.First(&table)

	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		panic(db.Error.Error())
	}

	if db.RowsAffected == 0 {
		return nil
	}

	return table
}
