package initialize

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/utils"
	"errors"
	"fmt"

	"github.com/dromara/carbon/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 通用 upsert 函数：如果记录不存在则创建，存在则跳过
func upsert[T any](db *gorm.DB, rows []T, label string) {
	// 临时静默 GORM 日志，避免 First 查询 ErrRecordNotFound 的日志噪音
	oldLogger := db.Logger
	db.Logger = oldLogger.LogMode(logger.Silent)
	defer func() { db.Logger = oldLogger }()

	for _, row := range rows {
		// 通过反射获取 Id 字段的值用于查询
		var id uint
		var name string
		switch v := any(row).(type) {
		case model.SystemMenu:
			id = v.Id
			name = v.Name
		case model.SystemRole:
			id = v.Id
			name = v.Name
		case model.SystemUser:
			id = v.Id
			name = v.Nickname
		case model.SystemApiCategory:
			id = v.Id
			name = v.Name
		case model.SystemApi:
			id = v.Id
			name = v.Name
		}

		var existing T
		if err := db.Where("id = ?", id).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err = db.Create(&row).Error; err != nil {
					common.SystemLog.Error(fmt.Sprintf("[失败] 添加%s（%s）：%s", label, name, err.Error()))
					return
				} else {
					common.SystemLog.Info(fmt.Sprintf("[成功] %s（%s）添加成功", label, name))
					continue
				}
			} else {
				common.SystemLog.Error(fmt.Sprintf("[失败] 查询%s（%s）：%s", label, name, err.Error()))
				return
			}
		} else {
			common.SystemLog.Info(fmt.Sprintf("[跳过] %s（%s）已存在", label, name))
		}
	}
}

// 执行初始化 SQL 脚本 sql/init.sql
func Data() {
	initSystemMenuData()
	initSystemRoleData()
	initSystemUserData()
	initSystemApiCategoryData()
	initSystemApiData()
}

// 菜单初始化
func initSystemMenuData() {
	common.SystemLog.Info("开始初始化系统菜单")
	menus := []model.SystemMenu{
		{Id: 1, Name: "工作空间", Path: "/dashboard", Icon: "DesktopOutlined", ParentId: 0},
		{Id: 10000, Name: "数据查询", Path: "/query", Icon: "FundOutlined", ParentId: 0},
		{Id: 20000, Name: "数据来源", Path: "/datasource", Icon: "HddOutlined", ParentId: 0},
		{Id: 30000, Name: "监控告警", Path: "/alarm", Icon: "AlertOutlined", ParentId: 0},
		{Id: 31000, Name: "告警事件", Path: "/alarm/events", Icon: "ExceptionOutlined", ParentId: 30000},
		{Id: 32000, Name: "告警策略", Path: "/alarm/rules", Icon: "BorderlessTableOutlined", ParentId: 30000},
		{Id: 33000, Name: "屏蔽策略", Path: "/alarm/shielding", Icon: "AudioMutedOutlined", ParentId: 30000},
		{Id: 34000, Name: "告警历史", Path: "/alarm/history", Icon: "FieldTimeOutlined", ParentId: 30000},
		{Id: 34100, Name: "等待处理", Path: "/alarm/history/todo", Icon: "WarningTwoTone", ParentId: 34000},
		{Id: 34200, Name: "完成处理", Path: "/alarm/history/finish", Icon: "CheckOutlined", ParentId: 34000},
		{Id: 40000, Name: "消息通知", Path: "/message", Icon: "MailOutlined", ParentId: 0},
		{Id: 41000, Name: "告警媒介", Path: "/message/type", Icon: "AppstoreAddOutlined", ParentId: 40000},
		{Id: 42000, Name: "通知模板", Path: "/message/template", Icon: "SnippetsOutlined", ParentId: 40000},
		{Id: 50000, Name: "人员组织", Path: "/user", Icon: "UsergroupAddOutlined", ParentId: 0},
		{Id: 51000, Name: "用户列表", Path: "/user/list", Icon: "SolutionOutlined", ParentId: 50000},
		{Id: 52000, Name: "用户分组", Path: "/user/groups", Icon: "CommentOutlined", ParentId: 50000},
		{Id: 53000, Name: "项目团队", Path: "/user/projects", Icon: "ClusterOutlined", ParentId: 50000},
		{Id: 54000, Name: "人员排班", Path: "/user/scheduling", Icon: "CalendarOutlined", ParentId: 50000},
		{Id: 90000, Name: "系统设置", Path: "/system", Icon: "SlidersOutlined", ParentId: 0},
		{Id: 91000, Name: "角色授权", Path: "/system/role", Icon: "IdcardOutlined", ParentId: 90000},
		{Id: 92000, Name: "菜单配置", Path: "/system/menu", Icon: "SisternodeOutlined", ParentId: 90000},
		{Id: 93000, Name: "接口配置", Path: "/system/api", Icon: "ApiOutlined", ParentId: 90000},
		{Id: 94000, Name: "通用设置", Path: "/system/setting", Icon: "SettingOutlined", ParentId: 90000},
		{Id: 99000, Name: "系统日志", Path: "/system/log", Icon: "AuditOutlined", ParentId: 0},
	}
	upsert(common.DB, menus, "菜单")
}

// 角色初始化
func initSystemRoleData() {
	common.SystemLog.Info("开始初始化系统角色")
	roles := []model.SystemRole{
		{BaseModel: model.BaseModel{Id: 1}, Name: "超级管理员", Description: "系统预设最高权限角色，无法删除，修改"},
		{BaseModel: model.BaseModel{Id: 2}, Name: "访客", Description: "系统预设访客权限角色，无法删除，修改，仅能查看系统基本信息"},
	}
	upsert(common.DB, roles, "角色")
}

// 用户初始化
func initSystemUserData() {
	common.SystemLog.Info("开始初始化系统用户")
	password, err := utils.PasswordEncrypt("helios")
	if err != nil {
		common.SystemLog.Error("[失败] 加密系统用户密码失败：" + err.Error())
		return
	}
	users := []model.SystemUser{
		{
			BaseModel:    model.BaseModel{Id: 1},
			Nickname:     "超管",
			Username:     "helios",
			Mobile:       "18888888888",
			HideMobile:   1,
			Email:        "helios@helios.com",
			Password:     password,
			Gender:       1,
			AvatarUrl:    "https://github.com/shadcn.png",
			ExpireAt:     carbon.Now().AddYears(1000),
			SystemRoleId: 1,
		},
	}
	upsert(common.DB, users, "用户")
}

// API 分类初始化
func initSystemApiCategoryData() {
	common.SystemLog.Info("开始初始化系统API分类")
	apiCategories := []model.SystemApiCategory{
		// 开放接口
		{Id: 1000, Name: "对外开放接口", ParentId: 0},
		// 免授权接口
		{Id: 2000, Name: "免授权接口", ParentId: 0},
		// 系统接口
		{Id: 10000, Name: "系统接口", ParentId: 0},
		{Id: 11000, Name: "用户接口", ParentId: 10000},
		{Id: 12000, Name: "菜单接口", ParentId: 10000},
		{Id: 13000, Name: "角色接口", ParentId: 10000},
		{Id: 14000, Name: "API分类接口", ParentId: 10000},
		{Id: 15000, Name: "API接口", ParentId: 10000},
	}
	upsert(common.DB, apiCategories, "API分类")
}

// API 初始化
func initSystemApiData() {
	common.SystemLog.Info("开始初始化系统API")
	api := []model.SystemApi{
		// 开放接口
		{Id: 1001, Name: "健康检查", Method: "GET", Api: "/openapi/v1/health", IsAuthApi: 0, SystemApiCategoryId: 1000},
		{Id: 1002, Name: "账号密码登录", Method: "POST", Api: "/openapi/v1/login", IsAuthApi: 0, SystemApiCategoryId: 1000},
		{Id: 1003, Name: "钉钉扫码登录", Method: "POST", Api: "/openapi/v1/login/dingtalk", IsAuthApi: 0, SystemApiCategoryId: 1000},
		{Id: 1004, Name: "飞书扫码登录", Method: "POST", Api: "/openapi/v1/login/feishu", IsAuthApi: 0, SystemApiCategoryId: 1000},
		{Id: 1005, Name: "企业微信扫码登录", Method: "POST", Api: "/openapi/v1/login/wechat", IsAuthApi: 0, SystemApiCategoryId: 1000},
		// 免授权接口
		{Id: 2001, Name: "退出登录", Method: "POST", Api: "/api/v1/logout", IsAuthApi: 0, SystemApiCategoryId: 2000},
		// 用户接口
		{Id: 11001, Name: "获取用户列表", Method: "GET", Api: "/api/v1/system/user/list", IsAuthApi: 1, SystemApiCategoryId: 11000},
		{Id: 11002, Name: "创建新用户", Method: "POST", Api: "/api/v1/system/user/create", IsAuthApi: 1, SystemApiCategoryId: 11000},
		// 菜单接口
		{Id: 12001, Name: "获取菜单列表", Method: "GET", Api: "/api/v1/system/menu/list", IsAuthApi: 1, SystemApiCategoryId: 12000},
		// 角色接口
		{Id: 13001, Name: "获取角色列表", Method: "GET", Api: "/api/v1/system/role/list", IsAuthApi: 1, SystemApiCategoryId: 13000},
		// API分类接口
		{Id: 14001, Name: "获取API分类列表", Method: "GET", Api: "/api/v1/system/api-category/list", IsAuthApi: 1, SystemApiCategoryId: 14000},
		// API接口
		{Id: 15001, Name: "获取API接口列表", Method: "GET", Api: "/api/v1/system/api/list", IsAuthApi: 1, SystemApiCategoryId: 15000},
	}
	upsert(common.DB, api, "API")
}
