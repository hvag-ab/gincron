package main

import (
	"myproject/db"
	"myproject/pkg/logging"
	"myproject/routers"
	"myproject/pkg/jobs"
)

//系统初始化
func init() {

	logging.Setup()
	// //mysql配置
	db.Setup()
	// //迁移数据
	// migrate()
	jobs.InitJobs()


}

//系统启动项
func main() {
	
	//初始化路由
	routersInit := routers.InitRouter()
	routersInit.Run("localhost:7890")


}

