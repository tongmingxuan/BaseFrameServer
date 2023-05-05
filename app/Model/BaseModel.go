// Package Model /*
package Model

import (
	"BaseFrameServer/app/Common/Utils"
	"BaseFrameServer/app/Constant"
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// InterfaceModel
// @Description: 基础模型接口
type InterfaceModel interface {
	//TableName 获取当前表名称
	TableName() string
	//GetConnection 返回默认mysql连接池Name
	GetConnection() string
	//GetDb 获取当前模型
	GetDb() *gorm.DB
}

// Connection
// @Description: 获取数据库连接实例
// @param connection
// @return *gorm.DB
func Connection(connection string) *gorm.DB {
	dbPoll := Utils.ContainerGetDataBase()

	conn, ok := dbPoll.Pool.Load(connection)

	if ok != true {
		panic("dbPoll.Pool.Load(connection):error")
	}

	return conn.(*gorm.DB)
}

// GetDefaultConnectionName
// @Description: 默认数据库连接池名称
// @return string
func GetDefaultConnectionName() string {
	return Constant.DefaultDatabaseConnection
}

// BaseModel
// @Description: 基础模型类
type BaseModel struct {
	CreatedAt MyTime `json:"created_at"`
	UpdatedAt MyTime `json:"updated_at"`
	DeletedAt MyTime `json:"deleted_at"`
}

// Model
// @Description: 获取模型实例
// @receiver m
// @param model
// @return *gorm.DB
func (m BaseModel) Model(model InterfaceModel) *gorm.DB {
	connection := model.GetConnection()

	if connection == "" {
		panic("GetModel:model.GetConnection():返回空")
	}

	conn := Connection(connection)

	return conn.Model(model)
}

type MyTime struct {
	time.Time
}

// MarshalJSON  重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t MyTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// Value 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
