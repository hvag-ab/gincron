package control

import (
	"github.com/gin-gonic/gin"
	"myproject/control/project"
)

// RegisterRouter 注册路由
func RegisterRouter(router *gin.RouterGroup) {

	router.GET("/joblist",project.List)
	router.POST("/addjob",project.Add)
	router.POST("/editjob",project.Edit)
	router.POST("/startjob",project.Start)
	router.POST("/pausejob",project.Pause)
	router.POST("/deljob",project.Deljob)
	router.POST("/runjob",project.Run)
	router.POST("/batchjob",project.Batch)

	router.POST("/deljoblogs",project.LogBatch)
	router.GET("/logs",project.Logs)
}