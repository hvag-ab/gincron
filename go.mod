module myproject

go 1.12

replace (
	github.com/Unknwon/com => github.com/unknwon/com v0.0.0-20190804042917-757f69c95f3e
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191227163750-53104e6ec876
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191228213918-04cbcbbfeed8
)

require (
	github.com/Unknwon/com v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.5.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.9
	github.com/json-iterator/go v1.1.7
	github.com/lisijie/cron v0.0.0-20151225081149-1c5ac61b9f22
	github.com/robfig/cron/v3 v3.0.0
	github.com/unknwon/com v1.0.1 // indirect
	go.uber.org/zap v1.13.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
