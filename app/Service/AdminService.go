// Package Service /*
package Service

import "BaseFrameServer/app/Dao"

type AdminService struct {
}

func (s AdminService) Login(adminName, adminPassword string) {
	Dao.GetAdminDao().FindInfoByWhere([]interface{}{
		[]interface{}{"admin_name", "=", adminName},
		[]interface{}{"admin_password", "=", adminPassword},
	})
}
