package main

import (
	"blog/config"
	"blog/router"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {
	fmt.Println("gin.Version: ", gin.Version)
	if config.Mode == config.ProductionMode {
		gin.SetMode(gin.ReleaseMode)
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.
		logFile, err := os.OpenFile(config.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(-1)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	}

	// Creates a router without any middleware by default
	app := gin.New()

	if config.Mode != config.ProductionMode {
		// Dev环境的时候,前端单独占用一个端口, 配置一下跨域, 有时候配置ｎｇｉｎｘ不如去复制代码简单的
		app.Use(cors.New(cors.Config{
			// 这里注意一下,nginx部署前端模板的情况下必须是从nginx请求过来的端口
			AllowOrigins:     []string{"http://192.168.238.133:8080"}, // 不能写成:http://192.168.238.133:8080/
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Origin","Content-Type", "X-Requested-With", "token", "Authorization", "Cookie"},
			AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		}))
	}

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())
	router.Route(app)

	app.Run(":" + fmt.Sprintf("%d", config.Port))
}
