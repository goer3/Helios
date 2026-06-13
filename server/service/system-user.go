package service

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/utils"
	"errors"

	"github.com/dromara/carbon/v2"
	"github.com/jinzhu/copier"
)

// 创建用户
func SystemUserCreate(req *dto.SystemUserCreateRequest) error {
	// 验证过期时间格式
	var expireAt *carbon.Carbon
	if req.ExpireAt != "" {
		c := carbon.ParseByLayout(req.ExpireAt, common.TIME_DATE, carbon.Local)
		if c == nil || c.IsInvalid() {
			return errors.New("过期时间格式不正确")
		}
		expireAt = c
	} else {
		// 默认过期时间为 1000 年
		expireAt = carbon.Now().AddYears(1000)
	}

	// 密码加密
	hashedPassword, err := utils.PasswordEncrypt(req.Password)
	if err != nil {
		return errors.New("密码加密失败：" + err.Error())
	}

	// 数据转换
	var user = model.SystemUser{}
	if err := copier.Copy(&user, &req); err != nil {
		common.SystemLog.Error("创建用户数据转换异常：", err.Error())
		return errors.New("创建用户数据转换异常：" + err.Error())
	}

	// 设置密码和过期时间
	user.Password = hashedPassword
	user.ExpireAt = expireAt

	// 创建用户
	if err := common.DB.Create(&user).Error; err != nil {
		common.SystemLog.Error("创建用户失败：", err.Error())
		return errors.New("创建用户失败：" + err.Error())
	}

	return nil
}
