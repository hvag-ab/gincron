package jobs

import (
	"github.com/robfig/cron/v3"
	"sync"
	"myproject/pkg/setting"
)

var (
	mainCron *cron.Cron
	workPool chan bool
	lock     sync.Mutex
)

func init() {
	workPool = make(chan bool, setting.WorkPoolSize)
	mainCron = cron.New()
	mainCron.Start()
}

func AddJob(spec string, job *Job) bool {
	lock.Lock()//防止在并发的时候添加多个相同job
	defer lock.Unlock()

	if GetEntryById(job.id) != nil { //如果存在这个job 那么就添加失败 不需要重复添加
		return false
	}
	_,err := mainCron.AddJob(spec, job)
	if err != nil{
		return false 
	}else{
		return true
	}

}

func RemoveJob(id int) {
	entry := GetEntryById(id)
	if entry == nil {
		return 
	}
	ID := entry.ID
	mainCron.Remove(ID)
	
}

func GetEntryById(id int) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return &e
			}
		}
	}
	return nil
}

//获取cron正在调度的entry  size表示最多取多少个entry
func GetEntries(size int) []cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}
