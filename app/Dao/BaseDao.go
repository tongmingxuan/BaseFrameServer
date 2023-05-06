// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/DaoPlugin"
)

type InterfaceDao interface {
	GetModel() Model.InterfaceModel
}

type BaseDao struct {
	DaoPlugin.DaoPlugin
}
