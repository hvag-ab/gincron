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
	//上传图片路径
	ImageSavePath = "upload/images/"
	//最大图片大小MB
	ImageMaxSize = 5
	//允许上传的文件格式
	ImageAllowExts = ".jpg,.jpeg,.png"
	//导出文件路径
	ExportSavePath = "export/"

	ProjectName = "demo"
	//debug or release
	RunMode = "debug"
	HttpPort = 8000
	ReadTimeout = 60
	WriteTimeout = 60

	//[log]
	ServiceName = "go-gin-demo"
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

	//[redis]
	RedisHost = "127.0.0.1"
	RedisPort = "6379"
	RedisPassword = ""
	RedisDb = 9
	//Cluster = false
	//Hosts =
	MaxIdle = 5
	MaxActive = 5
	IdleTimeout = 240

	//[session]
	SessionHost = "127.0.0.1"
	SessionPort = "6379"
	SessionPassword = ""
	SessionDB = "10"

	//project
	ImgPageSize = 10
	ViewImageSavePath = "/upload/views/"

	//定时任务设置 
	WorkPoolSize = 10 //并发任务数
	Interpreter = "python3" //执行python解释器 
	// Interpreter = "linux" //执行linux shell

	//权限模型设置
	RbacModel = `
	[request_definition]        
	r = sub, obj, act
	
	[policy_definition]         
	p = sub, obj, act
	
	[role_definition]           
	g = _, _
	
	[policy_effect]              
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
	`
	AbacModel = `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == r.obj.Owner
	`

)