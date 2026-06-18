package service

import (
	"Helios/common"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// CreateFromDto 将 DTO 通过 copier 转换为 model 并写入数据库。
// 注意：request dto 和 model 必须是指针类型。
func CreateFromDto[T any](db *gorm.DB, dto any) error {
	var entity = new(T)
	if err := copier.Copy(entity, dto); err != nil {
		common.SystemLog.Error("数据转换异常：", err)
		return err
	}
	if err := db.Create(entity).Error; err != nil {
		common.SystemLog.Error("创建数据失败: ", err)
		return err
	}
	return nil
}
