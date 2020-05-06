package routers

import (
	"git.trac.cn/nv/spider/pkg/app"
	"git.trac.cn/nv/spider/pkg/e"
	"git.trac.cn/nv/spider/pkg/setting"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net"
	"net/http"
	"os"
	"strings"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(registerRecovery())

	if setting.ServerSetting.RunMode == gin.DebugMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/spider", Episode)
	}

	return r
}

func registerRecovery() gin.HandlerFunc {
	return recovery()
}

func recovery() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					c.Error(err.(error))
					c.Abort()
				} else {
					g := app.Gin{C: c}
					g.EpisodeResponse(http.StatusInternalServerError, e.ERROR, nil, nil)
				}
			}
		}()
		c.Next()
	}
}
