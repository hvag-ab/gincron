package project

import (
	libcron "github.com/robfig/cron/v3"
	"myproject/pkg/jobs"
	"myproject/models"
	"time"
	"github.com/gin-gonic/gin"
	"myproject/pkg/app"
	"github.com/Unknwon/com"
	"fmt"
	"strconv"
)



// 任务列表
func List(c *gin.Context) {
	appG := app.Gin{C: c}
	page := com.StrTo(c.DefaultQuery("page","1")).MustInt()
	status := com.StrTo(c.DefaultQuery("status","1")).MustInt()

	cond := make(map[string]interface{})
	cond["status"] = status

	if page < 1 {
		page = 1
	}

	result,paginatorMap,errr := models.TaskGetList(page,cond)
	if errr != nil {
		appG.ErrorResponse(404, "查询任务失败")
		return
	}
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.TaskName
		row["cron_spec"] = v.CronSpec
		row["status"] = v.Status
		row["description"] = v.Description

		e := jobs.GetEntryById(v.Id)
		if e != nil {
			row["next_time"] = e.Next.Format("2006-01-02 15:04:05")
			row["prev_time"] = "-"
			if e.Prev.Unix() > 0 {
				row["prev_time"] = e.Prev.Format("2006-01-02 15:04:05")
			} else if v.PrevTime > 0 {
				row["prev_time"] = time.Unix(v.PrevTime, 0).Format("2006-01-02 15:04:05")
			}
			row["running"] = 1
		} else {
			row["next_time"] = "-"
			if v.PrevTime > 0 {
				row["prev_time"] = time.Unix(v.PrevTime, 0).Format("2006-01-02 15:04:05")
			} else {
				row["prev_time"] = "-"
			}
			row["running"] = 0
		}
		list[k] = row
	}
	paginatorMap["data"] = list
	appG.Response(200,paginatorMap)
}

// 添加任务
func Add(c *gin.Context) {
	appG := app.Gin{C: c}

	task := new(models.Task)

	err := c.ShouldBind(task)
	if err != nil{
		appG.ErrorResponse(404, err.Error())
		return 
	}

	if _, err := libcron.ParseStandard(task.CronSpec); err != nil {
		appG.ErrorResponse(404, "cron表达式无效")
		return 
	}
	if _, b := models.TaskAdd(task); b == false {
		appG.ErrorResponse(404, "添加任务失败")
		return 
	}

	appG.Response(200,"添加任务成功")
	

}

// 编辑任务
func Edit(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).MustInt()

	task, err := models.TaskGetById(id)
	if err != nil {
		appG.ErrorResponse(404, fmt.Sprintf("id:[%d]不存在",id))
		return 
	}
	if task == nil {
		appG.ErrorResponse(404, "此任务不存在")
		return 
	}

	jobs.RemoveJob(id)

	task_name := c.DefaultPostForm("task_name", "")
	if task_name != "" {
		task.TaskName = task_name
	}
	description := c.DefaultPostForm("description", "")
	if description != "" {
		task.Description = description
	}
	cron_spec := c.DefaultPostForm("cron_spec", "")
	if cron_spec != "" {
		task.CronSpec = cron_spec
	}
	command := c.DefaultPostForm("command", "")
	if command != "" {
		task.Command = command
	}
	timeout := c.DefaultPostForm("timeout", "")
	if timeout != "" {
		to,error := strconv.Atoi(timeout)
		if error != nil{
			appG.ErrorResponse(404, "timeout必须为一个整数")
			return 
		}
		task.Timeout = to
	}
	task_type := c.DefaultPostForm("task_type", "")
	if task_type != "" {
		ttype,error := strconv.Atoi(task_type)
		if error != nil{
			appG.ErrorResponse(404, "task_type必须为一个整数")
			return 
		}
		task.TaskType = ttype
	}

	if task.CronSpec !=""{
		if _, err := libcron.ParseStandard(task.CronSpec); err != nil {
			appG.ErrorResponse(404, "cron表达式无效")
			return 
		}
	}

	if b := task.Update(); b == false {
		appG.ErrorResponse(404, "编辑任务失败")
		return 
	}

	appG.Response(200,"编辑任务成功 请重新启动任务")
	
}



type BatchJson struct {
	Action     string `form:"action" json:"action"  binding:"required"`
    Ids []int `form:"ids" json:"ids" binding:"required"`
}
// // 批量操作
func Batch(c *gin.Context) {
	appG := app.Gin{C: c}
	var json BatchJson
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.ErrorResponse(404, "批量操作传入json失败")
		return 
	}
	for _, id := range json.Ids {

		if id < 1 {
			continue
		}
		switch json.Action {
		case "active":
			if task, err := models.TaskGetById(id); err == nil && task != nil {
				job, err := jobs.NewJobFromTask(task)
				if err == nil {
					jobs.AddJob(task.CronSpec, job)
					task.Status = 1
					task.Update()
				}
			}else{
				appG.Response(200,"批量操作任务失败")
				return
			}
		case "pause":
			if task, err := models.TaskGetById(id); err == nil && task != nil {
				task.Status = 0
				task.Update()
				jobs.RemoveJob(id)
			}else{
				appG.Response(200,"批量操作任务失败")
				return 
			}
			
		case "delete":
			if affected, err := models.TaskDel(id); err == nil && affected != 0 {
				models.TaskLogDelByTaskId(id)
				jobs.RemoveJob(id)
			}else{
				appG.Response(200,"批量操作任务失败")
				return 
			}
		}
	}
	appG.Response(200,"批量操作任务成功")
}


// 启动任务
func Start(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).MustInt()

	task, err := models.TaskGetById(id)
	if err != nil {
		appG.ErrorResponse(404, "启动失败")
		return 
	}
	if task == nil {
		appG.ErrorResponse(404, "此任务不存在")
		return 
	}

	job, err := jobs.NewJobFromTask(task)
	if err != nil {
		appG.ErrorResponse(404, "启动任务失败")
		return 
	}

	if jobs.AddJob(task.CronSpec, job) {
		task.Status = 1
		task.Update()
	}

	appG.Response(200,"启动任务成功")
}

// 暂停任务
func Pause(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).MustInt()
	task, err := models.TaskGetById(id)
	if err != nil {
		appG.ErrorResponse(404, "暂停任务失败")
		return 
	}
	if task == nil {
		appG.ErrorResponse(404, "此任务不存在")
		return 
	}

	jobs.RemoveJob(id)
	task.Status = 0
	task.Update()

	appG.Response(200,"暂停任务成功")
}


// 删除任务
func Deljob(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).MustInt()

	affected, err := models.TaskDel(id)
	if err != nil {
		appG.ErrorResponse(404, "删除任务失败")
		return 
	}
	if affected == 0 {
		appG.ErrorResponse(404, "此任务不存在 删除任务失败")
		return 
	}

	err2 := models.TaskLogDelByTaskId(id)
	if err2 != nil {
		fmt.Println(err2.Error())
		appG.ErrorResponse(404, "删除任务日志失败")
		return 
	}

	jobs.RemoveJob(id)

	appG.Response(200,"暂停任务成功")
}


// 立即执行
func Run(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.PostForm("id")).MustInt()

	task, err := models.TaskGetById(id)
	if err != nil {
		appG.ErrorResponse(404, "立即执行任务失败")
		return 
	}

	job, err := jobs.NewJobFromTask(task)
	if err != nil {
		appG.ErrorResponse(404, "立即执行任务失败")
		return 
	}

	job.Run()

	appG.Response(200,"立即执行任务成功")
}


// // 日志查询列表
func Logs(c *gin.Context) {
	appG := app.Gin{C: c}
	page := com.StrTo(c.DefaultQuery("page","1")).MustInt()
	task_id := com.StrTo(c.DefaultQuery("task_id","1")).MustInt()
	cond := make(map[string]interface{})
	cond["task_id"] = task_id
	if page < 1 {
		page = 1
	}

	result,paginatorMap,errr := models.TaskLogGetList(page, cond)
	if errr != nil {
		appG.ErrorResponse(404, "查询任务id日志失败")
		return
	}

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["start_time"] = time.Unix(v.CreateTime, 0).Format("2006-01-02 15:04:05")
		row["process_time"] = float64(v.ProcessTime) / 1000
		row["ouput_size"] = v.Output
		row["status"] = v.Status
		list[k] = row
	}
	paginatorMap["data"] = list
	appG.Response(200,paginatorMap)
}

// // 批量操作日志
func LogBatch(c *gin.Context) {
	appG := app.Gin{C: c}
	var json BatchJson
	if err := c.ShouldBindJSON(&json); err != nil {
		appG.ErrorResponse(404, "批量删除日志失败")
		return 
	}
	for _, id := range json.Ids {
		if id < 1 {
			continue
		}
		models.TaskLogDelById(id)
		
	}

	appG.Response(200,"批量删除日志成功")
}
