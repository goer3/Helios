package initialize

import (
	"Helios/common"
	"Helios/model"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// sqlCaptureLogger 包装 gorm 的 Logger，捕获 DDL SQL 写入文件
type sqlCaptureLogger struct {
	logger.Interface
	sb        strings.Builder
	hasChange bool
}

func (l *sqlCaptureLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	upper := strings.ToUpper(strings.TrimSpace(sql))
	// 只捕获 CREATE TABLE / ALTER TABLE / DROP TABLE
	if strings.HasPrefix(upper, "CREATE TABLE") || strings.HasPrefix(upper, "ALTER TABLE") || strings.HasPrefix(upper, "DROP TABLE") {
		l.sb.WriteString(sql)
		l.sb.WriteString(";\n")
		l.hasChange = true
	}
	l.Interface.Trace(ctx, begin, fc, err)
}

// 查询指定表的当前列名集合
func getTableColumns(db *gorm.DB, tableName string) (map[string]bool, error) {
	var columns []struct {
		ColumnName string `gorm:"column:COLUMN_NAME"`
	}
	err := db.Raw(
		"SELECT COLUMN_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ?",
		tableName,
	).Scan(&columns).Error
	if err != nil {
		return nil, err
	}
	colSet := make(map[string]bool, len(columns))
	for _, c := range columns {
		colSet[c.ColumnName] = true
	}
	if len(colSet) == 0 {
		return nil, nil
	}
	return colSet, nil
}

// tableNameOf 获取 model 的表名
func tableNameOf(m any) string {
	if t, ok := m.(interface{ TableName() string }); ok {
		return t.TableName()
	}
	return "unknown"
}

// diffColumns 对比列变更，返回新增列和删除列
func diffColumns(before, after map[string]bool) (added, removed []string) {
	for col := range after {
		if !before[col] {
			added = append(added, col)
		}
	}
	for col := range before {
		if !after[col] {
			removed = append(removed, col)
		}
	}
	return
}

// 数据表结构同步
func Table() {
	// 包装 GORM Logger 用于捕获 DDL SQL
	oldLogger := common.DB.Logger
	capture := &sqlCaptureLogger{Interface: oldLogger}
	common.DB.Logger = capture
	defer func() { common.DB.Logger = oldLogger }()

	models := []any{
		&model.SystemSetting{},
		&model.SystemRole{},
		&model.SystemMenu{},
		&model.SystemApiCategory{},
		&model.SystemApi{},
		&model.SystemUser{},
		&model.SystemUserExtraApi{},
		&model.SystemUserExtraMenu{},
	}

	for _, m := range models {
		tableName := tableNameOf(m)

		// 用 HasTable 判断表是否真实存在（区别于 GORM 前置创建导致的列被查到）
		tableExists := common.DB.Migrator().HasTable(tableName)

		// 迁移前记录当前列
		before, _ := getTableColumns(common.DB, tableName)

		// 执行迁移
		if err := common.DB.AutoMigrate(m); err != nil {
			common.SystemLog.Error("[失败] " + tableName + "：" + err.Error())
			continue
		}

		// 迁移后记录当前列
		after, _ := getTableColumns(common.DB, tableName)
		if after == nil {
			continue
		}
		if !tableExists {
			common.SystemLog.Info("[创建] " + tableName)
		} else {
			added, removed := diffColumns(before, after)
			if len(added) > 0 || len(removed) > 0 {
				parts := []string{"[变更] " + tableName + ","}
				if len(added) > 0 {
					parts = append(parts, "新增列（"+strings.Join(added, ", ")+"）")
				}
				if len(removed) > 0 {
					parts = append(parts, "删除列（"+strings.Join(removed, ", ")+"）")
				}
				common.SystemLog.Info(strings.Join(parts, " "))
			} else {
				common.SystemLog.Info("[跳过] " + tableName)
			}
		}
	}

	common.SystemLog.Info("数据表同步完成")
	// 如果有 DDL 变更，写入 update.sql（仅 debug 模式）
	if capture.hasChange && common.Config.System.Mode == "debug" {
		now := time.Now().Format(common.TIME_SECOND)
		header := fmt.Sprintf("\n-- Version: %s | Update at: %s\n", common.PROJECT_VERSION, now)
		content := header + capture.sb.String()

		// 确保 sql 目录存在
		os.MkdirAll("sql", 0755)

		// 文件不存在则创建，存在则追加
		f, err := os.OpenFile("sql/update.sql", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			common.SystemLog.Error("打开 sql/update.sql 文件失败：" + err.Error())
		} else {
			defer f.Close()
			if _, err := f.WriteString(content); err != nil {
				common.SystemLog.Error("写入 sql/update.sql 文件失败：" + err.Error())
			} else {
				common.SystemLog.Info("sql/update.sql 文件写入成功")
			}
		}
	}
}
