package common

import (
	"embed"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 参数初始化配置，用于命令行参数初始化
var (
	// 服务默认监听地址
	ParamSystemListenHost = ""
	// 服务默认监听端口，默认0
	ParamSystemListenPort = -1
	// 服务默认配置文件，默认空字符串，不能设置默认值
	ParamSystemConfigFile = ""
	// 数据表名称，默认 all
	ParamSystemMigrateTableName = "all"
	// 服务是否参与领导者选举, 0-不参与，1-参与
	ParamSystemRoleLeader = -1
	// 服务是否开启工作节点角色, 0-不开启，1-开启
	ParamSystemRoleWorker = -1
	// 服务是否开启Web后端服务角色, 0-不开启，1-开启
	ParamSystemRoleWeb = -1
)

// 全局工具变量
var FS embed.FS                  // 全局打包静态文件
var Config *Configuration        // 配置文件解析保存数据
var SystemLog *zap.SugaredLogger // 系统日志
var AccessLog *zap.SugaredLogger // 访问日志
var DB *gorm.DB                  // MySQL 连接
var Cache *redis.Client          // Redis 连接
