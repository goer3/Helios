package initialize

import (
	"Helios/common"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 初始化 MySQL 数据库连接
func MySQL() {
	// 数据库连接串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%dms&%s",
		common.Config.MySQL.Username,
		common.Config.MySQL.Password,
		common.Config.MySQL.Host,
		common.Config.MySQL.Port,
		common.Config.MySQL.Database,
		common.Config.MySQL.Timeout,
		common.Config.MySQL.Params)

	// 创建配置
	var c = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
			// TablePrefix:   "tb_", // 表名前缀
		},
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键
		IgnoreRelationshipsWhenMigrating:         false, // 开启会导致 many2many 的表创建失败
		QueryFields:                              true,  // 解决查询索引失效问题
		// TranslateError:                           true,  // 启用错误翻译
	}

	// 如果日志级别是 DEBUG，则打印所有 SQL
	if common.Config.Log.System.Level == "debug" {
		c.Logger = logger.Default.LogMode(logger.Info)
	}

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 170, // varchar 默认长度，太长影响查询
	}), c)

	// 错误处理
	if err != nil {
		common.SystemLog.Fatal("MySQL 数据库初始化连接失败: ", err)
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(common.Config.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(common.Config.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(common.Config.MySQL.MaxIdleTime) * time.Minute)

	// 设置全局数据库连接，方便后续使用
	common.DB = db

	// 初始化完成日志
	logDsn := fmt.Sprintf("%s@%s:%d/%s",
		common.Config.MySQL.Username,
		common.Config.MySQL.Host,
		common.Config.MySQL.Port,
		common.Config.MySQL.Database)
	common.SystemLog.Info("MySQL 数据库初始化连接成功: ", logDsn)
}
