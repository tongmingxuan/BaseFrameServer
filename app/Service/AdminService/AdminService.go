// Package AdminService /*
package AdminService

import (
	"BaseFrameServer/app/Common/Utils"
	"BaseFrameServer/app/Constant"
	"BaseFrameServer/app/Dao"
	"BaseFrameServer/app/Model"
	"BaseFrameServer/app/Service"
	"github.com/gin-gonic/gin"
	"github.com/syyongx/php2go"
	"github.com/tongmingxuan/tmx-server/plugin/pluginList/HelperFunction"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
	"time"
)

type AdminService struct {
	Service.BaseService
}

func (service AdminService) Login(adminName, adminPassword string) Service.Result {

	dao := tmxServer.CommonGetDao("")

	adminDao := Dao.GetAdminDao(dao)

	adminInfo, err := adminDao.Login(adminName, adminPassword)

	if err != nil {
		return service.ServiceError("账号或密码错误|数据不存在", adminInfo, 500)
	}

	token := service.setToken(adminInfo)

	return service.ServiceSuccess("登录成功", token)
}

// setToken
// @Description: 设置token
// @receiver service
// @param adminInfo
func (service AdminService) setToken(adminInfo Model.AdminModel) string {
	redis := Utils.GetRedisConnection("default")

	key := Constant.LoginTokenKey + HelperFunction.GetSnowflakeIdByInt64()

	redis.Set(key, adminInfo.Id, 86400*time.Second)

	return key
}

// TokenGetAdmin
// @Description: 获取登录人信息
// @receiver service
// @param c
// @return Result
func (service AdminService) TokenGetAdmin(c *gin.Context) Service.Result {

	param := HelperFunction.GetGinHttpParams(c)

	token, ok := param["token"]

	if ok != true {
		return service.ServiceError("TokenGetAdmin:异常:token_key不存在", nil, 501)
	}

	tokenStr, ok := token.(string)

	if ok == false {
		return service.ServiceError("TokenGetAdmin:异常:断言失败", nil, 501)
	}

	if php2go.Empty(tokenStr) {
		return service.ServiceError("TokenGetAdmin:异常:token不存在", nil, 501)
	}

	redis := Utils.GetRedisConnection("default")

	tokenStr = Constant.LoginTokenKey + tokenStr

	cmd := redis.Get(tokenStr)

	adminId := cmd.Val()

	if adminId == "" {
		return service.ServiceError("TokenGetAdmin:异常:token已过期", nil, 501)
	}

	info, err := Dao.AdminDao{}.FindInfoByWhere(map[string]string{
		"id": adminId,
	})

	if err != nil && err != gorm.ErrRecordNotFound {
		return service.ServiceError("TokenGetAdmin:异常:"+err.Error(), nil, 501)
	}

	if info.Id == 0 {
		return service.ServiceError("TokenGetAdmin:管理员不存在", nil, 501)
	}

	return service.ServiceSuccess("查询成功", map[string]string{
		"admin_name":     info.AdminName,
		"admin_password": info.AdminPassword,
		"admin_id":       adminId,
	})
}
