package jobs

import (

	"fmt"
	"myproject/models"
	"runtime/debug"
	"time"
	"myproject/pkg/logging"
	"go.uber.org/zap"
)



type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *models.Task                                      // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	if task.Id < 1 {
		logging.AppLogger.Fatal("ToJob: 缺少id", zap.Int("jobid",task.Id))
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task)
	job.task = task
	job.Concurrent = task.Concurrent == 1
	return job, nil
}

func NewCommandJob(task *models.Task) *Job {
	job := &Job{
		id:   task.Id,
		name: task.TaskName,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		switch task.TaskType {

		case models.API:
			return HttpGet(timeout, task)
		default:
			return Interprete(timeout, task)
		}
	}

	return job
}



//自定义job 必须实现RUN接口 
func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		logging.AppLogger.Warn("上一次执行尚未结束，本次被忽略", zap.Int("jobid",j.id))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			logging.AppLogger.Fatal("err:", zap.String("err",string(debug.Stack())))
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	logging.Logger.Info("开始执行任务:", zap.Int("jobid",j.id))

	j.status++
	defer func() {
		j.status--
	}()

	t := time.Now()
	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	cmdOut, cmdErr, err, isTimeout := j.runFunc(timeout)

	ut := time.Now().Sub(t) / time.Millisecond

	// 插入日志
	log := new(models.TaskLog)
	log.TaskId = j.id
	log.Output = cmdOut
	log.Error = cmdErr
	log.ProcessTime = int(ut)
	log.CreateTime = t.Unix()

	if isTimeout {
		log.Status = models.TASK_TIMEOUT
		log.Error = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), cmdErr)
	} else if err != nil {
		log.Status = models.TASK_ERROR
		log.Error = err.Error() + ":" + cmdErr
	}

	j.logId, _ = models.TaskLogAdd(log)

	// 更新上次执行时间
	j.task.PrevTime = t.Unix()
	j.task.ExecuteTimes++
	j.task.Update()

	// 发送邮件通知
	if (err != nil) {
		fmt.Print("send email")
	}
}
