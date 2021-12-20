package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/requestid"
	//_ "github.com/qingants/gin-skeleton/doc"
	"github.com/qingants/gin-skeleton/middleware"

	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
)

//var swagHandler gin.HandlerFunc
//
//func init() {
//	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
//}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	// Creates a router without any middleware by default
	app := gin.New()
	// Global middleware
	// Logger middleware will write the logs
	app.Use(middleware.LoggerWithZap(zap.L()))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(middleware.RecoveryWithZap(zap.L()))
	app.Use(middleware.CORSMiddleware())

	app.Use(requestid.New())

	// 文档路径
	//app.GET("/v1/api/gin-skeleton/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	help(app)
	movie(app)

	return app
}
