package models

import (
	"myproject/db"
	"time"
	"fmt"
	"myproject/pkg/util"
)

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时

	SHELL = 0
	API = 1
	PYTHON = 2
)

type Task struct {
	Id           int
	TaskName     string `form:"task_name"  json:"task_name" binding:"required"`
	TaskType     int `form:"task_type"  json:"task_type"`
	Description  string `form:"description"  json:"description"`
	CronSpec     string `form:"cron_spec"  json:"cron_spec" binding:"required"`
	Concurrent   int `form:"concurrent"  json:"concurrent" binding:"gte=0,lte=1"`
	Command      string `form:"command"  json:"command" binding:"required"`
	Status       int 
	Timeout      int `form:"timeout"  json:"timeout"`
	ExecuteTimes int
	PrevTime     int64
	CreateTime   int64
}


func (t *Task) Update() bool {
	affected, err := db.DB.Table("task").Where("task.id=?",t.Id).Update(t)
	if err != nil {
		return false
	}
	if affected == 0 {
		return false
	}else{
		return true
	}

}

func (t *Task) UpdateStatus() bool {
	/*
	这里需要注意，Update会自动从user结构体中提取非0和非nil得值作为需要更新的内容，因此，如果需要更新一个值为0，则此种方法将无法实现，因此有两种选择：
1.通过添加Cols函数指定需要更新结构体中的哪些值，未指定的将不更新，指定了的即使为0也会更新。
affected, err := engine.Id(id).Cols("age").Update(&user)
2.通过传入map[string]interface{}来进行更新，但这时需要额外指定更新到哪个表，因为通过map是无法自动检测更新哪个表的。
affected, err := engine.Table(new(User)).Id(id).Update(map[string]interface{}{"age":0})
	*/
	affected, err := db.DB.Table("task").Where("task.id=?",t.Id).Update(map[string]interface{}{"status":t.Status})
	if err != nil {
		return false
	}
	if affected == 0 {
		return false
	}else{
		return true
	}

}

func TaskAdd(task *Task) (int64, bool) {
	if task.TaskName == "" {
		return 0, false
	}
	if task.CronSpec == "" {
		return 0, false
	}
	if task.Command == "" {
		return 0, false
	}
	if task.CreateTime == 0 {
		task.CreateTime = time.Now().Unix()
	}
	affected, err := db.DB.Table("task").Insert(task)
	if err != nil {
		return affected, false
	}
	return affected, true
}

func TaskCount(condition map[string]interface{}) (int64,error) {

	task := new(Task)

	query := db.DB.Table("task")
	
	if len(condition) != 0 {
		for k,v := range condition {
			query = query.Where(fmt.Sprintf("task.%s = ?",k),v)
		}
	}
	resultCount, terr := query.Count(task)//特别注意链式调用只能调用一次 也就是Count后就不能再query基础上再find了

	return resultCount,terr
}

func TaskGetList(page int, pageSize int, condition map[string]interface{}) ([]*Task, map[string]interface{},error) {

	tasks := make([]*Task, 0)
	
	resultCount, terr := TaskCount(condition)
	if terr != nil{
		return tasks,map[string]interface{}{},terr
	}

	paginatorMap := util.Paginator(resultCount,page, pageSize)


	query := db.DB.Table("task")
	
	if len(condition) != 0 {
		for k,v := range condition {
			query = query.Where(fmt.Sprintf("task.%s = ?",k),v)
		}
	}
	
	err := query.Limit(paginatorMap["pageSize"].(int), paginatorMap["offset"].(int)).OrderBy("create_time desc").Find(&tasks)

	return tasks,paginatorMap,err
}


func TaskGetById(id int) (*Task, error) {
	task := &Task{}

	has,err := db.DB.Table("task").Where("task.id = ?", id).Get(task)
	if err != nil {
		return nil, err
	}
	if has == false {
		return nil,nil
	}
	return task, nil
}

func TaskDel(id int) (int64,error) {
	task := &Task{}
	affected, err := db.DB.Table("task").Where("task.id = ?", id).Delete(task)
	return affected, err
}

func TaskInitList(page, pageSize int, status int) ([]*Task, error) {
	offset := (page - 1) * pageSize
	tasks := make([]*Task, 0)

	query := db.DB.Table("task")

	query = query.Where("task.status = ?", status)

	err:=query.Limit(pageSize, offset).OrderBy("create_time desc").Find(&tasks)

	return tasks,err

}
