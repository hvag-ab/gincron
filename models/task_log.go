package models

import (
	"myproject/db"
	"myproject/pkg/util"
	"fmt"
)

type TaskLog struct {
	Id          int
	TaskId      int
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}



func TaskLogAdd(t *TaskLog) (int64, bool) {

	affected, err := db.DB.Table("task_log").Insert(t)
	if err != nil {
		return affected, false
	}
	return affected, true
}

func TaskLogCount(condition map[string]interface{}) (int64,error) {

	task := new(TaskLog)

	query := db.DB.Table("task_log")
	
	if len(condition) != 0 {
		for k,v := range condition {
			query = query.Where(fmt.Sprintf("task_log.%s = ?",k),v)
		}
	}
	resultCount, terr := query.Count(task)//特别注意链式调用只能调用一次 也就是Count后就不能再query基础上再find了

	return resultCount,terr
}

func TaskLogGetList(page int, condition map[string]interface{}) ([]*TaskLog, map[string]interface{},error) {

	tasks := make([]*TaskLog, 0)
	
	resultCount, terr := TaskLogCount(condition)
	if terr != nil{
		return tasks,map[string]interface{}{},terr
	}

	paginatorMap := util.Paginator(resultCount,page)


	query := db.DB.Table("task_log")
	
	if len(condition) != 0 {
		for k,v := range condition {
			query = query.Where(fmt.Sprintf("task_log.%s = ?",k),v)
		}
	}
	
	err := query.Limit(paginatorMap["pageSize"].(int), paginatorMap["offset"].(int)).OrderBy("create_time desc").Find(&tasks)

	return tasks,paginatorMap,err
}


func TaskLogGetById(id int) (*TaskLog, error) {
	tasklog := &TaskLog{}

	_,err := db.DB.Table("task_log").Where("id =?", id).Get(tasklog)
	if err != nil {
		return nil, err
	}
	return tasklog, nil
}

func TaskLogDelById(id int) error {
	task := &TaskLog{}
	_, err := db.DB.Table("task_log").Where("id =?", id).Delete(task)
	return err
}

func TaskLogDelByTaskId(taskId int) error {
	task := &TaskLog{}
	_, err := db.DB.Table("task_log").Where("task_id=?",taskId).Delete(task)
	return err
}
