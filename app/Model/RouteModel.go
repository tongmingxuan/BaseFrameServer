// Package Model /*
package Model

import (
	"BaseFrameServer/app/Constant"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
)

type RouteModel struct {
	tmxServer.BaseModel
	Id             int    `gorm:"primaryKey" json:"id"`
	MainRouteId    int    `json:"main_route_id"`
	UpRouteId      int    `json:"up_route_id"`
	RouteName      string `json:"route_name"`
	Describe       string `json:"describe"`
	QueueType      string `json:"queue_type"`
	QueueInfo      string `json:"queue_info"`
	OpenStatus     int    `json:"open_status"`
	RouteLevelType string `json:"route_level_type"`
}

func (m RouteModel) TableName() string {
	return "sync_route"
}

func (m RouteModel) GetPollName() string {
	return Constant.DefaultDatabaseConnection
}

func (m RouteModel) GetConnection() *gorm.DB {
	return m.Model(m)
}
