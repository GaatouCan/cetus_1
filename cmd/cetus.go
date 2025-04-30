package main

import (
	"demo/internal/handler"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
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

	router := gin.Default()

	// Hello

	helloHandler := handler.HelloHandler{}

	router.GET("/hello", helloHandler.HelloWorld)
	router.GET("/user", helloHandler.GetUser)
	router.POST("/user", helloHandler.CreateUser)
	router.PUT("/user", helloHandler.UpdateUser)
	router.DELETE("/user", helloHandler.DeleteUser)

	// GroceryItem

	groceryItemHandler := handler.GroceryItemHandler{}

	router.GET("/groceryItem", groceryItemHandler.GetGroceryItems)
	router.POST("/groceryItem", groceryItemHandler.CreateGroceryItem)

	// 创建服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 启动服务在 goroutine 中以便监听 OS 信号
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号触发
	<-ctx.Done()
	stop()
	log.Print("shutting down server")

	// 创建一个 5 秒超时的 context, 用于 graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 尝试优雅关闭 HTTP 服务器: 停止接受新请求 等待旧请求完成
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Print("Server exiting")
}
