package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"Helios/common"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志级别映射
var zapLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

// 自定义时间格式编码器
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(common.TIME_MILLISECOND))
}

// 创建日志写入器
func createLogWriter(cfg common.BaseLogConfiguration) zapcore.WriteSyncer {
	var writers []zapcore.WriteSyncer

	// 未启用日志时，仅输出到控制台
	if !cfg.Enabled {
		return zapcore.AddSync(os.Stdout)
	}

	// 确保日志目录存在
	if err := os.MkdirAll(cfg.Path, 0755); err != nil {
		log.Fatalf("创建日志目录失败：%s", err.Error())
	}

	// 构建日志文件名
	filename := fmt.Sprintf("%s/%s.log", cfg.Path, cfg.Prefix)

	// 配置日志文件写入器
	hook := &lumberjack.Logger{
		Filename: filename,
	}

	// 如果开启滚动，配置日志切割参数
	if cfg.Rolling.Enabled {
		hook.MaxSize = cfg.Rolling.MaxSize
		hook.MaxAge = cfg.Rolling.MaxAge
		hook.MaxBackups = cfg.Rolling.MaxBackups
		hook.Compress = cfg.Rolling.Compress
	}

	// 添加文件写入器
	writers = append(writers, zapcore.AddSync(hook))

	// 默认同时输出到控制台
	writers = append(writers, zapcore.AddSync(os.Stdout))
	return zapcore.NewMultiWriteSyncer(writers...)
}

// 创建日志记录器
func NewLogger(cfg common.BaseLogConfiguration) *zap.SugaredLogger {
	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建编码器
	var encoder zapcore.Encoder
	if cfg.Formatter == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 获取日志级别
	level, ok := zapLevelMap[cfg.Level]
	if !ok {
		level = zapcore.InfoLevel // 默认 INFO 级别
	}

	// 创建日志核心
	core := zapcore.NewCore(
		encoder,
		createLogWriter(cfg),
		level,
	)

	// 创建并返回日志记录器
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

// 系统日志初始化
func SystemLogger() {
	common.SystemLog = NewLogger(common.Config.Log.System)
}

// 访问日志初始化
func AccessLogger() {
	common.AccessLog = NewLogger(common.Config.Log.Access)
}
