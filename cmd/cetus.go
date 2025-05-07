package main

import (
	"demo/configs"
	"demo/internal"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 创建 OS 信号通道 监听 Ctrl+C 或 kill 信号
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := configs.GetConfig()

	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移表
	internal.InitDatabaseTables(db)

	r := gin.Default()
	internal.RegisterRouter(r, db)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	// 创建服务器
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 启动服务在 goroutine 中以便监听 OS 信号
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号触发
	<-ctx.Done()
	stop()
	log.Print("Shutting down server")

	// 创建一个 5 秒超时的 context, 用于 graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 尝试优雅关闭 HTTP 服务器: 停止接受新请求 等待旧请求完成
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Print("Server exiting")
}
