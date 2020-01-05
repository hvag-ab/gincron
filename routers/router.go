package routers

import (
	"github.com/gin-gonic/gin"
	"myproject/pkg/setting"

	"myproject/control"

)


// InitRouter 初始化路由
func InitRouter() *gin.Engine {

	router := gin.New()

	
	setUpConfig(router)
	setUpRouter(router)

	return router

}


// 初始化应用设置
func setUpConfig(router *gin.Engine) {
	// 设置静态文件处理


	//Logger实例将日志写入gin.DefaultWriter的日志记录器中间件。
	router.Use(gin.Logger())

	//Recovery返回一个中间件，该中间件从任何恐慌中恢复，如果有500，则写入500。
	router.Use(gin.Recovery())
	//设置mode-----"debug","release","test"
	gin.SetMode(setting.RunMode)

	
}



// 设置路由
func setUpRouter(router *gin.Engine) {
	api := router.Group("/hvag")
	{
		control.RegisterRouter(api)
	}
}


