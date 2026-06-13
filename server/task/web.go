package task

import (
	"Helios/common"
	"Helios/initialize"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Web 服务
func StartWebServer() {
	// 监听地址
	listenAddress := fmt.Sprintf("%s:%d", common.Config.System.Listen.Host, common.Config.System.Listen.Port)
	log.Printf("WebServer 服务启动监听地址：%s", listenAddress)

	// 路由初始化
	r := initialize.Router()

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    listenAddress,
		Handler: r,
	}

	// 启动 WebServer 服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("WebServer 服务启动失败：%s", err.Error())
			os.Exit(1)
		}
	}()

	// 监听信号，用于关闭 WebServer 服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 接收到关闭信号后，优雅关闭 WebServer 服务
	timeout := time.Duration(common.Config.System.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("WebServer 服务关闭失败：%s", err.Error())
		os.Exit(1)
	}
	log.Printf("WebServer 服务关闭成功")
}
