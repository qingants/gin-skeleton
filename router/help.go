package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qingants/gin-skeleton/pkg/bininfo"
	"net/http"
)

func help(app *gin.Engine) {
	help := app.Group("/help")
	{
		// 健康检查
		help.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		// 服务信息
		help.GET("/version", func(c *gin.Context) {
			c.String(http.StatusOK, string(bininfo.PrettyVersion()))
		})

		// 路由信息
		help.GET("/routers", func(c *gin.Context) {
			var routers []string
			for _, route := range app.Routes() {
				routers = append(routers, fmt.Sprintf("%-6s %-25s --> %s", route.Method, route.Path, route.Handler))
			}
			c.JSON(http.StatusOK, gin.H{
				"routers": routers,
			})
		})

		// 返回客户端IP
		help.GET("/ip", func(c *gin.Context) {
			c.String(http.StatusOK, c.ClientIP())
		})
	}
}
