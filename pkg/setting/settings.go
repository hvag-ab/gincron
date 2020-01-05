package setting

const (
	//分页
	PageSize = 10
	//token密码
	JwtSecret = "hvag"
	//查看静态文件路径
	PrefixUrl = "http://127.0.0.1:7890"
	//运行时数据存储
	RuntimeRootPath = "./runtime/"
	

	ProjectName = "gincron"
	//debug or release
	RunMode = "debug"
	HttpPort = 8000
	ReadTimeout = 60
	WriteTimeout = 60

	//[log]
	ServiceName = "go-gin-cron"
	Gin = "./logs/gin.log"
	Apps = "./logs/app.log"
	Http = "./logs/http.log"

	//[database]
	Type = "mysql"
	User = "root"
	Password = "hvag"
	Host = "127.0.0.1"
	Port = "3306"
	Name = "sche"
	TablePrefix = "hvag_"

	//定时任务设置 
	WorkPoolSize = 10 //并发任务数	

)